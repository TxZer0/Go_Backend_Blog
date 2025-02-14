package models

import "time"

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique; not null"`
	Password  string    `json:"-" gorm:"not null"`
	IsVerify  bool      `gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
