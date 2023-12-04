package models

import "time"

type User struct {
	Id        uint      `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username" gorm:"uniqueKey"`
	Email     string    `json:"email" gorm:"uniqueKey"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}

type UserResponse struct {
	Id        uint      `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserCreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserUpdateRequest struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserDeleteRequest struct {
	Id uint `json:"id"`
}
