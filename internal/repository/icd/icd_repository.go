package icd

type Repository interface {
	GetAll() ([]*IcdRow, error)
}

type IcdRow struct {
	ID   int
	Kode string
	Nama string
}
