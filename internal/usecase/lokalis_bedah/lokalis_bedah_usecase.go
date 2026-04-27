package lokalis_bedah

import (
	"time"

	entity "rme/internal/entity/lokalis_bedah"
	repo "rme/internal/repository/lokalis_bedah"
)

type Usecase interface {
	Create(d *entity.LokalisBedah) error
	GetByID(id int) (*entity.LokalisBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(d *entity.LokalisBedah) error {
	if d.Tanggal.IsZero() {
		d.Tanggal = time.Now()
	}
	return u.repo.Create(d)
}

func (u *usecase) GetByID(id int) (*entity.LokalisBedah, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	return u.repo.GetAll(idDokter, idPasien)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}

func (u *usecase) Hide(id int) error {
	return u.repo.Hide(id)
}
