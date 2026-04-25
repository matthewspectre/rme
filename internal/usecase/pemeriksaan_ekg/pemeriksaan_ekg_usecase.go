package pemeriksaan_ekg

import (
	entity "rme/internal/entity/pemeriksaan_ekg"
	repo "rme/internal/repository/pemeriksaan_ekg"
)

type Usecase interface {
	Create(p *entity.PemeriksaanEkg) (int, error)
	GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error)
	GetByID(id int) (*repo.PemeriksaanEkgWithNames, error)
	Update(id int, updates map[string]interface{}) (*repo.PemeriksaanEkgWithNames, error)
	Delete(id int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(p *entity.PemeriksaanEkg) (int, error) {
	return u.repo.Create(p)
}

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) GetByID(id int) (*repo.PemeriksaanEkgWithNames, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Update(id int, updates map[string]interface{}) (*repo.PemeriksaanEkgWithNames, error) {
	return u.repo.Update(id, updates)
}

func (u *usecase) Delete(id int) error {
	return u.repo.Delete(id)
}
