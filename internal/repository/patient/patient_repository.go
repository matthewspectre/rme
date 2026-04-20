package patient

import (
	"rme/internal/entity/patient"
)

type Repository interface {
	Create(data *patient.Patient) error
	Update(data *patient.Patient) error
	GetByNIK(nik string) (*patient.Patient, error)
	GetAll() ([]*patient.Patient, error)
	GetAllByIDDataKlinik(idDataKlinik int) ([]*patient.Patient, error)
	// GetByID mengambil pasien berdasarkan kolom primary key `id`.
	GetByID(idPasien int) (*patient.Patient, error)
	// GetAllByID mengambil daftar pasien yang cocok dengan nilai `id` (biasanya 0/1 result).
	GetAllByID(idPasien int) ([]*patient.Patient, error)
	// CreateWithGeneratedRM inserts patient and generates `no_rekam_medis` as RM-<year>-<seq>
	// when `data.NoRekamMedis` is empty. Operation is transactional and safe for concurrency.
	CreateWithGeneratedRM(data *patient.Patient) error
}
