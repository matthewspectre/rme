package pemeriksaanvital

import "time"

type PemeriksaanVitalModel struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id_pemeriksaan_vital"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
	TekananDarah   string    `gorm:"column:tekanan_darah"`
	DenyutNadi     int       `gorm:"column:denyut_nadi"`
	SuhuTubuh      float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas int       `gorm:"column:frekuensi_napas"`
	BeratBadan     float64   `gorm:"column:berat_badan"`
	TinggiBadan    float64   `gorm:"column:tinggi_badan"`
	Visible        int       `gorm:"column:visible"`
}

func (PemeriksaanVitalModel) TableName() string {
	return "pemeriksaan_vital"
}
