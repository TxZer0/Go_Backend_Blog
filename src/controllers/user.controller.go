package controllers

import (
	"net/http"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/request"
	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/TxZer0/Go_Backend_Blog/src/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.NewUserService(),
	}
}

func (uc *UserController) Login(c *gin.Context) {
	var request request.Login
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := uc.userService.Login(&request)
	c.JSON(code, obj)
}

func (uc *UserController) Register(c *gin.Context) {
	var request request.Register
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := uc.userService.Register(&request)
	c.JSON(code, obj)
}

func (uc *UserController) VerifyEmail(c *gin.Context) {
	code, obj := uc.userService.VerifyEmail(c.Param("token"))
	c.JSON(code, obj)
}

func (uc *UserController) ForgotPassword(c *gin.Context) {
	var email struct {
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&email); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := uc.userService.ForgotPassword(email.Email)
	c.JSON(code, obj)
}

func (uc *UserController) ChangePassword(c *gin.Context) {
	var newPassword struct {
		NewPassword string `json:"new_password"`
	}
	if err := c.ShouldBindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := uc.userService.ChangePassword(c.Param("token"), newPassword.NewPassword)
	c.JSON(code, obj)
}
