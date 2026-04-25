package pemeriksaan_laboratorium

// Entity pemeriksaan laboratorium
type PemeriksaanLaboratorium struct {
	ID           int
	IDPasien     int
	IDDokter     int
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
	Visible      int
}
