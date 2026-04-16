package doctor

// Domain entity for doctors.

type Doctor struct {
	ID           int
	IDUser       int
	Nama         string
	Poli         string
	NomorTelepon string
	IDDataKlinik int
}
