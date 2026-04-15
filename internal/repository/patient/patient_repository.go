package patient

import (
	"rme/internal/entity/patient"
)

type Repository interface {
	Create(data *patient.Patient) error
	Update(data *patient.Patient) error
	GetByNIK(nik string) (*patient.Patient, error)
	GetAll() ([]*patient.Patient, error)
}
