package tatalaksana

import (
	entity "rme/internal/entity/tatalaksana"
	repo "rme/internal/repository/tatalaksana"
)

type Usecase interface {
	Create(data *entity.Tatalaksana) error
	GetByID(id int) (*entity.Tatalaksana, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.Tatalaksana, error)
	Update(id int, updates map[string]interface{}) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.Tatalaksana) error {
	return u.repo.Create(data)
}

func (u *usecase) GetByID(id int) (*entity.Tatalaksana, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*entity.Tatalaksana, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}
