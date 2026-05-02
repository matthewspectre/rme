package pemeriksaan_penunjang_bedah

import (
	"time"

	entity "rme/internal/entity/pemeriksaan_penunjang_bedah"
	repo "rme/internal/repository/pemeriksaan_penunjang_bedah"
)

type Usecase interface {
	Create(d *entity.PemeriksaanPenunjangBedah) error
	GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
	HideByPatient(idPasien int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(d *entity.PemeriksaanPenunjangBedah) error {
	if d.DateMake.IsZero() {
		d.DateMake = time.Now()
	}
	if d.DateUpdate.IsZero() {
		d.DateUpdate = d.DateMake
	}
	return u.repo.Create(d)
}

func (u *usecase) GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	return u.repo.GetAll(idDokter, idPasien)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}

func (u *usecase) Hide(id int) error {
	return u.repo.Hide(id)
}

func (u *usecase) HideByPatient(idPasien int) error {
	return u.repo.HideByPatient(idPasien)
}
