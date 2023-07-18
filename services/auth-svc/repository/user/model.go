package user

import (
	"time"

	"golang.org/x/crypto/bcrypt"

	"go-monorepo-boilerplate/services/auth-svc/entity"
)

type User struct {
	Id          string `gorm:"primaryKey"`
	Name        string
	Email       string `gorm:"unique"`
	Username    string `gorm:"unique, omitempty"`
	Password    string
	PhoneNumber string    `gorm:"column:phone"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

func (user *User) ToEntityAuth() *entity.User {
	return &entity.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Local(),
		UpdatedAt:   user.UpdatedAt.Local(),
	}
}

func (user *User) ToEntity() *entity.User {
	return &entity.User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Local(),
		UpdatedAt:   user.UpdatedAt.Local(),
	}
}

func (User) FromEntity(user *entity.User) *User {
	return &User{
		Id:          user.Id,
		Name:        user.Name,
		Email:       user.Email,
		Username:    user.Username,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Local(),
		UpdatedAt:   user.UpdatedAt.Local(),
	}
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
