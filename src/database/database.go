package database

import (
	"fmt"
	"log"
	"os"

	"github.com/TxZer0/Go_Backend_Blog/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	username := os.Getenv("DATABASE_USERNAME")
	if username == "" {
		username = "root"
	}
	password := os.Getenv("DATABASE_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/blog?parseTime=true", username, password)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("Cannot connect to MySQL: ", err)
	}

	fmt.Println("Connected to MySQL successfully!")

	err = db.AutoMigrate(&models.Post{}, &models.Comment{}, &models.User{})
	if err != nil {
		log.Fatalln("Failed to migrate tables!")
	}

	DB = db
	fmt.Println("Database migration successfully!")
}
