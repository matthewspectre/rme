package doctor

import (
	entity "rme/internal/entity/doctor"
	repo "rme/internal/repository/doctor"
)

type Usecase interface {
	Create(data *entity.Doctor) error
	GetAll() ([]*entity.Doctor, error)
	GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Doctor, error)
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.Doctor) error {
	return u.repo.Create(data)
}

func (u *usecase) GetAll() ([]*entity.Doctor, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Doctor, error) {
	return u.repo.GetAllByIDDataKlinik(idDataKlinik)
}
