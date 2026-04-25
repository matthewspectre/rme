package antrian

// GORM model untuk tabel antrian_pasien
type AntrianModel struct {
	ID           int    `gorm:"primaryKey;autoIncrement;column:id"`
	IDPasien     int    `gorm:"column:id_pasien"`
	IDDokter     int    `gorm:"column:id_dokter"`
	NomorAntrian int    `gorm:"column:nomor_antrian"`
	IDPoli       int    `gorm:"column:id_poli"`
	Waktu        string `gorm:"column:waktu"`
	Status       int    `gorm:"column:status"`
}

func (AntrianModel) TableName() string {
	return "antrian_pasien"
}
