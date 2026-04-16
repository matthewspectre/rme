package doctor

// Request/response DTOs for doctor handlers.

type DoctorCreateRequest struct {
	IDUser       int    `json:"idUser"`
	NamaDokter   string `json:"namaDokter"`
	Poli         string `json:"poli"`
	NomorTelepon string `json:"nomorTelepon"`
	IDDataKlinik int    `json:"idDataKlinik"`
}

type DoctorResponse struct {
	ID           int    `json:"id"`
	IDUser       int    `json:"idUser"`
	NamaDokter   string `json:"namaDokter"`
	Poli         string `json:"poli"`
	NomorTelepon string `json:"nomorTelepon"`
	IDDataKlinik int    `json:"idDataKlinik"`
}
