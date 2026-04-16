package poli

// GORM model untuk tabel daftar_poli
type PoliModel struct {
	ID             int    `gorm:"primaryKey;autoIncrement;column:id"`
	NamaPoli       string `gorm:"column:nama_poli"`
	Deskripsi      string `gorm:"column:deskripsi"`
	IDDataKlinik   int    `gorm:"column:id_data_klinik"`
	Aktif          bool   `gorm:"column:aktif"`
	DibuatPada     string `gorm:"column:dibuat_pada"`
	DiperbaruiPada string `gorm:"column:diperbarui_pada"`
}

func (PoliModel) TableName() string {
	return "daftar_poli"
}
