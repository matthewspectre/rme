package auth

// Use case (business logic) for authentication.

import (
	"errors"

	entity "rme/internal/entity/user"
	repo "rme/internal/repository/user"
)

// Usecase mendefinisikan operasi bisnis untuk autentikasi.
type Usecase interface {
	Login(username, password string) (*entity.User, error)
}

type usecase struct {
	repo repo.Repository
}

// NewUsecase membuat instance baru usecase Auth.
func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

// Login memeriksa kredensial username dan password.
func (u *usecase) Login(username, password string) (*entity.User, error) {
	user, err := u.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("invalid credentials")
	}
	if user.Password != password {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
