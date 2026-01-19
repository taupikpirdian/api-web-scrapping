package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"api-web-scrapping/internal/application/usecases"
	"api-web-scrapping/internal/infrastructure/config"
	"api-web-scrapping/internal/infrastructure/persistence"
	"api-web-scrapping/internal/presentation/handlers"
	"api-web-scrapping/internal/presentation/routes"
	"api-web-scrapping/pkg/auth"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize dependencies
	jwtManager := auth.NewJWTManager(cfg.Auth.JWTSecret, cfg.Auth.TokenDuration)
	userRepo := persistence.NewInMemoryUserRepository()
	authUseCase := usecases.NewAuthUseCase(userRepo, jwtManager)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)

	// Setup Gin
	if cfg.Auth.JWTSecret == "your-secret-key-change-in-production" {
		log.Println("WARNING: Using default JWT secret. Please change in production!")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, authHandler)

	// Start server
	log.Printf("Starting server on %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
