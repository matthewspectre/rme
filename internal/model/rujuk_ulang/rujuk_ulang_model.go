package rujukulang

import "time"

type RujukUlangModel struct {
	ID                 int       `gorm:"primaryKey;autoIncrement;column:id_rujuk_ulang"`
	IDPasien           int       `gorm:"column:id_pasien"`
	IDDokter           int       `gorm:"column:id_dokter"`
	PoliAsal           int       `gorm:"column:poli_asal"`
	PoliTujuan         int       `gorm:"column:poli_tujuan"`
	DiagnosisSementara string    `gorm:"column:diagnosis_sementara"`
	Catatan            string    `gorm:"column:catatan"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	Visible            int       `gorm:"column:visible"`
}

func (RujukUlangModel) TableName() string {
	return "rujuk_ulang"
}
