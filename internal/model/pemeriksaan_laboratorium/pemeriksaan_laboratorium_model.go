package pemeriksaan_laboratorium

// GORM model for pemeriksaan_laboratorium table
type PemeriksaanLaboratoriumModel struct {
	ID           int      `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien     int      `gorm:"column:id_pasien"`
	IDDokter     int      `gorm:"column:id_dokter"`
	Hb           *float64 `gorm:"column:hb"`
	Ht           *float64 `gorm:"column:ht"`
	Leukosit     *int     `gorm:"column:leukosit"`
	Trombosit    *int     `gorm:"column:trombosit"`
	GulaPuasa    *float64 `gorm:"column:gula_puasa"`
	GulaSewaktu  *float64 `gorm:"column:gula_sewaktu"`
	HbA1c        *float64 `gorm:"column:hba1c"`
	Kolesterol   *float64 `gorm:"column:kolesterol_total"`
	HDL          *float64 `gorm:"column:hdl"`
	LDL          *float64 `gorm:"column:ldl"`
	Trigliserida *float64 `gorm:"column:trigliserida"`
	SGOT         *float64 `gorm:"column:sgot"`
	SGPT         *float64 `gorm:"column:sgpt"`
	Ureum        *float64 `gorm:"column:ureum"`
	Kreatinin    *float64 `gorm:"column:kreatinin"`
	AsamUrat     *float64 `gorm:"column:asam_urat"`
	Natrium      *float64 `gorm:"column:natrium"`
	Kalium       *float64 `gorm:"column:kalium"`
	Klorida      *float64 `gorm:"column:klorida"`
	Waktu        string   `gorm:"column:waktu"`
	Visible      int      `gorm:"column:visible"`
}

func (PemeriksaanLaboratoriumModel) TableName() string {
	return "pemeriksaan_laboratorium"
}
