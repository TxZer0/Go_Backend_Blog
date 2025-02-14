package repositories

import (
	"time"

	"github.com/TxZer0/Go_Backend_Blog/src/database"
	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"gorm.io/gorm"
)

type CommentRepo struct {
	db *gorm.DB
}

func NewCommentRepo() *CommentRepo {
	return &CommentRepo{
		db: database.DB,
	}
}

func (cr *CommentRepo) GetCommentByPostId(postId uint) ([]models.Comment, error) {
	var comments []models.Comment
	if err := cr.db.First(&comments, postId).Error; err != nil {
		return nil, err
	}
	return comments, nil
}

func (cr *CommentRepo) GetPostById(postId uint) (*models.Post, error) {
	var post models.Post
	err := cr.db.Where("id = ?", postId).First(&post).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (cr *CommentRepo) GetCommentById(commentId uint) (*models.Comment, error) {
	var comment *models.Comment
	if err := cr.db.First(&comment, commentId).Error; err != nil {
		return nil, err
	}
	return comment, nil
}
func (cr *CommentRepo) Create(comment *models.Comment) error {
	if err := cr.db.Create(comment).Error; err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepo) Update(comment *models.Comment) error {
	err := cr.db.Model(comment).Updates(map[string]interface{}{
		"content":    comment.Content,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (cr *CommentRepo) DeleteById(commentId uint) error {
	if err := cr.db.Delete(&models.Comment{}, commentId).Error; err != nil {
		return err
	}
	return nil
}
