package services

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

var (
	ErrProjectNotFound       = errors.New("project not found")
	ErrProjectRoleNotFound   = errors.New("project role not found")
	ErrUnauthorizedProject   = errors.New("unauthorized access to project")
	ErrProjectRoleNotOpen    = errors.New("project role is not open for applications")
)

// ProjectService handles project-related operations
type ProjectService struct {
	db *sqlx.DB
}

// NewProjectService creates a new project service
func NewProjectService(db *sqlx.DB) *ProjectService {
	return &ProjectService{
		db: db,
	}
}

// CreateProject creates a new project
func (s *ProjectService) CreateProject(userID int64, req *models.CreateProjectRequest) (*models.ProjectResponse, error) {
	project := &models.Project{
		RecruiterID: userID,
		Title:       req.Title,
		Description: req.Description,
		TeamSize:    req.TeamSize,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		IsActive:    true,
	}

	if project.TeamSize == 0 {
		project.TeamSize = 5 // Default team size
	}

	query := `
		INSERT INTO project (recruiter_id, title, description, team_size, created_at, is_active)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at
	`
	err := s.db.QueryRowx(query, project.RecruiterID, project.Title, project.Description,
		project.TeamSize, project.CreatedAt, project.IsActive).Scan(&project.ID, &project.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create project: %w", err)
	}

	// Convert to response
	response := &models.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		TeamSize:    project.TeamSize,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		IsActive:    project.IsActive,
	}

	return response, nil
}

