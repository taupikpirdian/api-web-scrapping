package routes

import (
	"github.com/gin-gonic/gin"

	"api-web-scrapping/internal/presentation/handlers"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler, stockPriceSummaryHandler *handlers.StockPriceSummaryHandler) {
	api := r.Group("/api/v1")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}

		// Stock price summary routes
		stockPrices := api.Group("/stock-prices")
		{
			// Get all stock prices with pagination
			stockPrices.GET("", stockPriceSummaryHandler.GetAll)

			// Get stock price by ID
			stockPrices.GET("/:id", stockPriceSummaryHandler.GetByID)

			// Get stock price by symbol
			stockPrices.GET("/symbol/:symbol", stockPriceSummaryHandler.GetBySymbol)

			// Get latest stock price by symbol
			stockPrices.GET("/symbol/:symbol/latest", stockPriceSummaryHandler.GetLatestBySymbol)

			// Get stock price by symbol and date
			stockPrices.GET("/symbol/:symbol/date/:date", stockPriceSummaryHandler.GetBySymbolAndDate)

			// Get stock prices by date range
			stockPrices.GET("/range", stockPriceSummaryHandler.GetByDateRange)

			// Get stock prices by symbol and date range
			stockPrices.GET("/symbol/:symbol/range", stockPriceSummaryHandler.GetBySymbolAndDateRange)

			// Get top movers (gainers & losers) for a specific date
			stockPrices.GET("/movers/:date", stockPriceSummaryHandler.GetTopMovers)

			// Create stock price summary (admin only)
			stockPrices.POST("", stockPriceSummaryHandler.Create)

			// Update stock price summary (admin only)
			stockPrices.PUT("/:id", stockPriceSummaryHandler.Update)

			// Delete stock price summary (admin only)
			stockPrices.DELETE("/:id", stockPriceSummaryHandler.Delete)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
