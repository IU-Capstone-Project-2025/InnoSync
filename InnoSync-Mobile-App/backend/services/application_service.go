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
	ErrApplicationNotFound      = errors.New("application not found")
	ErrApplicationAlreadyExists = errors.New("application already exists")
	ErrUnauthorizedApplication  = errors.New("unauthorized access to application")
	ErrInvalidApplicationStatus = errors.New("invalid application status")
)

// ApplicationService handles application-related operations
type ApplicationService struct {
	db             *sqlx.DB
	projectService *ProjectService
}

// NewApplicationService creates a new application service
func NewApplicationService(db *sqlx.DB, projectService *ProjectService) *ApplicationService {
	return &ApplicationService{
		db:             db,
		projectService: projectService,
	}
}

// CreateApplication creates a new role application
func (s *ApplicationService) CreateApplication(userID, projectRoleID int64, req *models.CreateApplicationRequest) (*models.ApplicationResponse, error) {
	// Check if application already exists
	var existingApp models.RoleApplication
	err := s.db.Get(&existingApp, "SELECT id FROM role_application WHERE user_id = $1 AND project_role_id = $2", userID, projectRoleID)
	if err == nil {
		return nil, ErrApplicationAlreadyExists
	}
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing application: %w", err)
	}

	// Get project role details
	role, err := s.projectService.GetProjectRole(projectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if role is open for applications
	if !role.IsOpen {
		return nil, ErrProjectRoleNotOpen
	}

	// Create application
	application := &models.RoleApplication{
		UserID:        userID,
		ProjectRoleID: projectRoleID,
		Status:        models.ApplicationStatusPending,
		AppliedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CoverLetter:   req.CoverLetter,
	}

	query := `
		INSERT INTO role_application (user_id, project_role_id, status, applied_at, updated_at, cover_letter)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`
	err = s.db.QueryRowx(query, application.UserID, application.ProjectRoleID, application.Status,
		application.AppliedAt, application.UpdatedAt, application.CoverLetter).Scan(&application.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create application: %w", err)
	}

	// Get response data
	return s.getApplicationResponse(application.ID)
}

// GetUserApplications retrieves all applications for a user
func (s *ApplicationService) GetUserApplications(userID int64) ([]models.ApplicationResponse, error) {
	query := `
		SELECT
			ra.id, ra.user_id, ra.project_role_id, ra.status, ra.applied_at, ra.updated_at, ra.cover_letter,
			u.full_name as user_full_name,
			pr.role_name,
			p.id as project_id, p.title as project_title
		FROM role_application ra
		JOIN users u ON ra.user_id = u.id
		JOIN project_role pr ON ra.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		WHERE ra.user_id = $1
		ORDER BY ra.applied_at DESC
	`

	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user applications: %w", err)
	}
	defer rows.Close()

	var applications []models.ApplicationResponse
	for rows.Next() {
		var app models.ApplicationResponse
		err := rows.Scan(&app.ID, &app.UserID, &app.ProjectRoleID, &app.Status, &app.AppliedAt,
			&app.UpdatedAt, &app.CoverLetter, &app.UserFullName, &app.RoleName, &app.ProjectID, &app.ProjectTitle)
		if err != nil {
			return nil, fmt.Errorf("failed to scan application: %w", err)
		}
		applications = append(applications, app)
	}

	return applications, nil
}

// GetRoleApplications retrieves all applications for a specific role
func (s *ApplicationService) GetRoleApplications(userID, projectRoleID int64) ([]models.ApplicationResponse, error) {
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
		return nil, ErrUnauthorizedApplication
	}

	query := `
		SELECT
			ra.id, ra.user_id, ra.project_role_id, ra.status, ra.applied_at, ra.updated_at, ra.cover_letter,
			u.full_name as user_full_name,
			pr.role_name,
			p.id as project_id, p.title as project_title
		FROM role_application ra
		JOIN users u ON ra.user_id = u.id
		JOIN project_role pr ON ra.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		WHERE ra.project_role_id = $1
		ORDER BY ra.applied_at DESC
	`

	rows, err := s.db.Query(query, projectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get role applications: %w", err)
	}
	defer rows.Close()

	var applications []models.ApplicationResponse
	for rows.Next() {
		var app models.ApplicationResponse
		err := rows.Scan(&app.ID, &app.UserID, &app.ProjectRoleID, &app.Status, &app.AppliedAt,
			&app.UpdatedAt, &app.CoverLetter, &app.UserFullName, &app.RoleName, &app.ProjectID, &app.ProjectTitle)
		if err != nil {
			return nil, fmt.Errorf("failed to scan application: %w", err)
		}
		applications = append(applications, app)
	}

	return applications, nil
}

// GetApplication retrieves a specific application by ID
func (s *ApplicationService) GetApplication(applicationID int64) (*models.ApplicationResponse, error) {
	return s.getApplicationResponse(applicationID)
}

