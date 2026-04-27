package lokalis_bedah

import "time"

type LokalisBedah struct {
	ID             int
	IDPasien       int
	IDDokter       int
	Tanggal        time.Time
	LokasiKelainan string
	JenisKelainan  string
	Ukuran         string
	Warna          string
	NyeriTekan     bool
	Konsistensi    string
	Mobilitas      string
	TandaRadang    bool
	Fluktuasi      bool
	Catatan        string
	Visible        int
	NamaPasien     string
}
