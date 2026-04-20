package anamnesis

// Use case (business logic) for anamnesis.

import (
	entity "rme/internal/entity/anamnesis"
	repo "rme/internal/repository/anamnesis"
)

// Usecase mendefinisikan operasi bisnis untuk Anamnesis.
// Di atas repository (DB), di bawah handler (HTTP/API).
type Usecase interface {
	Create(data *entity.Anamnesis) error
	GetByID(idPasien int, idPoli int) (*entity.Anamnesis, error)
	GetAll(idPoli int) ([]*entity.Anamnesis, error)
}

// usecase adalah implementasi konkret dari Usecase.
type usecase struct {
	repo repo.Repository
}

// NewUsecase membuat instance baru usecase Anamnesis.
func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

// Create menjalankan proses bisnis saat membuat data Anamnesis baru.
// Saat ini hanya meneruskan ke repository (bisa ditambah validasi/logika lain nanti).
func (u *usecase) Create(data *entity.Anamnesis) error {
	return u.repo.Create(data)
}

// GetByID mengambil satu data Anamnesis berdasarkan ID pasien.
func (u *usecase) GetByID(idPasien int, idPoli int) (*entity.Anamnesis, error) {
	return u.repo.GetByID(idPasien, idPoli)
}

// GetAll mengambil semua data Anamnesis yang masih visible dan sesuai poli.
func (u *usecase) GetAll(idPoli int) ([]*entity.Anamnesis, error) {
	return u.repo.GetAll(idPoli)
}
