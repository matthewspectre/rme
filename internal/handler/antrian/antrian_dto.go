package antrian

type AntrianCreateRequest struct {
	IDPasien int `json:"idPasien"`
	IDDokter int `json:"idDokter"`
	IDPoli   int `json:"idPoli"`
}

type AntrianResponse struct {
	ID           int    `json:"id"`
	IDPasien     int    `json:"idPasien"`
	NamaPasien   string `json:"namaPasien"`
	IDDokter     int    `json:"idDokter"`
	NamaDokter   string `json:"namaDokter"`
	NomorAntrian int    `json:"nomorAntrian"`
	IDPoli       int    `json:"idPoli"`
	Waktu        string `json:"waktu"`
}

type AntrianUpdateRequest struct {
	IDPasien     *int `json:"idPasien,omitempty"`
	IDDokter     *int `json:"idDokter,omitempty"`
	IDPoli       *int `json:"idPoli,omitempty"`
	NomorAntrian *int `json:"nomorAntrian,omitempty"`
}
