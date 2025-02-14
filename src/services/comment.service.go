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

type CommentService struct {
	commentRepo *repositories.CommentRepo
}

func NewCommentService() *CommentService {
	return &CommentService{
		commentRepo: repositories.NewCommentRepo(),
	}
}

func (cs *CommentService) Get(commentId uint) (int, interface{}) {
	comment, err := cs.commentRepo.GetCommentById(commentId)
	if err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}
	return http.StatusOK, comment
}

func (cs *CommentService) Create(accessToken string, request *request.CreateComment) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return 200, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return 500, response.NewBadRequest()
	}

	if _, err := cs.commentRepo.GetPostById(request.PostID); err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}
	comment := models.Comment{
		Content: request.Content,
		PostID:  request.PostID,
		UserID:  uint(idFloat),
	}

	if err := cs.commentRepo.Create(&comment); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, comment
}

func (cs *CommentService) Update(accessToken string, request *request.UpdateComment) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	comment, err := cs.commentRepo.GetCommentById(request.ID)
	if err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}

	if uint(idFloat) != comment.UserID {
		return http.StatusForbidden, response.NewForbidden()
	}

	if comment.PostID != request.PostID {
		return http.StatusNotFound, response.NewNotFound()
	}

	comment.Content = request.Content
	comment.UpdatedAt = time.Now().UTC()

	if err := cs.commentRepo.Update(comment); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, comment
}

func (cs *CommentService) Delete(accessToken string, commentId uint) (int, interface{}) {
	claims, err := utils.VerifyToken(accessToken)
	if err != nil {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	idFloat, ok := (*claims)["user_id"].(float64)
	if !ok {
		return http.StatusBadRequest, response.NewBadRequest()
	}

	comment, err := cs.commentRepo.GetCommentById(commentId)
	if err != nil {
		return http.StatusNotFound, response.NewNotFound()
	}

	if uint(idFloat) != comment.UserID {
		return http.StatusForbidden, response.NewForbidden()
	}

	if err := cs.commentRepo.DeleteById(commentId); err != nil {
		return http.StatusInternalServerError, response.NewInternalError()
	}
	return http.StatusOK, response.NewSuccessResponse(nil)
}
