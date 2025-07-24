package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team-18-project/InnoSync/backend/middleware"
	"github.com/team-18-project/InnoSync/backend/models"
	"github.com/team-18-project/InnoSync/backend/services"
)

// ApplicationHandler handles application-related endpoints
type ApplicationHandler struct {
	applicationService *services.ApplicationService
}

// NewApplicationHandler creates a new application handler
func NewApplicationHandler(applicationService *services.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{
		applicationService: applicationService,
	}
}

// CreateApplication handles role application creation
// @Summary Create application
// @Description Apply for a project role
// @Tags applications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param projectRoleId path int true "Project Role ID"
// @Param request body models.CreateApplicationRequest true "Application creation request"
// @Success 201 {object} models.ApplicationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications/project-roles/{projectRoleId} [post]
func (h *ApplicationHandler) CreateApplication(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projectRoleIDStr := c.Param("projectRoleId")
	projectRoleID, err := strconv.ParseInt(projectRoleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project role ID"})
		return
	}

	var req models.CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.applicationService.CreateApplication(userID, projectRoleID, &req)
	if err != nil {
		if errors.Is(err, services.ErrApplicationAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Application already exists"})
			return
		}
		if errors.Is(err, services.ErrProjectRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project role not found"})
			return
		}
		if errors.Is(err, services.ErrProjectRoleNotOpen) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Project role is not open for applications"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create application"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUserApplications handles retrieving user's applications
// @Summary Get user applications
// @Description Get all applications submitted by the authenticated user
// @Tags applications
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.ApplicationResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications [get]
func (h *ApplicationHandler) GetUserApplications(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	applications, err := h.applicationService.GetUserApplications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user applications"})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GetRoleApplications handles retrieving applications for a specific role
// @Summary Get role applications
// @Description Get all applications for a specific project role
// @Tags applications
// @Produce json
// @Security BearerAuth
// @Param projectRoleId path int true "Project Role ID"
// @Success 200 {array} models.ApplicationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications/project-roles/{projectRoleId} [get]
func (h *ApplicationHandler) GetRoleApplications(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projectRoleIDStr := c.Param("projectRoleId")
	projectRoleID, err := strconv.ParseInt(projectRoleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project role ID"})
		return
	}

	applications, err := h.applicationService.GetRoleApplications(userID, projectRoleID)
	if err != nil {
		if errors.Is(err, services.ErrProjectRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project role not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedApplication) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to applications"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get role applications"})
		return
	}

	c.JSON(http.StatusOK, applications)
}

// GetApplication handles retrieving a specific application
// @Summary Get application
// @Description Get a specific application by ID
// @Tags applications
// @Produce json
// @Security BearerAuth
// @Param id path int true "Application ID"
// @Success 200 {object} models.ApplicationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications/{id} [get]
func (h *ApplicationHandler) GetApplication(c *gin.Context) {
	applicationIDStr := c.Param("id")
	applicationID, err := strconv.ParseInt(applicationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	application, err := h.applicationService.GetApplication(applicationID)
	if err != nil {
		if errors.Is(err, services.ErrApplicationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get application"})
		return
	}

	c.JSON(http.StatusOK, application)
}

// UpdateApplicationStatus handles updating application status
// @Summary Update application status
// @Description Update the status of a specific application
// @Tags applications
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Application ID"
// @Param request body models.UpdateApplicationStatusRequest true "Status update request"
// @Success 200 {object} models.ApplicationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications/{id}/status [patch]
func (h *ApplicationHandler) UpdateApplicationStatus(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	applicationIDStr := c.Param("id")
	applicationID, err := strconv.ParseInt(applicationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	var req models.UpdateApplicationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.applicationService.UpdateApplicationStatus(userID, applicationID, &req)
	if err != nil {
		if errors.Is(err, services.ErrApplicationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedApplication) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to application"})
			return
		}
		if errors.Is(err, services.ErrInvalidApplicationStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application status transition"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update application status"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteApplication handles application deletion (withdrawal)
// @Summary Delete application
// @Description Withdraw a submitted application
// @Tags applications
// @Security BearerAuth
// @Param id path int true "Application ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /applications/{id} [delete]
func (h *ApplicationHandler) DeleteApplication(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	applicationIDStr := c.Param("id")
	applicationID, err := strconv.ParseInt(applicationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid application ID"})
		return
	}

	err = h.applicationService.DeleteApplication(userID, applicationID)
	if err != nil {
		if errors.Is(err, services.ErrApplicationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Application not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedApplication) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to application"})
			return
		}
		if errors.Is(err, services.ErrInvalidApplicationStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot withdraw application in current status"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete application"})
		return
	}

	c.Status(http.StatusNoContent)
}

// RegisterRoutes registers application routes
func (h *ApplicationHandler) RegisterRoutes(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	applications := r.Group("/applications")
	{
		applications.POST("/project-roles/:projectRoleId", authMiddleware, h.CreateApplication)
		applications.GET("", authMiddleware, h.GetUserApplications)
		applications.GET("/project-roles/:projectRoleId", authMiddleware, h.GetRoleApplications)
		applications.GET("/:id", authMiddleware, h.GetApplication)
		applications.PATCH("/:id/status", authMiddleware, h.UpdateApplicationStatus)
		applications.DELETE("/:id", authMiddleware, h.DeleteApplication)
	}
}