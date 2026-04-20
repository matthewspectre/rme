package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	Create(data *anamnesis.Anamnesis) error
	GetByID(idPasien int) (*anamnesis.Anamnesis, error)
	GetAll(idDokter *int, idPasien *int) ([]*anamnesis.Anamnesis, error)
	Update(idAnamnesis int, updates map[string]interface{}) error
}
