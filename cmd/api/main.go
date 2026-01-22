package main

import (
	"log"
	"strconv"

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
	port, _ := strconv.Atoi(cfg.Database.Port)
	dbConfig := &database.Config{
		Host:     cfg.Database.Host,
		Port:     port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.Database,
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

	// Initialize market data repository
	marketDataRepo := persistence.NewMarketDataRepository(db)

	// Initialize use cases
	authUseCase := usecases.NewAuthUseCase(userRepo, jwtManager)
	marketDataUseCase := usecases.NewMarketDataUseCase(marketDataRepo)

	// Initialize handlers
	authHandler := handlers.NewAuthHandler(authUseCase)
	marketDataHandler := handlers.NewMarketDataHandler(marketDataUseCase)

	// Setup Gin
	if cfg.Auth.JWTSecret == "your-secret-key-change-in-production" {
		log.Println("WARNING: Using default JWT secret. Please change in production!")
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, authHandler, marketDataHandler)

	// Start server
	log.Printf("Starting server on %s", cfg.Server.Port)
	if err := r.Run(cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
