package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

var (
	ErrInvitationNotFound       = errors.New("invitation not found")
	ErrInvitationAlreadyExists  = errors.New("invitation already exists")
	ErrUnauthorizedInvitation   = errors.New("unauthorized access to invitation")
	ErrInvalidInvitationStatus  = errors.New("invalid invitation status")
	ErrInvitationAlreadyReplied = errors.New("invitation already replied")
	ErrRecipientNotFound        = errors.New("recipient not found")
)

// InvitationService handles invitation-related operations
type InvitationService struct {
	db             *sqlx.DB
	projectService *ProjectService
}

// NewInvitationService creates a new invitation service
func NewInvitationService(db *sqlx.DB, projectService *ProjectService) *InvitationService {
	return &InvitationService{
		db:             db,
		projectService: projectService,
	}
}

// CreateInvitation creates a new invitation
func (s *InvitationService) CreateInvitation(userID int64, req *models.CreateInvitationRequest) (*models.InvitationResponse, error) {
	// Get project role details
	role, err := s.projectService.GetProjectRole(req.ProjectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if user is the project owner
	projectOwner, err := s.projectService.GetProjectOwner(role.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project owner: %w", err)
	}
	if projectOwner != userID {
		return nil, ErrUnauthorizedInvitation
	}

	// Check if recipient exists
	var recipientExists bool
	err = s.db.Get(&recipientExists, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", req.RecipientID)
	if err != nil {
		return nil, fmt.Errorf("failed to check recipient: %w", err)
	}
	if !recipientExists {
		return nil, ErrRecipientNotFound
	}

	// Check if invitation already exists
	var existingInvitation models.Invitation
	err = s.db.Get(&existingInvitation, "SELECT id FROM invitation WHERE project_role_id = $1 AND user_id = $2", req.ProjectRoleID, req.RecipientID)
	if err == nil {
		return nil, ErrInvitationAlreadyExists
	}
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing invitation: %w", err)
	}

	// Create invitation
	invitation := &models.Invitation{
		ProjectRoleID:    req.ProjectRoleID,
		UserID:           req.RecipientID,
		InvitationStatus: "INVITED",
		SentAt:           time.Now(),
		Message:          req.Message,
	}

	query := `
		INSERT INTO invitation (project_role_id, user_id, invitation_status, sent_at, message)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err = s.db.QueryRowx(query, invitation.ProjectRoleID, invitation.UserID, invitation.InvitationStatus,
		invitation.SentAt, invitation.Message).Scan(&invitation.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create invitation: %w", err)
	}

	// Get response data
	return s.getInvitationResponse(invitation.ID)
}

// GetSentInvitations retrieves all invitations sent by a user
func (s *InvitationService) GetSentInvitations(userID int64) ([]models.InvitationResponse, error) {
	query := `
		SELECT
			i.id,
			i.project_role_id,
			i.user_id AS recipient_id,
			i.invitation_status,
			i.sent_at,
			i.responded_at,
			i.message,
			pr.role_name,
			p.id AS project_id,
			p.title AS project_title,
			u.full_name AS recipient_name
		FROM invitation i
		JOIN project_role pr ON i.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		JOIN users u ON i.user_id = u.id
		WHERE p.recruiter_id = $1
		ORDER BY i.sent_at DESC
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get sent invitations: %w", err)
	}
	defer rows.Close()

	var invitations []models.InvitationResponse
	for rows.Next() {
		var inv models.InvitationResponse
		err := rows.Scan(
			&inv.ID,
			&inv.ProjectRoleID,
			&inv.RoleName,
			&inv.ProjectID,
			&inv.ProjectTitle,
			&inv.RecipientID,
			&inv.RecipientName,
			&inv.InvitationStatus,
			&inv.SentAt,
			&inv.RespondedAt,
			&inv.Message,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invitation: %w", err)
		}
		invitations = append(invitations, inv)
	}

	return invitations, nil
}

