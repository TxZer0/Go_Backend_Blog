package controllers

import (
	"net/http"
	"strconv"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/request"
	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/TxZer0/Go_Backend_Blog/src/services"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService *services.CommentService
}

func NewCommentController() *CommentController {
	return &CommentController{
		commentService: services.NewCommentService(),
	}
}

func (cc *CommentController) Get(c *gin.Context) {
	commentId, err := strconv.ParseUint(c.Param("commentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := cc.commentService.Get(uint(commentId))
	c.JSON(code, obj)
}

func (cc *CommentController) Create(c *gin.Context) {
	var request request.CreateComment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := cc.commentService.Create(c.GetHeader("access_token"), &request)
	c.JSON(code, obj)
}

func (cc *CommentController) Update(c *gin.Context) {
	var request request.UpdateComment
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := cc.commentService.Update(c.GetHeader("access_token"), &request)
	c.JSON(code, obj)
}

func (cc *CommentController) Delete(c *gin.Context) {
	commentId, err := strconv.ParseUint(c.Param("commentId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := cc.commentService.Delete(c.GetHeader("access_token"), uint(commentId))
	c.JSON(code, obj)
}
