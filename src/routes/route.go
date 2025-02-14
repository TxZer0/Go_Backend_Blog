package routes

import (
	"net/http"
	"os"

	"github.com/TxZer0/Go_Backend_Blog/src/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9999"
	}

	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Server running on port " + PORT,
		})
	})

	logger := middlewares.InitLogger()
	defer logger.Sync()
	r.Use(middlewares.RateLimitMiddleware(middlewares.InitRateLimit()))
	r.Use(middlewares.LoggerMiddleware(logger))
	r.Use(middlewares.ErrorHandleMiddleware())

	api := r.Group("/api")
	{
		InitUserRoutes(api)
		InitPostRoutes(api)
		InitCommentRoutes(api)
	}
	r.Run(":" + PORT)
	return r
}
