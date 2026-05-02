package pemeriksaan_penunjang_bedah

import "time"

type PemeriksaanPenunjangBedahModel struct {
	ID            int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien      int       `gorm:"column:id_pasien"`
	IDDokter      int       `gorm:"column:id_dokter"`
	ButuhUSG      bool      `gorm:"column:butuh_usg"`
	ButuhRontgen  bool      `gorm:"column:butuh_rontgen"`
	ButuhCTScan   bool      `gorm:"column:butuh_ctscan"`
	ButuhBiopsi   bool      `gorm:"column:butuh_biopsi"`
	StatusOperasi string    `gorm:"column:status_operasi"`
	JenisTindakan string    `gorm:"column:jenis_tindakan"`
	Prioritas     string    `gorm:"column:prioritas"`
	CatatanBedah  string    `gorm:"column:catatan_bedah"`
	JadwalBedah   time.Time `gorm:"column:jadwal_bedah"`
	Visible       int       `gorm:"column:visible"`
	DateMake      time.Time `gorm:"column:date_make"`
	DateUpdate    time.Time `gorm:"column:date_update"`
}

func (PemeriksaanPenunjangBedahModel) TableName() string {
	return "pemeriksaan_penunjang_bedah"
}
