package anamnesis

import (
	"rme/internal/entity/anamnesis"
)

type Repository interface {
	Create(data *anamnesis.Anamnesis) error
	// GetByID mengambil anamnesis berdasarkan id_anamnesis, bukan id pasien
	GetByID(idAnamnesis int) (*anamnesis.Anamnesis, error)
	GetAll(idDokter *int, idPasien *int) ([]*anamnesis.Anamnesis, error)
	Update(idAnamnesis int, updates map[string]interface{}) error
}
