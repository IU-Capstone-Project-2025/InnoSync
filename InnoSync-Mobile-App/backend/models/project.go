package models

import "time"

// Project represents a project in the system
type Project struct {
	ID          int64     `json:"id" db:"id"`
	RecruiterID int64     `json:"recruiter_id" db:"recruiter_id"`
	Title       string    `json:"title" db:"title"`
	Description *string   `json:"description,omitempty" db:"description"`
	TeamSize    int       `json:"team_size" db:"team_size"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	IsActive    bool      `json:"is_active" db:"is_active"`

	// Related data
	Roles []ProjectRole `json:"roles,omitempty" db:"-"`
}

// ProjectRole represents a role within a project
type ProjectRole struct {
	ID             int64               `json:"id" db:"id"`
	ProjectID      int64               `json:"project_id" db:"project_id"`
	RoleName       string              `json:"role_name" db:"role_name"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty" db:"expertise_level"`
	Technologies   *string             `json:"technologies,omitempty" db:"technologies"`
	IsOpen         bool                `json:"is_open" db:"is_open"`

	// Related data
	Project *Project `json:"project,omitempty" db:"-"`
}

// Project request/response models
type CreateProjectRequest struct {
	Title       string  `json:"title" binding:"required"`
	Description *string `json:"description,omitempty"`
	TeamSize    int     `json:"team_size,omitempty"`
}

type ProjectResponse struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Description *string       `json:"description,omitempty"`
	TeamSize    int           `json:"team_size"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	IsActive    bool          `json:"is_active"`
	Roles       []ProjectRole `json:"roles,omitempty"`
}

type UpdateProjectRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	TeamSize    *int    `json:"team_size,omitempty"`
}

// Project role request/response models
type CreateProjectRoleRequest struct {
	RoleName       string              `json:"role_name" binding:"required"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty"`
	Technologies   []string            `json:"technologies,omitempty"`
}

type ProjectRoleResponse struct {
	ID             int64               `json:"id"`
	ProjectID      int64               `json:"project_id"`
	RoleName       string              `json:"role_name"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty"`
	Technologies   []string            `json:"technologies,omitempty"`
	IsOpen         bool                `json:"is_open"`
	ProjectTitle   string              `json:"project_title,omitempty"`
}

type UpdateProjectRoleRequest struct {
	RoleName       *string             `json:"role_name,omitempty"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty"`
	Technologies   []string            `json:"technologies,omitempty"`
	IsOpen         *bool               `json:"is_open,omitempty"`
}

// Application response models
type ApplicationResponse struct {
	ID            int64                 `json:"id"`
	UserID        int64                 `json:"user_id"`
	UserFullName  string                `json:"user_full_name"`
	ProjectRoleID int64                 `json:"project_role_id"`
	RoleName      string                `json:"role_name"`
	ProjectID     int64                 `json:"project_id"`
	ProjectTitle  string                `json:"project_title"`
	Status        ApplicationStatusEnum `json:"status"`
	AppliedAt     time.Time             `json:"applied_at"`
	UpdatedAt     time.Time             `json:"updated_at"`
	CoverLetter   *string               `json:"cover_letter,omitempty"`
}

type CreateApplicationRequest struct {
	CoverLetter *string `json:"cover_letter,omitempty"`
}

type UpdateApplicationStatusRequest struct {
	Status ApplicationStatusEnum `json:"status" binding:"required"`
}

// Invitation response models
type InvitationResponse struct {
	ID               int64      `json:"id" db:"id"`
	ProjectRoleID    int64      `json:"project_role_id" db:"project_role_id"`
	RoleName         string     `json:"role_name" db:"role_name"`
	ProjectID        int64      `json:"project_id" db:"project_id"`
	ProjectTitle     string     `json:"project_title" db:"project_title"`
	RecipientID      int64      `json:"recipient_id" db:"recipient_id"`
	RecipientName    string     `json:"recipient_name" db:"recipient_name"`
	InvitationStatus string     `json:"invitation_status" db:"invitation_status"`
	SentAt           time.Time  `json:"sent_at" db:"sent_at"`
	RespondedAt      *time.Time `json:"responded_at,omitempty" db:"responded_at"`
	Message          *string    `json:"message,omitempty" db:"message"`
}

type CreateInvitationRequest struct {
	ProjectRoleID int64   `json:"project_role_id" binding:"required"`
	RecipientID   int64   `json:"recipient_id" binding:"required"`
	Message       *string `json:"message,omitempty"`
}

type RespondInvitationRequest struct {
	Response string `json:"response" binding:"required"`
}

// Helper methods
func (p *Project) IsValid() bool {
	return p.Title != "" && p.RecruiterID > 0
}

func (pr *ProjectRole) IsValid() bool {
	return pr.RoleName != "" && pr.ProjectID > 0
}
