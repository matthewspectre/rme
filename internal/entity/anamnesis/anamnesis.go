package anamnesis

import "time"

// Domain entity definitions for anamnesis.

type Anamnesis struct {
	ID                    int
	IDPasien              int
	IDDokter              int
	Text                  string
	DateMake              time.Time
	DateUpdate            time.Time
	IDDataKlinik          int
	RiwayatPengobatan     string
	RiwayatKeluarga       string
	RiwayatPenyakitDahulu string
	RiwayatPenyakitLain   string
	StatusKehamilan       string
	KeluhanTambahan       string
	Visible               int
	NamaPasien            string
}
