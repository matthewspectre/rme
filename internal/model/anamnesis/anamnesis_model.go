package anamnesis

import "time"

type AnamnesisModel struct {
	ID                    int       `gorm:"primaryKey;autoIncrement;column:id_anamnesis"`
	IDPasien              int       `gorm:"column:id_pasien"`
	IDDokter              int       `gorm:"column:id_dokter"`
	Text                  string    `gorm:"column:text"`
	DateMake              time.Time `gorm:"column:date_make"`
	DateUpdate            time.Time `gorm:"column:date_update"`
	IDDataKlinik          int       `gorm:"column:id_data_klinik"`
	RiwayatPengobatan     string    `gorm:"column:riwayat_pengobatan"`
	RiwayatKeluarga       string    `gorm:"column:riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `gorm:"column:riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `gorm:"column:riwayat_penyakit_lain"`
	StatusKehamilan       string    `gorm:"column:status_kehamilan"`
	KeluhanTambahan       string    `gorm:"column:keluhan_tambahan"`

	Visible int `gorm:"column:visible"`
}

func (AnamnesisModel) TableName() string {
	return "anamnesis"
}
