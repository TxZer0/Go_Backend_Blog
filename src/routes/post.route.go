package routes

import (
	"github.com/TxZer0/Go_Backend_Blog/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitPostRoutes(api *gin.RouterGroup) {
	post := api.Group("/post")
	{
		post.GET("/", controllers.NewPostController().GetAllPosts)
		post.GET("/:postId", controllers.NewPostController().GetPostById)
		post.POST("/", controllers.NewPostController().Create)
		post.PUT("/", controllers.NewPostController().Update)
		post.DELETE("/:postId", controllers.NewPostController().Delete)
	}

}
