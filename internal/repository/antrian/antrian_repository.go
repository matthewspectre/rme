package antrian

import "rme/internal/entity/antrian"

// AntrianWithNamaPasien adalah tipe hasil yang dikembalikan repository
// berisi nama pasien yang di-join dari tabel patients.
type AntrianWithNamaPasien struct {
	ID           int
	IDPasien     int
	NamaPasien   string
	IDDokter     int
	NamaDokter   string
	NomorAntrian int
	IDPoli       int
	Waktu        string
}

type Repository interface {
	// Create inserts antrian and returns the assigned nomorAntrian.
	Create(a *antrian.Antrian) (int, error)
	GetAll(idDokter *int) ([]*AntrianWithNamaPasien, error)
	// Update updates fields of antrian with given id and returns updated row (with joined names)
	Update(id int, updates map[string]interface{}) (*AntrianWithNamaPasien, error)
}
