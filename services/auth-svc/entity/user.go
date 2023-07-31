package entity

import (
	"time"
)

type AuthRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=8,alphanumunicode"`
}

type User struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserAccess struct {
	Id          string    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Email       string    `json:"email" validate:"required,email"`
	Username    string    `json:"username"`
	PhoneNumber string    `json:"phoneNumber" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"accessToken"`
}
