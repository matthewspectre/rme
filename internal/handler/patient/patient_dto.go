package patient

import "time"

// Request/response DTOs for patient handlers.

type PatientCreateRequest struct {
	Name           string `json:"namaPasien"`
	NoRekamMedis   string `json:"noRekamMedis"`
	AdmissionDate  string `json:"tanggalMasuk"` // format: "2006-01-02 15:04:05"
	NIK            string `json:"nik"`
	Gender         string `json:"jenisKelamin"`
	BloodType      string `json:"golonganDarah"`
	BirthPlaceDate string `json:"tempatTanggalLahir"`
	Phone          string `json:"nomorTelepon"`
	Address        string `json:"alamat"`
	Category       string `json:"kategori"`
	Job            string `json:"pekerjaan"`
	IDDataKlinik   int    `json:"idDataKlinik"`
}

type PatientResponse struct {
	ID             int       `json:"id"`
	Name           string    `json:"namaPasien"`
	NoRekamMedis   string    `json:"noRekamMedis"`
	AdmissionDate  time.Time `json:"tanggalMasuk"`
	NIK            string    `json:"nik"`
	Gender         string    `json:"jenisKelamin"`
	BloodType      string    `json:"golonganDarah"`
	BirthPlaceDate string    `json:"tempatTanggalLahir"`
	Phone          string    `json:"nomorTelepon"`
	Address        string    `json:"alamat"`
	Category       string    `json:"kategori"`
	Job            string    `json:"pekerjaan"`
	IDDataKlinik   int       `json:"idDataKlinik"`
}
