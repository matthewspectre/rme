package pemeriksaanvital

import (
	entity "rme/internal/entity/pemeriksaan_vital"
)

type Repository interface {
	Create(data *entity.PemeriksaanVital) error
	GetByID(id int) (*entity.PemeriksaanVital, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error)
	Update(id int, updates map[string]interface{}) error
}
