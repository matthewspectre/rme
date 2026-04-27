package pemeriksaan_fungsi_organ

import (
	"time"

	entity "rme/internal/entity/pemeriksaan_fungsi_organ"
	repo "rme/internal/repository/pemeriksaan_fungsi_organ"
)

type Usecase interface {
	Create(d *entity.PemeriksaanFungsiOrgan) error
	GetByID(id int) (*entity.PemeriksaanFungsiOrgan, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(d *entity.PemeriksaanFungsiOrgan) error {
	if d.Tanggal.IsZero() {
		d.Tanggal = time.Now()
	}
	return u.repo.Create(d)
}

func (u *usecase) GetByID(id int) (*entity.PemeriksaanFungsiOrgan, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error) {
	return u.repo.GetAll(idDokter, idPasien)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}

func (u *usecase) Hide(id int) error {
	return u.repo.Hide(id)
}