// UpdateApplicationStatus updates the status of an application
func (s *ApplicationService) UpdateApplicationStatus(userID, applicationID int64, req *models.UpdateApplicationStatusRequest) (*models.ApplicationResponse, error) {
	// Get the application
	var app models.RoleApplication
	query := `
		SELECT id, user_id, project_role_id, status, applied_at, updated_at, cover_letter
		FROM role_application
		WHERE id = $1
	`
	err := s.db.Get(&app, query, applicationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrApplicationNotFound
		}
		return nil, fmt.Errorf("failed to get application: %w", err)
	}

	// Get project role to check ownership
	role, err := s.projectService.GetProjectRole(app.ProjectRoleID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	// Check if user is the project owner
	projectOwner, err := s.projectService.GetProjectOwner(role.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project owner: %w", err)
	}
	if projectOwner != userID {
		return nil, ErrUnauthorizedApplication
	}

	// Validate status transition
	if !s.isValidStatusTransition(app.Status, req.Status) {
		return nil, ErrInvalidApplicationStatus
	}

	// Update application status
	updateQuery := `
		UPDATE role_application
		SET status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err = s.db.Exec(updateQuery, req.Status, time.Now(), applicationID)
	if err != nil {
		return nil, fmt.Errorf("failed to update application status: %w", err)
	}

	// If application is accepted, consider closing the role
	if req.Status == models.ApplicationStatusAccepted {
		// You might want to implement logic to close the role or limit applications
		// For now, we'll just keep it open
	}

	return s.getApplicationResponse(applicationID)
}

// DeleteApplication allows user to withdraw their application
func (s *ApplicationService) DeleteApplication(userID, applicationID int64) error {
	// Check if user owns the application
	var app models.RoleApplication
	query := "SELECT id, user_id, status FROM role_application WHERE id = $1"
	err := s.db.Get(&app, query, applicationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return ErrApplicationNotFound
		}
		return fmt.Errorf("failed to get application: %w", err)
	}

	if app.UserID != userID {
		return ErrUnauthorizedApplication
	}

	// Only allow withdrawal if application is pending or under review
	if app.Status != models.ApplicationStatusPending && app.Status != models.ApplicationStatusUnderReview {
		return ErrInvalidApplicationStatus
	}

	// Update status to withdrawn
	updateQuery := `
		UPDATE role_application
		SET status = $1, updated_at = $2
		WHERE id = $3
	`
	_, err = s.db.Exec(updateQuery, models.ApplicationStatusWithdrawn, time.Now(), applicationID)
	if err != nil {
		return fmt.Errorf("failed to withdraw application: %w", err)
	}

	return nil
}

// getApplicationResponse retrieves a full application response by ID
func (s *ApplicationService) getApplicationResponse(applicationID int64) (*models.ApplicationResponse, error) {
	query := `
		SELECT
			ra.id, ra.user_id, ra.project_role_id, ra.status, ra.applied_at, ra.updated_at, ra.cover_letter,
			u.full_name as user_full_name,
			pr.role_name,
			p.id as project_id, p.title as project_title
		FROM role_application ra
		JOIN users u ON ra.user_id = u.id
		JOIN project_role pr ON ra.project_role_id = pr.id
		JOIN project p ON pr.project_id = p.id
		WHERE ra.id = $1
	`

	var app models.ApplicationResponse
	err := s.db.Get(&app, query, applicationID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrApplicationNotFound
		}
		return nil, fmt.Errorf("failed to get application response: %w", err)
	}

	return &app, nil
}

// isValidStatusTransition checks if a status transition is valid
func (s *ApplicationService) isValidStatusTransition(currentStatus, newStatus models.ApplicationStatusEnum) bool {
	// Define valid transitions
	validTransitions := map[models.ApplicationStatusEnum][]models.ApplicationStatusEnum{
		models.ApplicationStatusPending: {
			models.ApplicationStatusUnderReview,
			models.ApplicationStatusAccepted,
			models.ApplicationStatusRejected,
		},
		models.ApplicationStatusUnderReview: {
			models.ApplicationStatusAccepted,
			models.ApplicationStatusRejected,
			models.ApplicationStatusPending,
		},
		models.ApplicationStatusAccepted: {
			models.ApplicationStatusRejected, // Can be revoked
		},
		models.ApplicationStatusRejected: {
			models.ApplicationStatusUnderReview, // Can be reconsidered
		},
		models.ApplicationStatusWithdrawn: {
			// No transitions from withdrawn
		},
	}

	allowedTransitions, exists := validTransitions[currentStatus]
	if !exists {
		return false
	}

	for _, allowedStatus := range allowedTransitions {
		if newStatus == allowedStatus {
			return true
		}
	}

	return false
}