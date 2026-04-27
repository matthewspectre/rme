package diagnosis

import "time"

type DiagnosisModel struct {
	IDDiagnosis      int       `gorm:"primaryKey;autoIncrement;column:id_diagnosis"`
	IDPasien         int       `gorm:"column:id_pasien"`
	IDDokter         int       `gorm:"column:id_dokter"`
	Tanggal          time.Time `gorm:"column:tanggal"`
	KodeIcdUtama     string    `gorm:"column:kode_icd_utama"`
	KodeIcdSekunder  string    `gorm:"column:kode_icd_sekunder"` // JSON-encoded array
	DiagnosisBanding string    `gorm:"column:diagnosis_banding"` // JSON-encoded array
	Status           string    `gorm:"column:status"`
	DasarDiagnosis   string    `gorm:"column:dasar_diagnosis"` // JSON-encoded array
	Catatan          string    `gorm:"column:catatan"`
	Visible          int       `gorm:"column:visible"`
	DateMake         time.Time `gorm:"column:date_make"`
	DateUpdate       time.Time `gorm:"column:date_update"`
}

func (DiagnosisModel) TableName() string {
	return "diagnosis"
}
