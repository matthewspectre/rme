package antrian

import (
	"rme/internal/entity/antrian"
	repository "rme/internal/repository/antrian"
)

type Usecase interface {
	Create(a *antrian.Antrian) (int, error)
	GetAll(idDokter *int) ([]*repository.AntrianWithNamaPasien, error)
	Update(id int, updates map[string]interface{}) (*repository.AntrianWithNamaPasien, error)
	UpdateByPatient(idPasien int, updates map[string]interface{}) ([]*repository.AntrianWithNamaPasien, error)
	DeleteByPatient(idPasien int) error
}

type usecase struct {
	repo repository.Repository
}

func NewUsecase(repo repository.Repository) Usecase {
	return &usecase{repo: repo}
}

func (u *usecase) Create(a *antrian.Antrian) (int, error) {
	return u.repo.Create(a)
}

func (u *usecase) GetAll(idDokter *int) ([]*repository.AntrianWithNamaPasien, error) {
	return u.repo.GetAll(idDokter)
}

func (u *usecase) Update(id int, updates map[string]interface{}) (*repository.AntrianWithNamaPasien, error) {
	return u.repo.Update(id, updates)
}

func (u *usecase) UpdateByPatient(idPasien int, updates map[string]interface{}) ([]*repository.AntrianWithNamaPasien, error) {
	return u.repo.UpdateByPatient(idPasien, updates)
}

func (u *usecase) DeleteByPatient(idPasien int) error {
	return u.repo.DeleteByPatient(idPasien)
}