// GetUserProjects retrieves all projects for a user
func (s *ProjectService) GetUserProjects(userID int64) ([]models.ProjectResponse, error) {
	var projects []models.Project
	query := `
		SELECT id, recruiter_id, title, description, team_size, created_at, is_active
		FROM project
		WHERE recruiter_id = $1 AND is_active = true
		ORDER BY created_at DESC
	`
	err := s.db.Select(&projects, query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user projects: %w", err)
	}

	// Convert to response
	var responses []models.ProjectResponse
	for _, project := range projects {
		response := models.ProjectResponse{
			ID:          project.ID,
			Title:       project.Title,
			Description: project.Description,
			TeamSize:    project.TeamSize,
			CreatedAt:   project.CreatedAt,
			UpdatedAt:   project.UpdatedAt,
			IsActive:    project.IsActive,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetProject retrieves a specific project by ID
func (s *ProjectService) GetProject(projectID int64) (*models.ProjectResponse, error) {
	var project models.Project
	query := `
		SELECT id, recruiter_id, title, description, team_size, created_at, is_active
		FROM project
		WHERE id = $1 AND is_active = true
	`
	err := s.db.Get(&project, query, projectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProjectNotFound
		}
		return nil, fmt.Errorf("failed to get project: %w", err)
	}

	// Get project roles
	roles, err := s.GetProjectRoles(projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project roles: %w", err)
	}

	response := &models.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		TeamSize:    project.TeamSize,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
		IsActive:    project.IsActive,
		Roles:       roles,
	}

	return response, nil
}

// CreateProjectRole creates a new role for a project
func (s *ProjectService) CreateProjectRole(userID, projectID int64, req *models.CreateProjectRoleRequest) (*models.ProjectRoleResponse, error) {
	// Check if user owns the project
	if !s.userOwnsProject(userID, projectID) {
		return nil, ErrUnauthorizedProject
	}

	// Convert technologies array to comma-separated string
	var techString *string
	if len(req.Technologies) > 0 {
		techStr := strings.Join(req.Technologies, ",")
		techString = &techStr
	}

	role := &models.ProjectRole{
		ProjectID:      projectID,
		RoleName:       req.RoleName,
		ExpertiseLevel: req.ExpertiseLevel,
		Technologies:   techString,
		IsOpen:         true,
	}

	query := `
		INSERT INTO project_role (project_id, role_name, expertise_level, technologies, is_open)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := s.db.QueryRowx(query, role.ProjectID, role.RoleName, role.ExpertiseLevel,
		role.Technologies, role.IsOpen).Scan(&role.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create project role: %w", err)
	}

	// Get project title
	var projectTitle string
	err = s.db.Get(&projectTitle, "SELECT title FROM project WHERE id = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project title: %w", err)
	}

	// Convert to response
	response := &models.ProjectRoleResponse{
		ID:             role.ID,
		ProjectID:      role.ProjectID,
		RoleName:       role.RoleName,
		ExpertiseLevel: role.ExpertiseLevel,
		Technologies:   req.Technologies,
		IsOpen:         role.IsOpen,
		ProjectTitle:   projectTitle,
	}

	return response, nil
}

// GetProjectRoles retrieves all roles for a project
func (s *ProjectService) GetProjectRoles(projectID int64) ([]models.ProjectRole, error) {
	var roles []models.ProjectRole
	query := `
		SELECT id, project_id, role_name, expertise_level, technologies, is_open
		FROM project_role
		WHERE project_id = $1
		ORDER BY id
	`
	err := s.db.Select(&roles, query, projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project roles: %w", err)
	}

	return roles, nil
}

// GetProjectRolesResponse retrieves all roles for a project with response format
func (s *ProjectService) GetProjectRolesResponse(projectID int64) ([]models.ProjectRoleResponse, error) {
	roles, err := s.GetProjectRoles(projectID)
	if err != nil {
		return nil, err
	}

	// Get project title
	var projectTitle string
	err = s.db.Get(&projectTitle, "SELECT title FROM project WHERE id = $1", projectID)
	if err != nil {
		return nil, fmt.Errorf("failed to get project title: %w", err)
	}

	// Convert to response format
	var responses []models.ProjectRoleResponse
	for _, role := range roles {
		var technologies []string
		if role.Technologies != nil && *role.Technologies != "" {
			technologies = strings.Split(*role.Technologies, ",")
		}

		response := models.ProjectRoleResponse{
			ID:             role.ID,
			ProjectID:      role.ProjectID,
			RoleName:       role.RoleName,
			ExpertiseLevel: role.ExpertiseLevel,
			Technologies:   technologies,
			IsOpen:         role.IsOpen,
			ProjectTitle:   projectTitle,
		}
		responses = append(responses, response)
	}

	return responses, nil
}

// GetProjectRole retrieves a specific project role
func (s *ProjectService) GetProjectRole(roleID int64) (*models.ProjectRole, error) {
	var role models.ProjectRole
	query := `
		SELECT id, project_id, role_name, expertise_level, technologies, is_open
		FROM project_role
		WHERE id = $1
	`
	err := s.db.Get(&role, query, roleID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrProjectRoleNotFound
		}
		return nil, fmt.Errorf("failed to get project role: %w", err)
	}

	return &role, nil
}

// UpdateProject updates a project
func (s *ProjectService) UpdateProject(userID, projectID int64, req *models.UpdateProjectRequest) (*models.ProjectResponse, error) {
	// Check if user owns the project
	if !s.userOwnsProject(userID, projectID) {
		return nil, ErrUnauthorizedProject
	}

	// Build update query dynamically
	updateFields := []string{}
	updateValues := []interface{}{}
	valueIndex := 1

	if req.Title != nil {
		updateFields = append(updateFields, fmt.Sprintf("title = $%d", valueIndex))
		updateValues = append(updateValues, *req.Title)
		valueIndex++
	}
	if req.Description != nil {
		updateFields = append(updateFields, fmt.Sprintf("description = $%d", valueIndex))
		updateValues = append(updateValues, *req.Description)
		valueIndex++
	}
	if req.TeamSize != nil {
		updateFields = append(updateFields, fmt.Sprintf("team_size = $%d", valueIndex))
		updateValues = append(updateValues, *req.TeamSize)
		valueIndex++
	}

	if len(updateFields) == 0 {
		return s.GetProject(projectID)
	}

	query := fmt.Sprintf("UPDATE project SET %s WHERE id = $%d",
		strings.Join(updateFields, ", "), valueIndex)
	updateValues = append(updateValues, projectID)

	_, err := s.db.Exec(query, updateValues...)
	if err != nil {
		return nil, fmt.Errorf("failed to update project: %w", err)
	}

	return s.GetProject(projectID)
}

// DeleteProject soft deletes a project
func (s *ProjectService) DeleteProject(userID, projectID int64) error {
	// Check if user owns the project
	if !s.userOwnsProject(userID, projectID) {
		return ErrUnauthorizedProject
	}

	query := "UPDATE project SET is_active = false WHERE id = $1"
	_, err := s.db.Exec(query, projectID)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}

	return nil
}

// userOwnsProject checks if a user owns a project
func (s *ProjectService) userOwnsProject(userID, projectID int64) bool {
	var count int
	query := "SELECT COUNT(*) FROM project WHERE id = $1 AND recruiter_id = $2 AND is_active = true"
	err := s.db.Get(&count, query, projectID, userID)
	return err == nil && count > 0
}

// GetProjectOwner gets the owner of a project
func (s *ProjectService) GetProjectOwner(projectID int64) (int64, error) {
	var recruiterID int64
	query := "SELECT recruiter_id FROM project WHERE id = $1 AND is_active = true"
	err := s.db.Get(&recruiterID, query, projectID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, ErrProjectNotFound
		}
		return 0, fmt.Errorf("failed to get project owner: %w", err)
	}
	return recruiterID, nil
}