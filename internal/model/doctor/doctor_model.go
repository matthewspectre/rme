package doctor

// GORM model for doctors table.

type DoctorModel struct {
	ID           int    `gorm:"primaryKey;autoIncrement;column:id"`
	IDUser       int    `gorm:"column:id_user"`
	NamaDokter   string `gorm:"column:nama_dokter"`
	Poli         string `gorm:"column:poli"`
	NomorTelepon string `gorm:"column:nomor_telepon"`
	IDDataKlinik int    `gorm:"column:id_data_klinik"`
}

func (DoctorModel) TableName() string {
	return "dokter"
}
