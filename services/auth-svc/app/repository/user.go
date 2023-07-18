package repository

import "go-monorepo-boilerplate/services/auth-svc/entity"

type IUserRepository interface {
	Login(string, string) (*entity.User, error)
	GetByEmail(string) (*entity.User, error)
	Register(string, string, string) error
	ForgotPassword(string) error
}
