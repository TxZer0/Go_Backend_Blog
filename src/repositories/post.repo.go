package repositories

import (
	"time"

	"github.com/TxZer0/Go_Backend_Blog/src/database"
	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	db *gorm.DB
}

func NewPostRepo() *PostRepo {
	return &PostRepo{
		db: database.DB,
	}
}

func (pr *PostRepo) Create(post *models.Post) error {
	if err := pr.db.Create(post).Error; err != nil {
		return err
	}
	return nil
}

func (pr *PostRepo) GetAllPosts() ([]models.Post, error) {
	var posts []models.Post
	if err := pr.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
func (pr *PostRepo) GetPostById(postId uint) (*models.Post, error) {
	var post models.Post
	if err := pr.db.Where("id = ?", postId).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (pr *PostRepo) Update(post *models.Post) error {
	err := pr.db.Model(&models.Post{}).Where("id = ?", post.ID).Updates(map[string]interface{}{
		"title":      post.Title,
		"content":    post.Content,
		"updated_at": time.Now(),
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *PostRepo) DeleteById(postId uint) error {
	if err := pr.db.Delete(&models.Post{}, postId).Error; err != nil {
		return err
	}
	return nil
}
