package icd

// GORM model for icd10 table
type IcdModel struct {
	ID   int    `gorm:"primaryKey;autoIncrement;column:id"`
	Kode string `gorm:"column:kode"`
	Nama string `gorm:"column:nama"`
}

func (IcdModel) TableName() string {
	return "icd10"
}
