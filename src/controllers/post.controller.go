package controllers

import (
	"net/http"
	"strconv"

	"github.com/TxZer0/Go_Backend_Blog/src/dto/request"
	"github.com/TxZer0/Go_Backend_Blog/src/dto/response"
	"github.com/TxZer0/Go_Backend_Blog/src/services"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *services.PostService
}

func NewPostController() *PostController {
	return &PostController{
		postService: services.NewPostService(),
	}
}

func (pc *PostController) Create(c *gin.Context) {
	var request request.CreatePost
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := pc.postService.Create(c.GetHeader("access_token"), &request)
	c.JSON(code, obj)
}

func (pc *PostController) GetAllPosts(c *gin.Context) {
	code, obj := pc.postService.Get()
	c.JSON(code, obj)
}

func (pc *PostController) GetPostById(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewNotFound())
		return
	}
	code, obj := pc.postService.GetPostById(uint(postId))
	c.JSON(code, obj)
}

func (pc *PostController) Update(c *gin.Context) {
	var request request.UpdatePost
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, response.NewBadRequest())
		return
	}
	code, obj := pc.postService.Update(c.GetHeader("access_token"), &request)
	c.JSON(code, obj)
}

func (pc *PostController) Delete(c *gin.Context) {
	postId, err := strconv.ParseUint(c.Param("postId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, response.NewNotFound())
		return
	}
	code, obj := pc.postService.Delete(c.GetHeader("access_token"), uint(postId))
	c.JSON(code, obj)
}
