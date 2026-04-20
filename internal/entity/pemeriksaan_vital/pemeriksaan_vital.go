package pemeriksaanvital

import "time"

// Domain entity untuk pemeriksaan vital.
type PemeriksaanVital struct {
	ID             int
	IDPasien       int
	IDDokter       int
	DateMake       time.Time
	DateUpdate     time.Time
	TekananDarah   string
	DenyutNadi     int
	SuhuTubuh      float64
	FrekuensiNapas int
	BeratBadan     float64
	TinggiBadan    float64
	Visible        int
	NamaPasien     string
	NamaDokter     string
}
