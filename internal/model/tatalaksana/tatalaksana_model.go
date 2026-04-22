package tatalaksana

import "time"

type TatalaksanaModel struct {
	ID         int       `gorm:"primaryKey;autoIncrement;column:id_tatalaksana"`
	IDPasien   int       `gorm:"column:id_pasien"`
	IDDokter   int       `gorm:"column:id_dokter"`
	DateMake   time.Time `gorm:"column:date_make"`
	DateUpdate time.Time `gorm:"column:date_update"`
	NamaObat   string    `gorm:"column:nama_obat"`
	Dosis      string    `gorm:"column:dosis"`
	Frekuensi  string    `gorm:"column:frekuensi"`
	Durasi     string    `gorm:"column:durasi"`
	CaraPakai  string    `gorm:"column:cara_pakai"`
	Catatan    string    `gorm:"column:catatan"`
	Visible    int       `gorm:"column:visible"`
}

func (TatalaksanaModel) TableName() string {
	return "tatalaksana"
}
