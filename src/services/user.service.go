package services

import (
	"net/http"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/request"
	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"github.com/TxZer0/Go_Backend_Blog/src/repositories"
	"github.com/TxZer0/Go_Backend_Blog/src/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repositories.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepo(),
	}
}

func (us *UserService) Login(request *request.Login) (int, interface{}) {
	user, err := us.userRepo.GetUserByEmail(request.Email)
	if err != nil || !user.IsVerify {
		return http.StatusOK, response.NewWrongEmailOrPassword()
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return http.StatusOK, response.NewWrongEmailOrPassword()
	}

	accessToken, refreshToken, err := utils.GenerateTokenPair(user.ID)
	if err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}

	return http.StatusOK, response.NewSuccessResponse(map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func (us *UserService) Register(request *request.Register) (int, interface{}) {
	if request.Password != request.VerifyPassword {
		return http.StatusOK, response.NewPasswordDoNotMatch()
	}
	if user, err := us.userRepo.GetUserByEmail(request.Email); err == nil {
		if user.IsVerify {
			return http.StatusOK, response.NewEmailAlreadyExists()
		} else {
			if err := utils.SendEmail(request.Email, utils.VerifyAccount(request.Email)); err != nil {
				return http.StatusInternalServerError, response.NewInternalError()
			}
			return http.StatusOK, response.NewSuccessResponse(nil)
		}
	}
	if err := utils.SendEmail(request.Email, utils.VerifyAccount(request.Email)); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}

	user := models.User{
		Email:    request.Email,
		Password: string(hashedPassword),
	}

	if err := us.userRepo.Create(&user); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewSuccessResponse(nil)
}

func (us *UserService) VerifyEmail(token string) (int, interface{}) {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	email, ok := (*claims)["email"].(string)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	if user.IsVerify {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	if err := us.userRepo.UpdateVerifyAccount(user); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewVerifyEmailResponse()
}

func (us *UserService) ForgotPassword(email string) (int, interface{}) {
	if _, err := us.userRepo.GetUserByEmail(email); err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	if err := utils.SendEmail(email, utils.ChangePassword(email)); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewSuccessResponse(nil)
}

func (us *UserService) ChangePassword(token, newPassword string) (int, interface{}) {
	claims, err := utils.VerifyToken(token)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	email, ok := (*claims)["email"].(string)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	user, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	if !user.IsVerify {
		return http.StatusBadRequest, response.NewBadRequest()
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	if err := us.userRepo.UpdateUserPassword(user, string(hashedPassword)); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewSuccessResponse(nil)
}
