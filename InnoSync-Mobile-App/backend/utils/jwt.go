package utils

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/team-18-project/InnoSync/backend/config"
)

var (
	ErrInvalidToken = errors.New("invalid token")
	ErrExpiredToken = errors.New("expired token")
)

// JWTClaims represents the claims stored in JWT tokens
type JWTClaims struct {
	UserID    int64  `json:"user_id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	TokenType string `json:"token_type"`
	jwt.RegisteredClaims
}

// TokenManager handles JWT token operations
type TokenManager struct {
	secretKey       string
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

// NewTokenManager creates a new token manager
func NewTokenManager(cfg *config.Config) *TokenManager {
	return &TokenManager{
		secretKey:       cfg.JWT.SecretKey,
		accessTokenTTL:  cfg.JWT.AccessTokenTTL,
		refreshTokenTTL: cfg.JWT.RefreshTokenTTL,
	}
}

// GenerateTokens generates both access and refresh tokens
func (tm *TokenManager) GenerateTokens(userID int64, email, fullName string) (string, string, error) {
	// Generate access token
	accessToken, err := tm.generateToken(userID, email, fullName, "access", tm.accessTokenTTL)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err := tm.generateToken(userID, email, fullName, "refresh", tm.refreshTokenTTL)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate refresh token: %w", err)
	}

	return accessToken, refreshToken, nil
}

// generateToken generates a JWT token with the given parameters
func (tm *TokenManager) generateToken(userID int64, email, fullName, tokenType string, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := JWTClaims{
		UserID:    userID,
		Email:     email,
		FullName:  fullName,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userID, 10),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
			NotBefore: jwt.NewNumericDate(now),
			ID:        uuid.New().String(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tm.secretKey))
}

// ValidateToken validates a JWT token and returns the claims
func (tm *TokenManager) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(tm.secretKey), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// ValidateAccessToken validates an access token
func (tm *TokenManager) ValidateAccessToken(tokenString string) (*JWTClaims, error) {
	claims, err := tm.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "access" {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// ValidateRefreshToken validates a refresh token
func (tm *TokenManager) ValidateRefreshToken(tokenString string) (*JWTClaims, error) {
	claims, err := tm.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	if claims.TokenType != "refresh" {
		return nil, ErrInvalidToken
	}

	return claims, nil
}

// GetAccessTokenTTL returns the access token TTL in seconds
func (tm *TokenManager) GetAccessTokenTTL() int64 {
	return int64(tm.accessTokenTTL.Seconds())
}

// GetRefreshTokenTTL returns the refresh token TTL in seconds
func (tm *TokenManager) GetRefreshTokenTTL() int64 {
	return int64(tm.refreshTokenTTL.Seconds())
}