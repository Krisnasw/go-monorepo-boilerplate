package user

import (
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	"go-monorepo-boilerplate/exceptions"
	"go-monorepo-boilerplate/services/auth-svc/app/repository"
	"go-monorepo-boilerplate/services/auth-svc/entity"
)

type IUseCaseImplementation struct {
	userRepository repository.IUserRepository
}

func NewUseCase(userRepository repository.IUserRepository) IUseCase {
	return &IUseCaseImplementation{userRepository: userRepository}
}

// ForgotPassword implements IUseCase.
func (uc *IUseCaseImplementation) ForgotPassword(string) error {
	panic("unimplemented")
}

// Login implements IUseCase.
func (uc *IUseCaseImplementation) Login(email, password string) (*entity.User, error) {
	user, err := uc.userRepository.GetByEmail(email)
	if err != nil {
		if errors.Is(err, exceptions.DataNotFound) {
			logrus.Errorf("%w: %v", exceptions.InternalServerError, err)
			return nil, exceptions.DataNotFound
		}
		logrus.Errorf("%w: %v", exceptions.InternalServerError, err)
		return nil, exceptions.InternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		logrus.Errorf("%w: %v", exceptions.InternalServerError, err)
		return nil, exceptions.PasswordNotMatch
	}

	return user, nil
}

// Register implements IUseCase.
func (uc *IUseCaseImplementation) Register(string, string, string) error {
	panic("unimplemented")
}
