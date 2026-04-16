package doctor

import "rme/internal/entity/doctor"

type Repository interface {
	Create(data *doctor.Doctor) error
	GetAll() ([]*doctor.Doctor, error)
	GetAllByIDDataKlinik(idDataKlinik int) ([]*doctor.Doctor, error)
}
