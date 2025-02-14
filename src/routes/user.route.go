package routes

import (
	"github.com/TxZer0/Go_Backend_Blog/src/controllers"
	"github.com/gin-gonic/gin"
)

func InitUserRoutes(api *gin.RouterGroup) {
	{
		api.GET("/verify/:token", controllers.NewUserController().VerifyEmail)
		api.POST("/forgot", controllers.NewUserController().ForgotPassword)
		api.POST("/register", controllers.NewUserController().Register)
		api.POST("/login", controllers.NewUserController().Login)
		api.POST("/change/:token", controllers.NewUserController().ChangePassword)
	}

}
