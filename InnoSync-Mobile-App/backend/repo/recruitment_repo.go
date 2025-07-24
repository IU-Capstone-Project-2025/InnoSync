package repo

import (
	"context"
	_ "database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

type RecruitmentRepository struct {
	db *sqlx.DB
}

func NewRecruitmentRepository(db *sqlx.DB) *RecruitmentRepository {
	return &RecruitmentRepository{db: db}
}

func (r *RecruitmentRepository) CreateApplication(ctx context.Context, app *models.RoleApplication) error {
	query := `INSERT INTO role_application (user_id, project_role_id) 
              VALUES ($1, $2) RETURNING id, applied_at, updated_at`
	return r.db.QueryRowContext(ctx, query, app.UserID, app.ProjectRoleID).
		Scan(&app.ID, &app.AppliedAt, &app.UpdatedAt)
}

func (r *RecruitmentRepository) GetUserInvitations(ctx context.Context, userID int64) ([]models.Invitation, error) {
	var invitations []models.Invitation
	query := `SELECT * FROM invitation WHERE user_id = $1`
	err := r.db.SelectContext(ctx, &invitations, query, userID)
	return invitations, err
}

func (r *RecruitmentRepository) UpdateApplicationStatus(ctx context.Context,
	appID int64, status models.ApplicationStatusEnum) error {

	query := `UPDATE role_application SET status = $1 WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, appID)
	return err
}
