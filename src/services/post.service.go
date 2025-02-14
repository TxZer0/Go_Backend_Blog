package services

import (
	"net/http"
	"time"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/request"
	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"github.com/TxZer0/Go_Backend_Blog/src/repositories"
	"github.com/TxZer0/Go_Backend_Blog/src/utils"
)

type PostService struct {
	postRepo *repositories.PostRepo
}

func NewPostService() *PostService {
	return &PostService{
		postRepo: repositories.NewPostRepo(),
	}
}

func (ps *PostService) Create(accessToken string, request *request.CreatePost) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	if uint(idFloat) != 1 {
		return http.StatusForbidden, response.NewForbidden()
	}

	post := models.Post{
		Title:   request.Title,
		Content: request.Content,
	}

	if err := ps.postRepo.Create(&post); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, post
}
func (ps *PostService) Get() (int, interface{}) {
	posts, err := ps.postRepo.GetAllPosts()
	if err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, posts
}

func (ps *PostService) GetPostById(postId uint) (int, interface{}) {
	post, err := ps.postRepo.GetPostById(postId)
	if err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}
	return http.StatusOK, post
}
func (ps *PostService) Update(accessToken string, request *request.UpdatePost) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	if uint(idFloat) != 1 {
		return http.StatusForbidden, response.NewForbidden()
	}

	if _, err := ps.postRepo.GetPostById(request.ID); err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}

	post := models.Post{
		ID:        request.ID,
		Title:     request.Title,
		Content:   request.Content,
		UpdatedAt: time.Now(),
	}

	if err := ps.postRepo.Update(&post); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, post
}

func (ps *PostService) Delete(accessToken string, postId uint) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	if uint(idFloat) != 1 {
		return http.StatusForbidden, response.NewForbidden()
	}

	if _, err := ps.postRepo.GetPostById(postId); err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}

	if err := ps.postRepo.DeleteById(postId); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewSuccessResponse(nil)
}
