package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/team-18-project/InnoSync/backend/config"
	_ "github.com/team-18-project/InnoSync/backend/docs"
	"github.com/team-18-project/InnoSync/backend/handlers"
	"github.com/team-18-project/InnoSync/backend/middleware"
	"github.com/team-18-project/InnoSync/backend/services"
	"github.com/team-18-project/InnoSync/backend/utils"
)

// @title InnoSync API
// @version 1.0
// @description RESTful API for InnoSync Mobile Application
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:3000
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Set gin mode
	gin.SetMode(cfg.Server.Mode)

	// Initialize database
	db, err := utils.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Run database migrations
	if err := db.Migrate(); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	// Initialize token manager
	tokenManager := utils.NewTokenManager(cfg)

	// Initialize services
	authService := services.NewAuthService(db.DB, tokenManager)
	profileService := services.NewProfileService(db.DB)
	projectService := services.NewProjectService(db.DB)
	applicationService := services.NewApplicationService(db.DB, projectService)
	invitationService := services.NewInvitationService(db.DB, projectService)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authService)
	profileHandler := handlers.NewProfileHandler(profileService, cfg)
	projectHandler := handlers.NewProjectHandler(projectService)
	applicationHandler := handlers.NewApplicationHandler(applicationService)
	invitationHandler := handlers.NewInvitationHandler(invitationService)

	// Initialize router
	router := gin.Default()

	// Add middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Serve static files
	router.Static("/uploads", cfg.Upload.UploadDir)

	// Swagger UI endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Root endpoint - redirect to Swagger UI
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message":     "Welcome to InnoSync API",
			"version":     "1.0",
			"swagger_ui":  "/swagger/index.html",
			"health":      "/health",
			"api_base":    "/api/v1",
			"timestamp":   time.Now().Format(time.RFC3339),
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "healthy",
			"timestamp": time.Now().Format(time.RFC3339),
		})
	})

	// API routes
	apiV1 := router.Group("/api/v1")
	{
		// Authentication routes
		authHandler.RegisterRoutes(apiV1)

		// Authentication middleware
		authMiddleware := middleware.AuthMiddleware(tokenManager)

		// Profile routes (with authentication middleware)
		profileHandler.RegisterRoutes(apiV1, authMiddleware)

		// Project routes (with authentication middleware)
		projectHandler.RegisterRoutes(apiV1, authMiddleware)

		// Application routes (with authentication middleware)
		applicationHandler.RegisterRoutes(apiV1, authMiddleware)

		// Invitation routes (with authentication middleware)
		invitationHandler.RegisterRoutes(apiV1, authMiddleware)

		// Protected route for getting current user info
		apiV1.GET("/auth/me", authMiddleware, authHandler.Me)
	}

	// Start server
	log.Printf("Server starting on %s", cfg.GetServerAddress())
	if err := router.Run(cfg.GetServerAddress()); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// corsMiddleware handles CORS
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
