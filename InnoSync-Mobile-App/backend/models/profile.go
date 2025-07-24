package models

import (
	"time"
)

// UserProfile represents a user's profile information
type UserProfile struct {
	ID             int64               `json:"id" db:"id"`
	UserID         int64               `json:"user_id" db:"user_id"`
	Telegram       *string             `json:"telegram,omitempty" db:"telegram"`
	GitHub         *string             `json:"github,omitempty" db:"github"`
	Bio            *string             `json:"bio,omitempty" db:"bio"`
	Position       *string             `json:"position,omitempty" db:"position"`
	Education      *EducationEnum      `json:"education,omitempty" db:"education"`
	Expertise      *string             `json:"expertise,omitempty" db:"expertise"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty" db:"expertise_level"`
	ResumeURL      *string             `json:"resume_url,omitempty" db:"resume_url"`
	AvatarURL      *string             `json:"avatar_url,omitempty" db:"avatar_url"`

	// Related data
	WorkExperiences []WorkExperience `json:"work_experiences,omitempty" db:"-"`
	Technologies    []Technology     `json:"technologies,omitempty" db:"-"`
	User            *User            `json:"user,omitempty" db:"-"`
}

// WorkExperience represents a user's work experience
type WorkExperience struct {
	ID            int64      `json:"id" db:"id"`
	UserProfileID int64      `json:"user_profile_id" db:"user_profile_id"`
	StartDate     time.Time  `json:"start_date" db:"start_date"`
	EndDate       *time.Time `json:"end_date,omitempty" db:"end_date"`
	Position      string     `json:"position" db:"position"`
	Company       string     `json:"company" db:"company"`
	Description   *string    `json:"description,omitempty" db:"description"`
}

// Technology represents a technology/skill
type Technology struct {
	ID       int64   `json:"id" db:"id"`
	Name     string  `json:"name" db:"name"`
	Category *string `json:"category,omitempty" db:"category"`
}

// UserProfileTechnology represents the many-to-many relationship between user profiles and technologies
type UserProfileTechnology struct {
	UserProfileID int64 `json:"user_profile_id" db:"user_profile_id"`
	TechnologyID  int64 `json:"technology_id" db:"technology_id"`
}

// Profile request/response models
type CreateProfileRequest struct {
	Name           string                    `json:"name" binding:"required"`
	Email          string                    `json:"email" binding:"required,email"`
	Telegram       *string                   `json:"telegram,omitempty"`
	GitHub         *string                   `json:"github,omitempty"`
	Bio            *string                   `json:"bio,omitempty"`
	Position       *string                   `json:"position,omitempty"`
	Education      *EducationEnum            `json:"education,omitempty"`
	Expertise      *string                   `json:"expertise,omitempty"`
	ExpertiseLevel *ExpertiseLevelEnum       `json:"expertise_level,omitempty"`
	Technologies   []string                  `json:"technologies,omitempty"`
	WorkExperience []WorkExperienceRequest   `json:"work_experience,omitempty"`
}

type WorkExperienceRequest struct {
	StartDate   string  `json:"start_date" binding:"required"`
	EndDate     *string `json:"end_date,omitempty"`
	Position    string  `json:"position" binding:"required"`
	Company     string  `json:"company" binding:"required"`
	Description *string `json:"description,omitempty"`
}

type UpdateProfileRequest struct {
	Telegram       *string                   `json:"telegram,omitempty"`
	GitHub         *string                   `json:"github,omitempty"`
	Bio            *string                   `json:"bio,omitempty"`
	Position       *string                   `json:"position,omitempty"`
	Education      *EducationEnum            `json:"education,omitempty"`
	Expertise      *string                   `json:"expertise,omitempty"`
	ExpertiseLevel *ExpertiseLevelEnum       `json:"expertise_level,omitempty"`
	Technologies   []string                  `json:"technologies,omitempty"`
	WorkExperience []WorkExperienceRequest   `json:"work_experience,omitempty"`
}

type ProfileResponse struct {
	ID             int64               `json:"id"`
	UserID         int64               `json:"user_id"`
	Name           string              `json:"name"`
	Email          string              `json:"email"`
	Telegram       *string             `json:"telegram,omitempty"`
	GitHub         *string             `json:"github,omitempty"`
	Bio            *string             `json:"bio,omitempty"`
	Position       *string             `json:"position,omitempty"`
	Education      *EducationEnum      `json:"education,omitempty"`
	Expertise      *string             `json:"expertise,omitempty"`
	ExpertiseLevel *ExpertiseLevelEnum `json:"expertise_level,omitempty"`
	ResumeURL      *string             `json:"resume_url,omitempty"`
	AvatarURL      *string             `json:"avatar_url,omitempty"`
	WorkExperiences []WorkExperience   `json:"work_experiences,omitempty"`
	Technologies    []Technology       `json:"technologies,omitempty"`
	CreatedAt      time.Time           `json:"created_at"`
}

type FileUploadResponse struct {
	URL      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Message  string `json:"message"`
}

// Helper methods
func (p *UserProfile) IsValid() bool {
	return p.UserID > 0
}

func (we *WorkExperience) IsValid() bool {
	return we.Position != "" && we.Company != "" && we.UserProfileID > 0
}

func (t *Technology) IsValid() bool {
	return t.Name != ""
}
