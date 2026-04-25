package pemeriksaan_ekg

import "time"

type PemeriksaanEkgModel struct {
	ID                 int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien           int       `gorm:"column:id_pasien"`
	IDDokter           int       `gorm:"column:id_dokter"`
	DetakJantung       *int      `gorm:"column:detak_jantung"`
	Irama              string    `gorm:"column:irama"`
	PRInterval         *float64  `gorm:"column:pr_interval"`
	QRSDuration        *float64  `gorm:"column:qrs_duration"`
	QTTcInterval       *float64  `gorm:"column:qt_qtc_interval"`
	AxisJantung        string    `gorm:"column:axis_jantung"`
	STElevationDepress string    `gorm:"column:st_elevation_depression"`
	TWaveAbnormality   string    `gorm:"column:t_wave_abnormality"`
	InterpretasiDokter string    `gorm:"column:interpretasi_dokter"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	Visible            int       `gorm:"column:visible"`
}

func (PemeriksaanEkgModel) TableName() string {
	return "pemeriksaan_ekg"
}
