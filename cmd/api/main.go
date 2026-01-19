package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"api-web-scrapping/internal/application/usecases"
	"api-web-scrapping/internal/infrastructure/config"
	"api-web-scrapping/internal/infrastructure/database"
	"api-web-scrapping/internal/infrastructure/persistence"
	"api-web-scrapping/internal/presentation/handlers"
	"api-web-scrapping/internal/presentation/routes"
	"api-web-scrapping/pkg/auth"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database connection
	dbConfig := &database.Config{
		Host:     cfg.Database.Host,
		Port:     5432, // Default PostgreSQL port
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.Database,
		SSLMode:  cfg.Database.SSLMode,
	}

	db, err := database.NewConnection(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	log.Println("Database connected successfully")

	// Initialize dependencies
	jwtManager := auth.NewJWTManager(cfg.Auth.JWTSecret, cfg.Auth.TokenDuration)
	userRepo := persistence.NewInMemoryUserRepository()

	// Initialize stock price summary repository
	stockPriceSummaryRepo := persistence.NewStockPriceSummaryRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, jwtManager)
	stockPriceSummaryUseCase := usecases.NewStockPriceSummaryUseCase(stockPriceSummaryRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	stockPriceSummaryHandler := handlers.NewStockPriceSummaryHandler(stockPriceSummaryUseCase)

	// Setup Gin
	if cfg.Auth.JWTSecret == "your-secret-key-change-in-production" {
		log.Println("WARNING: Using default JWT secret. Please change in production!")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, authHandler, stockPriceSummaryHandler)

	// Start server
	log.Printf("Starting server on %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
