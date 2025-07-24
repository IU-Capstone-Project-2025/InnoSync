package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Enums for recruitment status
type ApplicationStatusEnum string

const (
	ApplicationStatusPending     ApplicationStatusEnum = "PENDING"
	ApplicationStatusUnderReview ApplicationStatusEnum = "UNDER_REVIEW"
	ApplicationStatusAccepted    ApplicationStatusEnum = "ACCEPTED"
	ApplicationStatusRejected    ApplicationStatusEnum = "REJECTED"
	ApplicationStatusWithdrawn   ApplicationStatusEnum = "WITHDRAWN"
)

// Implement the sql.Scanner interface for ApplicationStatusEnum
func (e *ApplicationStatusEnum) Scan(value interface{}) error {
	if value == nil {
		*e = ""
		return nil
	}
	switch v := value.(type) {
	case string:
		*e = ApplicationStatusEnum(v)
	case []byte:
		*e = ApplicationStatusEnum(string(v))
	default:
		return fmt.Errorf("cannot scan %T into ApplicationStatusEnum", value)
	}
	return nil
}

// Implement the driver.Valuer interface for ApplicationStatusEnum
func (e ApplicationStatusEnum) Value() (driver.Value, error) {
	return string(e), nil
}

// Enums for recruitment status

// Invitation represents a project invitation
type Invitation struct {
	ID               int64      `json:"id" db:"id"`
	ProjectRoleID    int64      `json:"project_role_id" db:"project_role_id"`
	UserID           int64      `json:"user_id" db:"user_id"`
	InvitationStatus string     `json:"invitation_status" db:"invitation_status"`
	SentAt           time.Time  `json:"sent_at" db:"sent_at"`
	RespondedAt      *time.Time `json:"responded_at,omitempty" db:"responded_at"`
	Message          *string    `json:"message,omitempty" db:"message"`
}

// RoleApplication represents a role application
type RoleApplication struct {
	ID            int64                 `json:"id" db:"id"`
	UserID        int64                 `json:"user_id" db:"user_id"`
	ProjectRoleID int64                 `json:"project_role_id" db:"project_role_id"`
	Status        ApplicationStatusEnum `json:"status" db:"status"`
	AppliedAt     time.Time             `json:"applied_at" db:"applied_at"`
	UpdatedAt     time.Time             `json:"updated_at" db:"updated_at"`
	CoverLetter   *string               `json:"cover_letter,omitempty" db:"cover_letter"`
}
