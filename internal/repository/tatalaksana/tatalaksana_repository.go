package tatalaksana

import (
	entity "rme/internal/entity/tatalaksana"
)

type Repository interface {
	Create(data *entity.Tatalaksana) error
	GetByID(id int) (*entity.Tatalaksana, error)
	GetAll(idPasien *int, idDokter *int) ([]*entity.Tatalaksana, error)
	Update(id int, updates map[string]interface{}) error
}
