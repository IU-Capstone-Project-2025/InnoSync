package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/team-18-project/InnoSync/backend/utils"
)

// AuthMiddleware creates a middleware for JWT authentication
func AuthMiddleware(tokenManager *utils.TokenManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Check if the header starts with "Bearer "
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := tokenManager.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("user_id", claims.UserID)
		c.Set("user_email", claims.Email)
		c.Set("user_full_name", claims.FullName)
		c.Set("claims", claims)

		c.Next()
	}
}

// GetUserID extracts user ID from context
func GetUserID(c *gin.Context) int64 {
	userID, exists := c.Get("user_id")
	if !exists {
		return 0
	}
	return userID.(int64)
}

// GetUserEmail extracts user email from context
func GetUserEmail(c *gin.Context) string {
	email, exists := c.Get("user_email")
	if !exists {
		return ""
	}
	return email.(string)
}

// GetUserFullName extracts user full name from context
func GetUserFullName(c *gin.Context) string {
	fullName, exists := c.Get("user_full_name")
	if !exists {
		return ""
	}
	return fullName.(string)
}

// GetClaims extracts JWT claims from context
func GetClaims(c *gin.Context) *utils.JWTClaims {
	claims, exists := c.Get("claims")
	if !exists {
		return nil
	}
	return claims.(*utils.JWTClaims)
}