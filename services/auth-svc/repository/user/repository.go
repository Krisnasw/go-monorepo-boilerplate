package user

import (
	"errors"

	"gorm.io/gorm"

	"go-monorepo-boilerplate/exceptions"
	"go-monorepo-boilerplate/services/auth-svc/app/repository"
	"go-monorepo-boilerplate/services/auth-svc/entity"
)

type Repository struct {
	db    *gorm.DB
	table string
}

func New(db *gorm.DB, table string) repository.IUserRepository {
	return &Repository{db, table}
}

// GetByEmail implements repository.IUserRepository.
func (r *Repository) GetByEmail(email string) (*entity.User, error) {
	userData := &User{}

	err := r.db.Raw(`SELECT * FROM drivers
		WHERE email = ?
	`, email).First(&userData).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, exceptions.DataNotFound
		}
		return nil, err
	}

	data := userData.ToEntity()
	return data, err
}

// ForgotPassword implements repository.IUserRepository.
func (*Repository) ForgotPassword(string) error {
	panic("unimplemented")
}

// Login implements repository.IUserRepository.
func (*Repository) Login(string, string) (*entity.User, error) {
	panic("unimplemented")
}

// Register implements repository.IUserRepository.
func (*Repository) Register(string, string, string) error {
	panic("unimplemented")
}
