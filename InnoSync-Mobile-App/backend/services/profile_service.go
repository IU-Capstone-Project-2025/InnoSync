package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/team-18-project/InnoSync/backend/models"
)

var (
	ErrProfileNotFound      = errors.New("profile not found")
	ErrProfileAlreadyExists = errors.New("profile already exists")
	ErrTechnologyNotFound   = errors.New("technology not found")
)

// ProfileService handles profile-related operations
type ProfileService struct {
	db *sqlx.DB
}

// NewProfileService creates a new profile service
func NewProfileService(db *sqlx.DB) *ProfileService {
	return &ProfileService{
		db: db,
	}
}

// CreateProfile creates a new user profile
func (s *ProfileService) CreateProfile(userID int64, req *models.CreateProfileRequest) (*models.ProfileResponse, error) {
	// Check if profile already exists
	var existingProfile models.UserProfile
	err := s.db.Get(&existingProfile, "SELECT id FROM user_profile WHERE user_id = $1", userID)
	if err == nil {
		log.Printf("[ProfileService:CreateProfile] Profile already exists for user %d", userID)
		return nil, ErrProfileAlreadyExists
	}
	if err != sql.ErrNoRows {
		log.Printf("[ProfileService:CreateProfile] Failed to check existing profile: %v", err)
		return nil, fmt.Errorf("failed to check existing profile: %w", err)
	}

	// Get user information
	var user models.User
	err = s.db.Get(&user, "SELECT id, email, full_name FROM users WHERE id = $1", userID)
	if err != nil {
		log.Printf("[ProfileService:CreateProfile] Failed to get user: %v", err)
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	// Update user's full name if provided
	if req.Name != user.FullName {
		_, err = s.db.Exec("UPDATE users SET full_name = $1 WHERE id = $2", req.Name, userID)
		if err != nil {
			log.Printf("[ProfileService:CreateProfile] Failed to update user name: %v", err)
			return nil, fmt.Errorf("failed to update user name: %w", err)
		}
		user.FullName = req.Name
	}

	// Begin transaction
	tx, err := s.db.Beginx()
	if err != nil {
		log.Printf("[ProfileService:CreateProfile] Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Create user profile
	profile := &models.UserProfile{
		UserID:         userID,
		Telegram:       req.Telegram,
		GitHub:         req.GitHub,
		Bio:            req.Bio,
		Position:       req.Position,
		Education:      req.Education,
		Expertise:      req.Expertise,
		ExpertiseLevel: req.ExpertiseLevel,
	}

	query := `
		INSERT INTO user_profile (user_id, telegram, github, bio, position, education, expertise, expertise_level)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`
	err = tx.QueryRowx(query, profile.UserID, profile.Telegram, profile.GitHub, profile.Bio,
		profile.Position, profile.Education, profile.Expertise, profile.ExpertiseLevel).Scan(&profile.ID)
	if err != nil {
		log.Printf("[ProfileService:CreateProfile] Failed to create profile: %v", err)
		return nil, fmt.Errorf("failed to create profile: %w", err)
	}

	// Add work experiences
	var workExperiences []models.WorkExperience
	for _, we := range req.WorkExperience {
		workExp := models.WorkExperience{
			UserProfileID: profile.ID,
			Position:      we.Position,
			Company:       we.Company,
			Description:   we.Description,
		}

		// Parse dates
		startDate, err := time.Parse("2006-01-02", we.StartDate)
		if err != nil {
			return nil, fmt.Errorf("invalid start date format: %w", err)
		}
		workExp.StartDate = startDate

		if we.EndDate != nil {
			endDate, err := time.Parse("2006-01-02", *we.EndDate)
			if err != nil {
				return nil, fmt.Errorf("invalid end date format: %w", err)
			}
			workExp.EndDate = &endDate
		}

		query := `
			INSERT INTO work_experience (user_profile_id, start_date, end_date, position, company, description)
			VALUES ($1, $2, $3, $4, $5, $6)
			RETURNING id
		`
		err = tx.QueryRowx(query, workExp.UserProfileID, workExp.StartDate, workExp.EndDate,
			workExp.Position, workExp.Company, workExp.Description).Scan(&workExp.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to create work experience: %w", err)
		}
		workExperiences = append(workExperiences, workExp)
	}

	// Add technologies
	var technologies []models.Technology
	for _, techName := range req.Technologies {
		// Get or create technology
		var tech models.Technology
		err = tx.Get(&tech, "SELECT id, name, category FROM technology WHERE name = $1", techName)
		if err != nil {
			if err == sql.ErrNoRows {
				// Create new technology
				err = tx.QueryRowx("INSERT INTO technology (name) VALUES ($1) RETURNING id, name, category",
					techName).StructScan(&tech)
				if err != nil {
					return nil, fmt.Errorf("failed to create technology: %w", err)
				}
			} else {
				return nil, fmt.Errorf("failed to get technology: %w", err)
			}
		}

		// Link technology to profile
		_, err = tx.Exec("INSERT INTO user_profile_technology (user_profile_id, technology_id) VALUES ($1, $2)",
			profile.ID, tech.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to link technology to profile: %w", err)
		}
		technologies = append(technologies, tech)
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Create response
	response := &models.ProfileResponse{
		ID:              profile.ID,
		UserID:          userID,
		Name:            user.FullName,
		Email:           user.Email,
		Telegram:        profile.Telegram,
		GitHub:          profile.GitHub,
		Bio:             profile.Bio,
		Position:        profile.Position,
		Education:       profile.Education,
		Expertise:       profile.Expertise,
		ExpertiseLevel:  profile.ExpertiseLevel,
		WorkExperiences: workExperiences,
		Technologies:    technologies,
		CreatedAt:       time.Now(),
	}

	return response, nil
}

// GetProfile retrieves a user profile
func (s *ProfileService) GetProfile(userID int64) (*models.ProfileResponse, error) {
	// Get user and profile information
	var profile models.UserProfile
	var user models.User

	query := `
		SELECT
			up.id, up.user_id, up.telegram, up.github, up.bio, up.position,
			up.education, up.expertise, up.expertise_level, up.resume_url, up.avatar_url,
			u.full_name, u.email, u.created_at
		FROM user_profile up
		JOIN users u ON up.user_id = u.id
		WHERE up.user_id = $1
	`

	row := s.db.QueryRowx(query, userID)
	err := row.Scan(&profile.ID, &profile.UserID, &profile.Telegram, &profile.GitHub,
		&profile.Bio, &profile.Position, &profile.Education, &profile.Expertise,
		&profile.ExpertiseLevel, &profile.ResumeURL, &profile.AvatarURL,
		&user.FullName, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[ProfileService:GetProfile] Profile not found for user %d", userID)
			return nil, ErrProfileNotFound
		}
		log.Printf("[ProfileService:GetProfile] Failed to get profile: %v", err)
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	// Get work experiences
	workExperiences := []models.WorkExperience{}
	err = s.db.Select(&workExperiences,
		"SELECT id, user_profile_id, start_date, end_date, position, company, description FROM work_experience WHERE user_profile_id = $1 ORDER BY start_date DESC",
		profile.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get work experiences: %w", err)
	}

	// Get technologies
	technologies := []models.Technology{}
	query = `
		SELECT t.id, t.name, t.category
		FROM technology t
		JOIN user_profile_technology upt ON t.id = upt.technology_id
		WHERE upt.user_profile_id = $1
		ORDER BY t.name
	`
	err = s.db.Select(&technologies, query, profile.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get technologies: %w", err)
	}

	// Create response
	response := &models.ProfileResponse{
		ID:              profile.ID,
		UserID:          profile.UserID,
		Name:            user.FullName,
		Email:           user.Email,
		Telegram:        profile.Telegram,
		GitHub:          profile.GitHub,
		Bio:             profile.Bio,
		Position:        profile.Position,
		Education:       profile.Education,
		Expertise:       profile.Expertise,
		ExpertiseLevel:  profile.ExpertiseLevel,
		ResumeURL:       profile.ResumeURL,
		AvatarURL:       profile.AvatarURL,
		WorkExperiences: workExperiences,
		Technologies:    technologies,
		CreatedAt:       user.CreatedAt,
	}

	return response, nil
}

// UpdateProfile updates a user profile
func (s *ProfileService) UpdateProfile(userID int64, req *models.UpdateProfileRequest) (*models.ProfileResponse, error) {
	// Check if profile exists
	var profileID int64
	err := s.db.Get(&profileID, "SELECT id FROM user_profile WHERE user_id = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[ProfileService:UpdateProfile] Profile not found for user %d", userID)
			return nil, ErrProfileNotFound
		}
		log.Printf("[ProfileService:UpdateProfile] Failed to get profile: %v", err)
		return nil, fmt.Errorf("failed to get profile: %w", err)
	}

	// Begin transaction
	tx, err := s.db.Beginx()
	if err != nil {
		log.Printf("[ProfileService:UpdateProfile] Failed to begin transaction: %v", err)
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Update profile
	updateFields := []string{}
	updateValues := []interface{}{}
	valueIndex := 1

	if req.Telegram != nil {
		updateFields = append(updateFields, fmt.Sprintf("telegram = $%d", valueIndex))
		updateValues = append(updateValues, req.Telegram)
		valueIndex++
	}
	if req.GitHub != nil {
		updateFields = append(updateFields, fmt.Sprintf("github = $%d", valueIndex))
		updateValues = append(updateValues, req.GitHub)
		valueIndex++
	}
	if req.Bio != nil {
		updateFields = append(updateFields, fmt.Sprintf("bio = $%d", valueIndex))
		updateValues = append(updateValues, req.Bio)
		valueIndex++
	}
	if req.Position != nil {
		updateFields = append(updateFields, fmt.Sprintf("position = $%d", valueIndex))
		updateValues = append(updateValues, req.Position)
		valueIndex++
	}
	if req.Education != nil {
		updateFields = append(updateFields, fmt.Sprintf("education = $%d", valueIndex))
		updateValues = append(updateValues, req.Education)
		valueIndex++
	}
	if req.Expertise != nil {
		updateFields = append(updateFields, fmt.Sprintf("expertise = $%d", valueIndex))
		updateValues = append(updateValues, req.Expertise)
		valueIndex++
	}
	if req.ExpertiseLevel != nil {
		updateFields = append(updateFields, fmt.Sprintf("expertise_level = $%d", valueIndex))
		updateValues = append(updateValues, req.ExpertiseLevel)
		valueIndex++
	}

	if len(updateFields) > 0 {
		query := fmt.Sprintf("UPDATE user_profile SET %s WHERE id = $%d",
			strings.Join(updateFields, ", "), valueIndex)
		updateValues = append(updateValues, profileID)
		_, err = tx.Exec(query, updateValues...)
		if err != nil {
			log.Printf("[ProfileService:UpdateProfile] Failed to update profile: %v", err)
			return nil, fmt.Errorf("failed to update profile: %w", err)
		}
	}

	// Update work experiences if provided
	if req.WorkExperience != nil {
		// Delete existing work experiences
		_, err = tx.Exec("DELETE FROM work_experience WHERE user_profile_id = $1", profileID)
		if err != nil {
			log.Printf("[ProfileService:UpdateProfile] Failed to delete work experiences: %v", err)
			return nil, fmt.Errorf("failed to delete work experiences: %w", err)
		}

		// Add new work experiences
		for _, we := range req.WorkExperience {
			workExp := models.WorkExperience{
				UserProfileID: profileID,
				Position:      we.Position,
				Company:       we.Company,
				Description:   we.Description,
			}

			// Parse dates
			startDate, err := time.Parse("2006-01-02", we.StartDate)
			if err != nil {
				log.Printf("[ProfileService:UpdateProfile] Invalid start date format: %v", err)
				return nil, fmt.Errorf("invalid start date format: %w", err)
			}
			workExp.StartDate = startDate

			if we.EndDate != nil {
				endDate, err := time.Parse("2006-01-02", *we.EndDate)
				if err != nil {
					log.Printf("[ProfileService:UpdateProfile] Invalid end date format: %v", err)
					return nil, fmt.Errorf("invalid end date format: %w", err)
				}
				workExp.EndDate = &endDate
			}

			query := `
				INSERT INTO work_experience (user_profile_id, start_date, end_date, position, company, description)
				VALUES ($1, $2, $3, $4, $5, $6)
			`
			_, err = tx.Exec(query, workExp.UserProfileID, workExp.StartDate, workExp.EndDate,
				workExp.Position, workExp.Company, workExp.Description)
			if err != nil {
				log.Printf("[ProfileService:UpdateProfile] Failed to create work experience: %v", err)
				return nil, fmt.Errorf("failed to create work experience: %w", err)
			}
		}
	}

	// Update technologies if provided
	if req.Technologies != nil {
		// Delete existing technology associations
		_, err = tx.Exec("DELETE FROM user_profile_technology WHERE user_profile_id = $1", profileID)
		if err != nil {
			log.Printf("[ProfileService:UpdateProfile] Failed to delete technology associations: %v", err)
			return nil, fmt.Errorf("failed to delete technology associations: %w", err)
		}

		// Add new technologies
		for _, techName := range req.Technologies {
			// Get or create technology
			var tech models.Technology
			err = tx.Get(&tech, "SELECT id, name, category FROM technology WHERE name = $1", techName)
			if err != nil {
				if err == sql.ErrNoRows {
					// Create new technology
					err = tx.QueryRowx("INSERT INTO technology (name) VALUES ($1) RETURNING id, name, category",
						techName).StructScan(&tech)
					if err != nil {
						log.Printf("[ProfileService:UpdateProfile] Failed to create technology: %v", err)
						return nil, fmt.Errorf("failed to create technology: %w", err)
					}
				} else {
					log.Printf("[ProfileService:UpdateProfile] Failed to get technology: %v", err)
					return nil, fmt.Errorf("failed to get technology: %w", err)
				}
			}

			// Link technology to profile
			_, err = tx.Exec("INSERT INTO user_profile_technology (user_profile_id, technology_id) VALUES ($1, $2)",
				profileID, tech.ID)
			if err != nil {
				log.Printf("[ProfileService:UpdateProfile] Failed to link technology to profile: %v", err)
				return nil, fmt.Errorf("failed to link technology to profile: %w", err)
			}
		}
	}

	// Commit transaction
	if err = tx.Commit(); err != nil {
		log.Printf("[ProfileService:UpdateProfile] Failed to commit transaction: %v", err)
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	// Get updated profile
	return s.GetProfile(userID)
}

// UpdateProfileImage updates profile image URL
func (s *ProfileService) UpdateProfileImage(userID int64, imageURL string) error {
	_, err := s.db.Exec("UPDATE user_profile SET avatar_url = $1 WHERE user_id = $2", imageURL, userID)
	if err != nil {
		log.Printf("[ProfileService:UpdateProfileImage] Failed to update profile image: %v", err)
		return fmt.Errorf("failed to update profile image: %w", err)
	}
	return nil
}

// UpdateResumeURL updates resume URL
func (s *ProfileService) UpdateResumeURL(userID int64, resumeURL string) error {
	_, err := s.db.Exec("UPDATE user_profile SET resume_url = $1 WHERE user_id = $2", resumeURL, userID)
	if err != nil {
		log.Printf("[ProfileService:UpdateResumeURL] Failed to update resume URL: %v", err)
		return fmt.Errorf("failed to update resume URL: %w", err)
	}
	return nil
}

// GetTechnologies retrieves all available technologies
func (s *ProfileService) GetTechnologies() ([]models.Technology, error) {
	var technologies []models.Technology
	err := s.db.Select(&technologies, "SELECT id, name, category FROM technology ORDER BY name")
	if err != nil {
		log.Printf("[ProfileService:GetTechnologies] Failed to get technologies: %v", err)
		return nil, fmt.Errorf("failed to get technologies: %w", err)
	}
	return technologies, nil
}
