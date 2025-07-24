package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// User represents a user in the system
type User struct {
	ID        int64     `json:"id" db:"id"`
	Email     string    `json:"email" db:"email"`
	FullName  string    `json:"full_name" db:"full_name"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	LastLogin *time.Time `json:"last_login,omitempty" db:"last_login"`
}

// RefreshToken represents a refresh token
type RefreshToken struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	Token      string    `json:"token" db:"token"`
	ExpiryDate time.Time `json:"expiry_date" db:"expiry_date"`
	DeviceID   *string   `json:"device_id,omitempty" db:"device_id"`
}

// Enums for education and expertise levels
type EducationEnum string

const (
	EducationNoDegree EducationEnum = "NO_DEGREE"
	EducationBachelor EducationEnum = "BACHELOR"
	EducationMaster   EducationEnum = "MASTER"
	EducationPhD      EducationEnum = "PHD"
)

type ExpertiseLevelEnum string

const (
	ExpertiseLevelEntry      ExpertiseLevelEnum = "ENTRY"
	ExpertiseLevelJunior     ExpertiseLevelEnum = "JUNIOR"
	ExpertiseLevelMid        ExpertiseLevelEnum = "MID"
	ExpertiseLevelSenior     ExpertiseLevelEnum = "SENIOR"
	ExpertiseLevelResearcher ExpertiseLevelEnum = "RESEARCHER"
)

// Implement the sql.Scanner interface for EducationEnum
func (e *EducationEnum) Scan(value interface{}) error {
	if value == nil {
		*e = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*e = EducationEnum(v)
	case []byte:
		*e = EducationEnum(string(v))
	default:
		return fmt.Errorf("cannot scan %T into EducationEnum", value)
	}
	return nil
}

// Implement the driver.Valuer interface for EducationEnum
func (e EducationEnum) Value() (driver.Value, error) {
	return string(e), nil
}

// Implement the sql.Scanner interface for ExpertiseLevelEnum
func (e *ExpertiseLevelEnum) Scan(value interface{}) error {
	if value == nil {
		*e = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*e = ExpertiseLevelEnum(v)
	case []byte:
		*e = ExpertiseLevelEnum(string(v))
	default:
		return fmt.Errorf("cannot scan %T into ExpertiseLevelEnum", value)
	}
	return nil
}

// Implement the driver.Valuer interface for ExpertiseLevelEnum
func (e ExpertiseLevelEnum) Value() (driver.Value, error) {
	return string(e), nil
}

// Authentication request/response models
type SignupRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

type AuthResponse struct {
	User   *User          `json:"user"`
	Tokens *TokenResponse `json:"tokens"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Helper methods
func (u *User) IsValid() bool {
	return u.Email != "" && u.FullName != "" && u.Password != ""
}

func (rt *RefreshToken) IsExpired() bool {
	return time.Now().After(rt.ExpiryDate)
}

func (rt *RefreshToken) IsValid() bool {
	return rt.Token != "" && rt.UserID > 0 && !rt.IsExpired()
}
