package handlers

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/team-18-project/InnoSync/backend/middleware"
	"github.com/team-18-project/InnoSync/backend/models"
	"github.com/team-18-project/InnoSync/backend/services"
)

// InvitationHandler handles invitation-related endpoints
type InvitationHandler struct {
	invitationService *services.InvitationService
}

// NewInvitationHandler creates a new invitation handler
func NewInvitationHandler(invitationService *services.InvitationService) *InvitationHandler {
	return &InvitationHandler{
		invitationService: invitationService,
	}
}

// CreateInvitation handles invitation creation
// @Summary Create invitation
// @Description Create a new invitation for a project role
// @Tags invitations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateInvitationRequest true "Invitation creation request"
// @Success 201 {object} models.InvitationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 409 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations [post]
func (h *InvitationHandler) CreateInvitation(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var req models.CreateInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.invitationService.CreateInvitation(userID, &req)
	if err != nil {
		log.Printf("CreateInvitation error: %v", err)
		if errors.Is(err, services.ErrInvitationAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "Invitation already exists"})
			return
		}
		if errors.Is(err, services.ErrProjectRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project role not found"})
			return
		}
		if errors.Is(err, services.ErrRecipientNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Recipient not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedInvitation) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to create invitation"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invitation"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// GetSentInvitations handles retrieving sent invitations
// @Summary Get sent invitations
// @Description Get all invitations sent by the authenticated user
// @Tags invitations
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.InvitationResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/sent [get]
func (h *InvitationHandler) GetSentInvitations(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	invitations, err := h.invitationService.GetSentInvitations(userID)
	if err != nil {
		log.Printf("GetSentInvitations error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get sent invitations"})
		return
	}

	c.JSON(http.StatusOK, invitations)
}

// GetReceivedInvitations handles retrieving received invitations
// @Summary Get received invitations
// @Description Get all invitations received by the authenticated user
// @Tags invitations
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.InvitationResponse
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/received [get]
func (h *InvitationHandler) GetReceivedInvitations(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	invitations, err := h.invitationService.GetReceivedInvitations(userID)
	if err != nil {
		log.Printf("GetReceivedInvitations error: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get received invitations"})
		return
	}

	c.JSON(http.StatusOK, invitations)
}

// GetInvitation handles retrieving a specific invitation
// @Summary Get invitation
// @Description Get a specific invitation by ID
// @Tags invitations
// @Produce json
// @Security BearerAuth
// @Param id path int true "Invitation ID"
// @Success 200 {object} models.InvitationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/{id} [get]
func (h *InvitationHandler) GetInvitation(c *gin.Context) {
	invitationIDStr := c.Param("id")
	invitationID, err := strconv.ParseInt(invitationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invitation ID"})
		return
	}

	invitation, err := h.invitationService.GetInvitation(invitationID)
	if err != nil {
		log.Printf("GetInvitation error: %v", err)
		if errors.Is(err, services.ErrInvitationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get invitation"})
		return
	}

	c.JSON(http.StatusOK, invitation)
}

// RespondToInvitation handles responding to an invitation
// @Summary Respond to invitation
// @Description Respond to an invitation (accept or decline)
// @Tags invitations
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Invitation ID"
// @Param request body models.RespondInvitationRequest true "Invitation response request"
// @Success 200 {object} models.InvitationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/{id}/respond [patch]
func (h *InvitationHandler) RespondToInvitation(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	invitationIDStr := c.Param("id")
	invitationID, err := strconv.ParseInt(invitationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invitation ID"})
		return
	}

	var req models.RespondInvitationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.invitationService.RespondToInvitation(userID, invitationID, &req)
	if err != nil {
		log.Printf("RespondToInvitation error: %v", err)
		if errors.Is(err, services.ErrInvitationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedInvitation) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to respond to invitation"})
			return
		}
		if errors.Is(err, services.ErrInvitationAlreadyReplied) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invitation already responded to"})
			return
		}
		if errors.Is(err, services.ErrInvalidInvitationStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invitation response"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to respond to invitation"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// RevokeInvitation handles revoking an invitation
// @Summary Revoke invitation
// @Description Revoke a sent invitation
// @Tags invitations
// @Security BearerAuth
// @Param id path int true "Invitation ID"
// @Success 200 {object} models.InvitationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/{id}/revoke [patch]
func (h *InvitationHandler) RevokeInvitation(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	invitationIDStr := c.Param("id")
	invitationID, err := strconv.ParseInt(invitationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invitation ID"})
		return
	}

	response, err := h.invitationService.RevokeInvitation(userID, invitationID)
	if err != nil {
		log.Printf("RevokeInvitation error: %v", err)
		if errors.Is(err, services.ErrInvitationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedInvitation) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to revoke invitation"})
			return
		}
		if errors.Is(err, services.ErrInvalidInvitationStatus) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot revoke invitation in current status"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to revoke invitation"})
		return
	}

	c.JSON(http.StatusOK, response)
}

// DeleteInvitation handles invitation deletion
// @Summary Delete invitation
// @Description Delete a sent invitation
// @Tags invitations
// @Security BearerAuth
// @Param id path int true "Invitation ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/{id} [delete]
func (h *InvitationHandler) DeleteInvitation(c *gin.Context) {
	userID := middleware.GetUserID(c)
	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	invitationIDStr := c.Param("id")
	invitationID, err := strconv.ParseInt(invitationIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invitation ID"})
		return
	}

	err = h.invitationService.DeleteInvitation(userID, invitationID)
	if err != nil {
		log.Printf("DeleteInvitation error: %v", err)
		if errors.Is(err, services.ErrInvitationNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Invitation not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedInvitation) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized to delete invitation"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invitation"})
		return
	}

	c.Status(http.StatusNoContent)
}

// GetInvitationsByProjectRole handles retrieving invitations for a specific project role
// @Summary Get invitations by project role
// @Description Get all invitations for a specific project role
// @Tags invitations
// @Produce json
// @Security BearerAuth
// @Param projectRoleId path int true "Project Role ID"
// @Success 200 {array} models.InvitationResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /invitations/project-roles/{projectRoleId} [get]
func (h *InvitationHandler) GetInvitationsByProjectRole(c *gin.Context) {
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

	invitations, err := h.invitationService.GetInvitationsByProjectRole(userID, projectRoleID)
	if err != nil {
		log.Printf("GetInvitationsByProjectRole error: %v", err)
		if errors.Is(err, services.ErrProjectRoleNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Project role not found"})
			return
		}
		if errors.Is(err, services.ErrUnauthorizedInvitation) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access to invitations"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get invitations"})
		return
	}

	c.JSON(http.StatusOK, invitations)
}

// RegisterRoutes registers invitation routes
func (h *InvitationHandler) RegisterRoutes(r *gin.RouterGroup, authMiddleware gin.HandlerFunc) {
	invitations := r.Group("/invitations")
	{
		invitations.POST("", authMiddleware, h.CreateInvitation)
		invitations.GET("/sent", authMiddleware, h.GetSentInvitations)
		invitations.GET("/received", authMiddleware, h.GetReceivedInvitations)
		invitations.GET("/:id", authMiddleware, h.GetInvitation)
		invitations.PATCH("/:id/respond", authMiddleware, h.RespondToInvitation)
		invitations.PATCH("/:id/revoke", authMiddleware, h.RevokeInvitation)
		invitations.DELETE("/:id", authMiddleware, h.DeleteInvitation)
		invitations.GET("/project-roles/:projectRoleId", authMiddleware, h.GetInvitationsByProjectRole)
	}
}
