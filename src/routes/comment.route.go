package routes

import (
	"github.com/TxZer0/Go_Backend_Blog/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitCommentRoutes(api *gin.RouterGroup) {
	comment := api.Group("/comment")
	{
		comment.GET("/:commentId", controllers.NewCommentController().Get)
		comment.POST("/", controllers.NewCommentController().Create)
		comment.PUT("/", controllers.NewCommentController().Update)
		comment.DELETE("/:commentId", controllers.NewCommentController().Delete)
	}

}
