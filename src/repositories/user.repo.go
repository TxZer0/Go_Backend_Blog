package repositories

import (
	"github.com/TxZer0/Go_Backend_Blog/src/database"
	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo() *UserRepo {
	return &UserRepo{
		db: database.DB,
	}
}

func (ur *UserRepo) Create(user *models.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) GetUserById(userId uint) (*models.User, error) {
	var user models.User
	if err := ur.db.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepo) UpdateVerifyAccount(updatedUser *models.User) error {
	err := ur.db.Model(&models.User{}).Where("id = ?", updatedUser.ID).Updates(map[string]interface{}{
		"IsVerify": true,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepo) UpdateUserPassword(updatedUser *models.User, hashedPassword string) error {
	err := ur.db.Model(&models.User{}).Where("id = ?", updatedUser.ID).Updates(map[string]interface{}{
		"password": hashedPassword,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
