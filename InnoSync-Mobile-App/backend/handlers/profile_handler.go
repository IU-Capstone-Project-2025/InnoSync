package handlers

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/team-18-project/InnoSync/backend/config"
	"github.com/team-18-project/InnoSync/backend/middleware"
	"github.com/team-18-project/InnoSync/backend/models"
	"github.com/team-18-project/InnoSync/backend/services"
)

// ProfileHandler handles profile-related endpoints
type ProfileHandler struct {
	profileService *services.ProfileService
	config         *config.Config
}

// NewProfileHandler creates a new profile handler
func NewProfileHandler(profileService *services.ProfileService, config *config.Config) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		config:         config,
	}
}

// CreateProfile handles profile creation
// @Summary Create user profile
// @Description Create a new user profile with detailed information
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateProfileRequest true "Profile creation request"
// @Success 201 {object} models.ProfileResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profile [post]
func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.profileService.CreateProfile(userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrProfileAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Profile already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create profile"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetProfile handles profile retrieval
// @Summary Get user profile
// @Description Get detailed profile information for the authenticated user
// @Tags profile
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.ProfileResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profile/me [get]
func (h *ProfileHandler) GetProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	response, err := h.profileService.GetProfile(userID)
	if err != nil {
		if errors.Is(err, services.ErrProfileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UpdateProfile handles profile updates
// @Summary Update user profile
// @Description Update user profile information
// @Tags profile
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.UpdateProfileRequest true "Profile update request"
// @Success 200 {object} models.ProfileResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profile [put]
func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.UpdateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.profileService.UpdateProfile(userID, &req)
	if err != nil {
		if errors.Is(err, services.ErrProfileNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// UploadProfileImage handles profile image upload
// @Summary Upload profile image
// @Description Upload a profile image for the authenticated user
// @Tags profile
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param image formData file true "Profile image file"
// @Success 200 {object} models.FileUploadResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profile/image [post]
func (h *ProfileHandler) UploadProfileImage(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return
	}
	defer file.Close()

	// Validate file
	if err := h.validateImageFile(header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save file
	filename, err := h.saveFile(file, header, "images")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Update profile image URL
	imageURL := fmt.Sprintf("/uploads/images/%s", filename)
	err = h.profileService.UpdateProfileImage(userID, imageURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile image"})
		return
	}

	response := &models.FileUploadResponse{
		URL:      imageURL,
		Filename: filename,
		Size:     header.Size,
		Message:  "Profile image uploaded successfully",
	}

	c.JSON(http.StatusOK, response)
}

// UploadResume handles resume upload
// @Summary Upload resume
// @Description Upload a resume file for the authenticated user
// @Tags profile
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param resume formData file true "Resume file"
// @Success 200 {object} models.FileUploadResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /profile/resume [post]
func (h *ProfileHandler) UploadResume(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	file, header, err := c.Request.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get file from request"})
		return
	}
	defer file.Close()

	// Validate file
	if err := h.validateResumeFile(header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save file
	filename, err := h.saveFile(file, header, "resumes")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	// Update resume URL
	resumeURL := fmt.Sprintf("/uploads/resumes/%s", filename)
	err = h.profileService.UpdateResumeURL(userID, resumeURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update resume URL"})
		return
	}

	response := &models.FileUploadResponse{
		URL:      resumeURL,
		Filename: filename,
		Size:     header.Size,
		Message:  "Resume uploaded successfully",
	}

	c.JSON(http.StatusOK, response)
}

// GetTechnologies handles technology list retrieval
// @Summary Get technologies
// @Description Get list of all available technologies
// @Tags profile
// @Produce json
// @Success 200 {array} models.Technology
// @Failure 500 {object} map[string]interface{}
// @Router /profile/technologies [get]
func (h *ProfileHandler) GetTechnologies(c *gin.Context) {
	technologies, err := h.profileService.GetTechnologies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get technologies"})
		return
	}

	c.JSON(http.StatusOK, technologies)
}

// validateImageFile validates image file upload
func (h *ProfileHandler) validateImageFile(header *multipart.FileHeader) error {
	// Check file size
	if header.Size > h.config.Upload.MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of %d bytes", h.config.Upload.MaxFileSize)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := []string{".jpg", ".jpeg", ".png", ".gif"}

	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return nil
		}
	}

	return fmt.Errorf("invalid file type. Allowed types: %s", strings.Join(allowedExts, ", "))
}

// validateResumeFile validates resume file upload
func (h *ProfileHandler) validateResumeFile(header *multipart.FileHeader) error {
	// Check file size
	if header.Size > h.config.Upload.MaxFileSize {
		return fmt.Errorf("file size exceeds maximum limit of %d bytes", h.config.Upload.MaxFileSize)
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowedExts := []string{".pdf", ".doc", ".docx"}

	for _, allowedExt := range allowedExts {
		if ext == allowedExt {
			return nil
		}
	}

	return fmt.Errorf("invalid file type. Allowed types: %s", strings.Join(allowedExts, ", "))
}

// saveFile saves uploaded file to disk
func (h *ProfileHandler) saveFile(file multipart.File, header *multipart.FileHeader, subdir string) (string, error) {
	// Create upload directory if it doesn't exist
	uploadDir := filepath.Join(h.config.Upload.UploadDir, subdir)
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create upload directory: %w", err)
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s-%s%s",
		strconv.FormatInt(time.Now().Unix(), 10),
		uuid.New().String()[:8],
		ext)

	// Create destination file
	destPath := filepath.Join(uploadDir, filename)
	destFile, err := os.Create(destPath)
	if err != nil {
		return "", fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy file content
	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", fmt.Errorf("failed to copy file: %w", err)
	}

	return filename, nil
}

// RegisterRoutes registers profile routes
func (h *ProfileHandler) RegisterRoutes(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	profile := r.Group("/profile")
	{
		profile.POST("", authMiddleware, h.CreateProfile)
		profile.GET("/me", authMiddleware, h.GetProfile)
		profile.PUT("", authMiddleware, h.UpdateProfile)
		profile.POST("/image", authMiddleware, h.UploadProfileImage)
		profile.POST("/resume", authMiddleware, h.UploadResume)
		profile.GET("/technologies", h.GetTechnologies)
	}
}