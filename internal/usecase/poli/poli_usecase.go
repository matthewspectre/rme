package poli

import (
	"rme/internal/entity/poli"
	repository "rme/internal/repository/poli"
)

type Usecase interface {
	GetAll(idDataKlinik *int) ([]*poli.Poli, error)
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) GetAll(idDataKlinik *int) ([]*poli.Poli, error) {
	return u.repo.GetAll(idDataKlinik)
}
