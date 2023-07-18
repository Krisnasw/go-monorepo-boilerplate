package user

import "go-monorepo-boilerplate/services/auth-svc/entity"

type IUseCase interface {
	Login(string, string) (*entity.User, error)
	Register(string, string, string) error
	ForgotPassword(string) error
}
