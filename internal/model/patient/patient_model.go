package patient

import "time"

// GORM model for patients table.

type PatientModel struct {
	ID             int       `gorm:"primaryKey;autoIncrement;column:id"`
	Name           string    `gorm:"column:name"`
	AdmissionDate  time.Time `gorm:"column:admission_date"`
	NIK            string    `gorm:"column:nik"`
	Gender         string    `gorm:"column:gender"`
	BloodType      string    `gorm:"column:blood_type"`
	BirthPlaceDate string    `gorm:"column:birth_place_date"`
	Phone          string    `gorm:"column:phone"`
	Address        string    `gorm:"column:address"`
	Category       string    `gorm:"column:category"`
	Job            string    `gorm:"column:job"`
	IDDataKlinik   int       `gorm:"column:id_data_klinik"`
}

func (PatientModel) TableName() string {
	return "patients"
}
