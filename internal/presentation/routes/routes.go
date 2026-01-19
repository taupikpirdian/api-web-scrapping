package routes

import (
	"github.com/gin-gonic/gin"

	"api-web-scrapping/internal/presentation/handlers"
)

func SetupRoutes(r *gin.Engine, authHandler *handlers.AuthHandler) {
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
			auth.POST("/register", authHandler.Register)
		}
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
}
