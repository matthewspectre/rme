package icd

import repo "rme/internal/repository/icd"

type Usecase interface {
	GetAll() ([]*repo.IcdRow, error)
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) GetAll() ([]*repo.IcdRow, error) {
	return u.repo.GetAll()
}
