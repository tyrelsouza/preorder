package models

import (
	"time"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateUserInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Username string `json:"username"`
}

func NewUser(id uint, username, password string, createdAt, updatedAt time.Time) User {
	user := User{
		ID:       id,
		Username: username,
		Password: password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
	return user
}
