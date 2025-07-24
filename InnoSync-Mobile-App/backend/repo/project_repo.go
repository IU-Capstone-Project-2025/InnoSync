package repo

import (
	"context"
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

func (r *ProjectRepository) CreateProject(ctx context.Context, project *models.Project) error {
	query := `INSERT INTO project (recruiter_id, title, description, team_size) 
              VALUES ($1, $2, $3, $4) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query,
		project.RecruiterID,
		project.Title,
		project.Description,
		project.TeamSize).
		Scan(&project.ID, &project.CreatedAt)
}

func (r *ProjectRepository) GetUserProjects(ctx context.Context, userID int64) ([]models.Project, error) {
	var projects []models.Project
	query := `SELECT * FROM project WHERE recruiter_id = $1`
	err := r.db.SelectContext(ctx, &projects, query, userID)
	return projects, err
}

func (r *ProjectRepository) GetProjectRoles(ctx context.Context, projectID int64) ([]models.ProjectRole, error) {
	var roles []models.ProjectRole
	query := `SELECT * FROM project_role WHERE project_id = $1`
	err := r.db.SelectContext(ctx, &roles, query, projectID)
	return roles, err
}
