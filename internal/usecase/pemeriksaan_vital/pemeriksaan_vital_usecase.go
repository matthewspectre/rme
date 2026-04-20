package pemeriksaanvital

import (
	entity "rme/internal/entity/pemeriksaan_vital"
	repo "rme/internal/repository/pemeriksaan_vital"
)

type Usecase interface {
	Create(data *entity.PemeriksaanVital) error
	GetByID(id int) (*entity.PemeriksaanVital, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error)
	Update(id int, updates map[string]interface{}) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.PemeriksaanVital) error {
	return u.repo.Create(data)
}

func (u *usecase) GetByID(id int) (*entity.PemeriksaanVital, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}
