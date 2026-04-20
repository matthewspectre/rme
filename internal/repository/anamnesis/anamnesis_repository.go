package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	Create(data *anamnesis.Anamnesis) error
	GetByID(idPasien int, idPoli int) (*anamnesis.Anamnesis, error)
	GetAll(idPoli int) ([]*anamnesis.Anamnesis, error)
}
