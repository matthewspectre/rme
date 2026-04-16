package patient

// Use case (business logic) for patients.

import (
	"fmt"

	entity "rme/internal/entity/patient"
	repo "rme/internal/repository/patient"
)

type Usecase interface {
	Create(data *entity.Patient) error
	CreateOrUpdateByNIK(data *entity.Patient) (bool, error)
	GetAll() ([]*entity.Patient, error)
	GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Patient, error)
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.Patient) error {
	if data == nil {
		return fmt.Errorf("data is nil")
	}
	// If no rekam medis provided, use repository transactional generator
	if data.NoRekamMedis == "" {
		return u.repo.CreateWithGeneratedRM(data)
	}
	return u.repo.Create(data)
}

// CreateOrUpdateByNIK akan membuat pasien baru jika NIK belum ada,
// atau memperbarui data pasien jika NIK sudah terdaftar.
// Mengembalikan updated=true jika data dilakukan update.
func (u *usecase) CreateOrUpdateByNIK(data *entity.Patient) (bool, error) {
	existing, err := u.repo.GetByNIK(data.NIK)
	if err != nil {
		return false, err
	}
	if existing != nil {
		// Pakai ID lama supaya dilakukan update, bukan insert baru.
		data.ID = existing.ID
		if err := u.repo.Update(data); err != nil {
			return false, err
		}
		return true, nil
	}
	// when creating new patient, use generator to ensure no_rekam_medis is set
	if err := u.repo.CreateWithGeneratedRM(data); err != nil {
		return false, err
	}
	return false, nil
}

func (u *usecase) GetAll() ([]*entity.Patient, error) {
	return u.repo.GetAll()
}

func (u *usecase) GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Patient, error) {
	return u.repo.GetAllByIDDataKlinik(idDataKlinik)
}
