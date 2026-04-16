package poli

import "rme/internal/entity/poli"

type Repository interface {
	GetAll(idDataKlinik *int) ([]*poli.Poli, error)
}
