package pemeriksaan_laboratorium

import (
	entity "rme/internal/entity/pemeriksaan_laboratorium"
	repo "rme/internal/repository/pemeriksaan_laboratorium"
)

type Usecase interface {
	Create(p *entity.PemeriksaanLaboratorium) (int, error)
	GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanLaboratoriumWithNames, error)
	GetByID(id int) (*repo.PemeriksaanLaboratoriumWithNames, error)
	Update(id int, updates map[string]interface{}) (*repo.PemeriksaanLaboratoriumWithNames, error)
	Delete(id int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(p *entity.PemeriksaanLaboratorium) (int, error) {
	return u.repo.Create(p)
}

func (u *usecase) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanLaboratoriumWithNames, error) {
	return u.repo.GetAll(idPasien, idDokter)
}

func (u *usecase) GetByID(id int) (*repo.PemeriksaanLaboratoriumWithNames, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) Update(id int, updates map[string]interface{}) (*repo.PemeriksaanLaboratoriumWithNames, error) {
	return u.repo.Update(id, updates)
}

func (u *usecase) Delete(id int) error {
	return u.repo.Delete(id)
}
