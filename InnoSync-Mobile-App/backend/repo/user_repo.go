package repo

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	query := `INSERT INTO users (email, full_name, password) 
              VALUES ($1, $2, $3) RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query, user.Email, user.FullName, user.Password).
		Scan(&user.ID, &user.CreatedAt)
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT * FROM users WHERE email = $1`
	err := r.db.GetContext(ctx, user, query, email)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (r *UserRepository) GetUserProfile(ctx context.Context, userID int64) (*models.UserProfile, error) {
	profile := &models.UserProfile{}
	query := `SELECT * FROM user_profile WHERE user_id = $1`
	err := r.db.GetContext(ctx, profile, query, userID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return profile, err
}
