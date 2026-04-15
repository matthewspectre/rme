package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	Create(data *anamnesis.Anamnesis) error
	GetByID(idPasien int) (*anamnesis.Anamnesis, error)
	GetAll() ([]*anamnesis.Anamnesis, error)
}