// GetReceivedInvitations retrieves all invitations received by a user
func (s *InvitationService) GetReceivedInvitations(userID int64) ([]models.InvitationResponse, error) {
	query := `
		SELECT
			i.id,
			i.project_role_id,
			i.user_id AS recipient_id,
			i.invitation_status,
			i.sent_at,
			i.responded_at,
			i.message,
			pr.role_name,
			p.id AS project_id,
			p.title AS project_title,
			u.full_name AS recipient_name
		FROM invitation i
		JOIN project_role pr ON i.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		JOIN users u ON i.user_id = u.id
		WHERE i.user_id = $1
		ORDER BY i.sent_at DESC
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get received invitations: %w", err)
	}
	defer rows.Close()

	var invitations []models.InvitationResponse
	for rows.Next() {
		var inv models.InvitationResponse
		err := rows.Scan(
			&inv.ID,               // i.id
			&inv.ProjectRoleID,    // i.project_role_id
			&inv.RecipientID,      // i.user_id AS recipient_id
			&inv.InvitationStatus, // i.invitation_status
			&inv.SentAt,           // i.sent_at
			&inv.RespondedAt,      // i.responded_at
			&inv.Message,          // i.message
			&inv.RoleName,         // pr.role_name
			&inv.ProjectID,        // p.id AS project_id
			&inv.ProjectTitle,     // p.title AS project_title
			&inv.RecipientName,    // u.full_name AS recipient_name
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invitation: %w", err)
		}
		invitations = append(invitations, inv)
	}

	return invitations, nil
}

// GetInvitation retrieves a specific invitation by ID
func (s *InvitationService) GetInvitation(invitationID int64) (*models.InvitationResponse, error) {
	return s.getInvitationResponse(invitationID)
}

// RespondToInvitation allows a user to respond to an invitation
func (s *InvitationService) RespondToInvitation(userID, invitationID int64, req *models.RespondInvitationRequest) (*models.InvitationResponse, error) {
	// Get the invitation
	var invitation models.Invitation
	query := `
		SELECT id, project_role_id, user_id, invitation_status, sent_at, responded_at, message
		FROM invitation
		WHERE id = $1
	`
	err := s.db.Get(&invitation, query, invitationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvitationNotFound
		}
		return nil, fmt.Errorf("failed to get invitation: %w", err)
	}

	// Check if user is the recipient
	if invitation.UserID != userID {
		return nil, ErrUnauthorizedInvitation
	}

	// Check if invitation is still pending
	if invitation.InvitationStatus != "INVITED" {
		return nil, ErrInvitationAlreadyReplied
	}

	// Validate response
	if req.Response != "ACCEPTED" && req.Response != "DECLINED" {
		return nil, ErrInvalidInvitationStatus
	}

	// Update invitation status
	now := time.Now()
	updateQuery := `
		UPDATE invitation
		SET invitation_status = $1, responded_at = $2
		WHERE id = $3
	`
	_, err = s.db.Exec(updateQuery, req.Response, now, invitationID)
	if err != nil {
		return nil, fmt.Errorf("failed to update invitation: %w", err)
	}

	// If invitation is accepted, you might want to:
	// 1. Create a project member record
	// 2. Send notification to project owner
	// 3. Close the role if it's filled
	// For now, we'll just update the status

	return s.getInvitationResponse(invitationID)
}

// RevokeInvitation allows project owner to revoke an invitation
func (s *InvitationService) RevokeInvitation(userID, invitationID int64) (*models.InvitationResponse, error) {
	// Get the invitation
	var invitation models.Invitation
	query := `
		SELECT id, project_role_id, user_id, invitation_status, sent_at, responded_at, message
		FROM invitation
		WHERE id = $1
	`
	err := s.db.Get(&invitation, query, invitationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvitationNotFound
		}
		return nil, fmt.Errorf("failed to get invitation: %w", err)
	}

	// Get project role to check ownership
	role, err := s.projectService.GetProjectRole(invitation.ProjectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if user is the project owner
	projectOwner, err := s.projectService.GetProjectOwner(role.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project owner: %w", err)
	}
	if projectOwner != userID {
		return nil, ErrUnauthorizedInvitation
	}

	// Check if invitation can be revoked (only pending invitations)
	if invitation.InvitationStatus != "INVITED" {
		return nil, ErrInvalidInvitationStatus
	}

	// Update invitation status
	now := time.Now()
	updateQuery := `
		UPDATE invitation
		SET invitation_status = $1, responded_at = $2
		WHERE id = $3
	`
	_, err = s.db.Exec(updateQuery, "REVOKED", now, invitationID)
	if err != nil {
		return nil, fmt.Errorf("failed to revoke invitation: %w", err)
	}

	return s.getInvitationResponse(invitationID)
}

// DeleteInvitation allows project owner to delete an invitation
func (s *InvitationService) DeleteInvitation(userID, invitationID int64) error {
	// Get the invitation
	var invitation models.Invitation
	query := `
		SELECT id, project_role_id, user_id, invitation_status, sent_at, responded_at, message
		FROM invitation
		WHERE id = $1
	`
	err := s.db.Get(&invitation, query, invitationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrInvitationNotFound
		}
		return fmt.Errorf("failed to get invitation: %w", err)
	}

	// Get project role to check ownership
	role, err := s.projectService.GetProjectRole(invitation.ProjectRoleID)
	if err != nil {
		return fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if user is the project owner
	projectOwner, err := s.projectService.GetProjectOwner(role.ProjectID)
	if err != nil {
		return fmt.Errorf("failed to get project owner: %w", err)
	}
	if projectOwner != userID {
		return ErrUnauthorizedInvitation
	}

	// Delete invitation
	deleteQuery := "DELETE FROM invitation WHERE id = $1"
	_, err = s.db.Exec(deleteQuery, invitationID)
	if err != nil {
		return fmt.Errorf("failed to delete invitation: %w", err)
	}

	return nil
}

// getInvitationResponse retrieves a full invitation response by ID
func (s *InvitationService) getInvitationResponse(invitationID int64) (*models.InvitationResponse, error) {
	query := `
		SELECT
			i.id,
			i.project_role_id,
			i.user_id AS recipient_id,
			i.invitation_status,
			i.sent_at,
			i.responded_at,
			i.message,
			pr.role_name,
			p.id AS project_id,
			p.title AS project_title,
			u.full_name AS recipient_name
		FROM invitation i
		JOIN project_role pr ON i.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		JOIN users u ON i.user_id = u.id
		WHERE i.id = $1
	`

	var inv models.InvitationResponse
	err := s.db.Get(&inv, query, invitationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvitationNotFound
		}
		return nil, fmt.Errorf("failed to get invitation response: %w", err)
	}

	return &inv, nil
}

// GetInvitationsByProjectRole retrieves all invitations for a specific project role
func (s *InvitationService) GetInvitationsByProjectRole(userID, projectRoleID int64) ([]models.InvitationResponse, error) {
	// Get project role to check ownership
	role, err := s.projectService.GetProjectRole(projectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if user is the project owner
	projectOwner, err := s.projectService.GetProjectOwner(role.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project owner: %w", err)
	}
	if projectOwner != userID {
		return nil, ErrUnauthorizedInvitation
	}

	query := `
		SELECT
			i.id, i.project_role_id, i.user_id, i.invitation_status, i.sent_at, i.responded_at, i.message,
			pr.role_name,
			p.id as project_id, p.title as project_title,
			u.full_name as recipient_name
		FROM invitation i
		JOIN project_role pr ON i.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		JOIN users u ON i.user_id = u.id
		WHERE i.project_role_id = $1
		ORDER BY i.sent_at DESC
	`

	rows, err := s.db.Query(query, projectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get invitations by project role: %w", err)
	}
	defer rows.Close()

	var invitations []models.InvitationResponse
	for rows.Next() {
		var inv models.InvitationResponse
		err := rows.Scan(
			&inv.ID,
			&inv.ProjectRoleID,
			&inv.RoleName,
			&inv.ProjectID,
			&inv.ProjectTitle,
			&inv.RecipientID,
			&inv.RecipientName,
			&inv.InvitationStatus,
			&inv.SentAt,
			&inv.RespondedAt,
			&inv.Message,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan invitation: %w", err)
		}
		invitations = append(invitations, inv)
	}

	return invitations, nil
}
