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

// ProjectHandler handles project-related endpoints
type ProjectHandler struct {
	projectService *services.ProjectService
}

// NewProjectHandler creates a new project handler
func NewProjectHandler(projectService *services.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		projectService: projectService,
	}
}

// CreateProject handles project creation
// @Summary Create project
// @Description Create a new project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateProjectRequest true "Project creation request"
// @Success 201 {object} models.ProjectResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects [post]
func (h *ProjectHandler) CreateProject(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.CreateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.projectService.CreateProject(userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetUserProjects handles retrieving user's projects
// @Summary Get user projects
// @Description Get all projects created by the authenticated user
// @Tags projects
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.ProjectResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/me [get]
func (h *ProjectHandler) GetUserProjects(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projects, err := h.projectService.GetUserProjects(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get projects"})
		return
	}

	c.JSON(http.StatusOK, projects)
}

// GetProject handles retrieving a specific project
// @Summary Get project
// @Description Get a specific project by ID
// @Tags projects
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Success 200 {object} models.ProjectResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id} [get]
func (h *ProjectHandler) GetProject(c *gin.Context) {
	projectIDStr := c.Param("id")
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	project, err := h.projectService.GetProject(projectID)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project"})
		return
	}

	c.JSON(http.StatusOK, project)
}

// UpdateProject handles project updates
// @Summary Update project
// @Description Update a project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Param request body models.UpdateProjectRequest true "Project update request"
// @Success 200 {object} models.ProjectResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id} [put]
func (h *ProjectHandler) UpdateProject(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projectIDStr := c.Param("id")
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req models.UpdateProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.projectService.UpdateProject(userID, projectID, &req)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedProject) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to project"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update project"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteProject handles project deletion
// @Summary Delete project
// @Description Delete a project (soft delete)
// @Tags projects
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id} [delete]
func (h *ProjectHandler) DeleteProject(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projectIDStr := c.Param("id")
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	err = h.projectService.DeleteProject(userID, projectID)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedProject) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to project"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete project"})
		return
	}

	c.Status(http.StatusNoContent)
}

// CreateProjectRole handles project role creation
// @Summary Create project role
// @Description Create a new role for a project
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Param request body models.CreateProjectRoleRequest true "Project role creation request"
// @Success 201 {object} models.ProjectRoleResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id}/roles [post]
func (h *ProjectHandler) CreateProjectRole(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	projectIDStr := c.Param("id")
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	var req models.CreateProjectRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.projectService.CreateProjectRole(userID, projectID, &req)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedProject) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to project"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project role"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetProjectRoles handles retrieving project roles
// @Summary Get project roles
// @Description Get all roles for a specific project
// @Tags projects
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Success 200 {array} models.ProjectRoleResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id}/roles [get]
func (h *ProjectHandler) GetProjectRoles(c *gin.Context) {
	projectIDStr := c.Param("id")
	projectID, err := strconv.ParseInt(projectIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid project ID"})
		return
	}

	roles, err := h.projectService.GetProjectRolesResponse(projectID)
	if err != nil {
		if errors.Is(err, services.ErrProjectNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project roles"})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// GetProjectRole handles retrieving a specific project role
// @Summary Get project role
// @Description Get a specific project role by ID
// @Tags projects
// @Produce json
// @Security BearerAuth
// @Param id path int true "Project ID"
// @Param roleId path int true "Role ID"
// @Success 200 {object} models.ProjectRoleResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /projects/{id}/roles/{roleId} [get]
func (h *ProjectHandler) GetProjectRole(c *gin.Context) {
	roleIDStr := c.Param("roleId")
	roleID, err := strconv.ParseInt(roleIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.projectService.GetProjectRole(roleID)
	if err != nil {
		if errors.Is(err, services.ErrProjectRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project role not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get project role"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// RegisterRoutes registers project routes
func (h *ProjectHandler) RegisterRoutes(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	projects := r.Group("/projects")
	{
		projects.POST("", authMiddleware, h.CreateProject)
		projects.GET("/me", authMiddleware, h.GetUserProjects)
		projects.GET("/:id", authMiddleware, h.GetProject)
		projects.PUT("/:id", authMiddleware, h.UpdateProject)
		projects.DELETE("/:id", authMiddleware, h.DeleteProject)
		projects.POST("/:id/roles", authMiddleware, h.CreateProjectRole)
		projects.GET("/:id/roles", authMiddleware, h.GetProjectRoles)
		projects.GET("/:id/roles/:roleId", authMiddleware, h.GetProjectRole)
	}
}