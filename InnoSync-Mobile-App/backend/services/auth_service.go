package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
	"github.com/team-18-project/InnoSync/backend/utils"
)

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidRefreshToken = errors.New("invalid refresh token")
)

// AuthService handles authentication logic
type AuthService struct {
	db           *sqlx.DB
	tokenManager *utils.TokenManager
}

// NewAuthService creates a new authentication service
func NewAuthService(db *sqlx.DB, tokenManager *utils.TokenManager) *AuthService {
	return &AuthService{
		db:           db,
		tokenManager: tokenManager,
	}
}

// Signup creates a new user account
func (s *AuthService) Signup(req *models.SignupRequest) (*models.AuthResponse, error) {
	// Check if user already exists
	var existingUser models.User
	err := s.db.Get(&existingUser, "SELECT id FROM users WHERE email = $1", req.Email)
	if err == nil {
		return nil, ErrEmailAlreadyExists
	}
	if err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed to check existing user: %w", err)
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create user
	user := &models.User{
		Email:     req.Email,
		FullName:  req.FullName,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	query := `
		INSERT INTO users (email, full_name, password, created_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id, email, full_name, created_at
	`
	err = s.db.QueryRowx(query, user.Email, user.FullName, user.Password, user.CreatedAt).StructScan(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate tokens
	accessToken, refreshToken, err := s.tokenManager.GenerateTokens(user.ID, user.Email, user.FullName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Store refresh token
	if err := s.storeRefreshToken(user.ID, refreshToken); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    s.tokenManager.GetAccessTokenTTL(),
	}

	return &models.AuthResponse{
		User:   user,
		Tokens: tokens,
	}, nil
}

// Login authenticates a user
func (s *AuthService) Login(req *models.LoginRequest) (*models.AuthResponse, error) {
	// Find user by email
	var user models.User
	err := s.db.Get(&user, "SELECT id, email, full_name, password FROM users WHERE email = $1", req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidCredentials
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	// Check password
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, ErrInvalidCredentials
	}

	// Update last login
	now := time.Now()
	_, err = s.db.Exec("UPDATE users SET last_login = $1 WHERE id = $2", now, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update last login: %w", err)
	}

	// Generate tokens
	accessToken, refreshToken, err := s.tokenManager.GenerateTokens(user.ID, user.Email, user.FullName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Store refresh token
	if err := s.storeRefreshToken(user.ID, refreshToken); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	tokens := &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    s.tokenManager.GetAccessTokenTTL(),
	}

	// Don't return password in response
	user.Password = ""

	return &models.AuthResponse{
		User:   &user,
		Tokens: tokens,
	}, nil
}

// RefreshToken refreshes an access token
func (s *AuthService) RefreshToken(req *models.RefreshTokenRequest) (*models.TokenResponse, error) {
	// Validate refresh token
	claims, err := s.tokenManager.ValidateRefreshToken(req.RefreshToken)
	if err != nil {
		return nil, ErrInvalidRefreshToken
	}

	// Check if refresh token exists and is valid in database
	var refreshToken models.RefreshToken
	err = s.db.Get(&refreshToken, "SELECT id, user_id, token, expiry_date FROM refresh_token WHERE token = $1", req.RefreshToken)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrInvalidRefreshToken
		}
		return nil, fmt.Errorf("failed to find refresh token: %w", err)
	}

	// Check if token is expired
	if refreshToken.IsExpired() {
		// Clean up expired token
		s.db.Exec("DELETE FROM refresh_token WHERE id = $1", refreshToken.ID)
		return nil, ErrInvalidRefreshToken
	}

	// Generate new tokens
	accessToken, newRefreshToken, err := s.tokenManager.GenerateTokens(claims.UserID, claims.Email, claims.FullName)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	// Update refresh token in database
	if err := s.updateRefreshToken(refreshToken.ID, newRefreshToken); err != nil {
		return nil, fmt.Errorf("failed to update refresh token: %w", err)
	}

	return &models.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		TokenType:    "Bearer",
		ExpiresIn:    s.tokenManager.GetAccessTokenTTL(),
	}, nil
}

// Logout invalidates a refresh token
func (s *AuthService) Logout(req *models.LogoutRequest) error {
	// Delete refresh token from database
	_, err := s.db.Exec("DELETE FROM refresh_token WHERE token = $1", req.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to delete refresh token: %w", err)
	}

	return nil
}

// storeRefreshToken stores a refresh token in the database
func (s *AuthService) storeRefreshToken(userID int64, token string) error {
	expiryDate := time.Now().Add(time.Duration(s.tokenManager.GetRefreshTokenTTL()) * time.Second)

	query := `
		INSERT INTO refresh_token (user_id, token, expiry_date)
		VALUES ($1, $2, $3)
	`
	_, err := s.db.Exec(query, userID, token, expiryDate)
	return err
}

// updateRefreshToken updates a refresh token in the database
func (s *AuthService) updateRefreshToken(tokenID int64, newToken string) error {
	expiryDate := time.Now().Add(time.Duration(s.tokenManager.GetRefreshTokenTTL()) * time.Second)

	query := `
		UPDATE refresh_token
		SET token = $1, expiry_date = $2
		WHERE id = $3
	`
	_, err := s.db.Exec(query, newToken, expiryDate, tokenID)
	return err
}

// GetUserByID retrieves a user by ID
func (s *AuthService) GetUserByID(userID int64) (*models.User, error) {
	var user models.User
	err := s.db.Get(&user, "SELECT id, email, full_name, created_at, last_login FROM users WHERE id = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserNotFound
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return &user, nil
}

// CleanupExpiredTokens removes expired refresh tokens
func (s *AuthService) CleanupExpiredTokens() error {
	_, err := s.db.Exec("DELETE FROM refresh_token WHERE expiry_date < NOW()")
	return err
}