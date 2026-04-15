package anamnesis

import "time"

type AnamnesisModel struct {
	ID                      int       `gorm:"primaryKey;autoIncrement;column:id_anamnesis"`
	IDPasien                int       `gorm:"column:id_pasien"`
	IDDokter                int       `gorm:"column:id_dokter"`
	Text                    string    `gorm:"column:text"`
	DateMake                time.Time `gorm:"column:date_make"`
	DateUpdate              time.Time `gorm:"column:date_update"`
	IDDataKlinik            int       `gorm:"column:id_data_klinik"`
	RiwayatPengobatan       string    `gorm:"column:riwayat_pengobatan"`
	RiwayatKeluarga         string    `gorm:"column:riwayat_keluarga"`
	RiwayatPekerjaan        string    `gorm:"column:riwayat_pekerjaan"`
	RiwayatAutoanamnesis    string    `gorm:"column:riwayat_autoanamnesis"`
	RiwayatPenyakitDahulu   string    `gorm:"column:riwayat_penyakit_dahulu"`
	RiwayatPenyakitSekarang string    `gorm:"column:riwayat_penyakit_sekarang"`
	RiwayatPenyakitLain     string    `gorm:"column:riwayat_penyakit_lain"`
	HubunganPasien          string    `gorm:"column:hubungan_pasien"`
	RiwayatAnestesiBedah    string    `gorm:"column:riwayat_anestesi_bedah"`
	RiwayatKeluhanUtama     string    `gorm:"column:riwayat_keluhan_utama"`
	StatusKehamilan         string    `gorm:"column:status_kehamilan"`
	KeluhanTambahan         string    `gorm:"column:keluhan_tambahan"`
	Catatan                 string    `gorm:"column:catatan"`

	Visible int `gorm:"column:visible"`
}

func (AnamnesisModel) TableName() string {
	return "anamnesis"
}
