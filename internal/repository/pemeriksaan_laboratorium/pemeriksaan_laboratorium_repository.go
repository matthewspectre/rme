package pemeriksaan_laboratorium

import "rme/internal/entity/pemeriksaan_laboratorium"

type Repository interface {
	Create(p *pemeriksaan_laboratorium.PemeriksaanLaboratorium) (int, error)
	GetAll(idPasien *int, idDokter *int) ([]*PemeriksaanLaboratoriumWithNames, error)
	GetByID(id int) (*PemeriksaanLaboratoriumWithNames, error)
	Update(id int, updates map[string]interface{}) (*PemeriksaanLaboratoriumWithNames, error)
	Delete(id int) error
}

type PemeriksaanLaboratoriumWithNames struct {
	ID           int
	IDPasien     int
	NamaPasien   string
	IDDokter     int
	NamaDokter   string
	Visible      int
	Hb           *float64
	Ht           *float64
	Leukosit     *int
	Trombosit    *int
	GulaPuasa    *float64
	GulaSewaktu  *float64
	HbA1c        *float64
	Kolesterol   *float64
	HDL          *float64
	LDL          *float64
	Trigliserida *float64
	SGOT         *float64
	SGPT         *float64
	Ureum        *float64
	Kreatinin    *float64
	AsamUrat     *float64
	Natrium      *float64
	Kalium       *float64
	Klorida      *float64
	Waktu        string
}
