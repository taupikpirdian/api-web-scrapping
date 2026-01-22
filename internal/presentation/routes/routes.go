package routes

import (
	"github.com/gin-gonic/gin"

	"api-web-scrapping/internal/presentation/handlers"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, marketDataHandler *handlers.MarketDataHandler) {
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		// Market data routes (from v_latest_market_data view)
		marketData := api.Group("/market-data")
		{
			// Get all market data from view
			marketData.GET("", marketDataHandler.GetAll)

			// Get latest market data for all emitens
			marketData.GET("/latest", marketDataHandler.GetLatestByAllEmiten)

			// Get market data by emiten
			marketData.GET("/emiten/:emiten", marketDataHandler.GetByEmiten)

			// Get latest market data by emiten
			marketData.GET("/emiten/:emiten/latest", marketDataHandler.GetLatestByEmiten)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
