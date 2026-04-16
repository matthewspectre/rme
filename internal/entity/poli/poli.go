package poli

// Entity untuk poli (daftar_poli)
type Poli struct {
	ID             int
	NamaPoli       string
	Deskripsi      string
	IDDataKlinik   int
	Aktif          bool
	DibuatPada     string // atau time.Time jika ingin parsing otomatis
	DiperbaruiPada string // atau time.Time
}
