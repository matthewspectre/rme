package poli

type PoliResponse struct {
	ID             int    `json:"id"`
	NamaPoli       string `json:"namaPoli"`
	Deskripsi      string `json:"deskripsi"`
	IDDataKlinik   int    `json:"idDataKlinik"`
	Aktif          bool   `json:"aktif"`
	DibuatPada     string `json:"dibuatPada"`
	DiperbaruiPada string `json:"diperbaruiPada"`
}
