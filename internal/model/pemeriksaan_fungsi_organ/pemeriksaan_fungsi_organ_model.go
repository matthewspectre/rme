package pemeriksaan_fungsi_organ

import "time"

type PemeriksaanFungsiOrganModel struct {
	ID            int       `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien      int       `gorm:"column:id_pasien"`
	IDDokter      int       `gorm:"column:id_dokter"`
	Tanggal       time.Time `gorm:"column:tanggal"`
	GangguanBAB   bool      `gorm:"column:gangguan_bab"`
	GangguanBAK   bool      `gorm:"column:gangguan_bak"`
	MualMuntah    bool      `gorm:"column:mual_muntah"`
	Demam         bool      `gorm:"column:demam"`
	Perdarahan    bool      `gorm:"column:perdarahan"`
	PenurunanBB   bool      `gorm:"column:penurunan_bb"`
	GangguanGerak bool      `gorm:"column:gangguan_gerak"`
	Catatan       string    `gorm:"column:catatan"`
	Visible       int       `gorm:"column:visible"`
	DateMake      time.Time `gorm:"column:date_make"`
	DateUpdate    time.Time `gorm:"column:date_update"`
}

func (PemeriksaanFungsiOrganModel) TableName() string {
	return "pemeriksaan_fungsi_organ"
}
