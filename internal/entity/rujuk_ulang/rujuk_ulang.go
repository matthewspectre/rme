package rujukulang

import "time"

// Domain entity untuk rujuk ulang.
type RujukUlang struct {
	ID                 int
	IDPasien           int
	IDDokter           int
	PoliAsal           int
	PoliTujuan         int
	DiagnosisSementara string
	Catatan            string
	DateMake           time.Time
	DateUpdate         time.Time
	Visible            int
	NamaPasien         string
	NamaPoliAsal       string
	NamaPoliTujuan     string
	NamaDokter         string
}
