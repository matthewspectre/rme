package anamnesis

import "time"

// Domain entity definitions for anamnesis.

type Anamnesis struct {
	IDPasien                int
	IDDokter                int
	Text                    string
	DateMake                time.Time
	DateUpdate              time.Time
	IDDataKlinik            int
	RiwayatPengobatan       string
	RiwayatKeluarga         string
	RiwayatPekerjaan        string
	RiwayatAutoanamnesis    string
	RiwayatPenyakitDahulu   string
	RiwayatPenyakitSekarang string
	RiwayatPenyakitLain     string
	HubunganPasien          string
	RiwayatAnestesiBedah    string
	RiwayatKeluhanUtama     string
	StatusKehamilan         string
	KeluhanTambahan         string
	Catatan                 string
	Visible                 int
}
