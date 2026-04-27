package lokalis_bedah

type LokalisBedahCreateRequest struct {
	IDPasien       int    `json:"id_pasien"`
	IDDokter       int    `json:"id_dokter"`
	Tanggal        string `json:"tanggal"` // RFC3339 optional
	LokasiKelainan string `json:"lokasiKelainan"`
	JenisKelainan  string `json:"jenisKelainan"`
	Ukuran         string `json:"ukuran"`
	Warna          string `json:"warna"`
	NyeriTekan     bool   `json:"nyeriTekan"`
	Konsistensi    string `json:"konsistensi"`
	Mobilitas      string `json:"mobilitas"`
	TandaRadang    bool   `json:"tandaRadang"`
	Fluktuasi      bool   `json:"fluktuasi"`
	Catatan        string `json:"catatan"`
}

type LokalisBedahResponse struct {
	ID             int    `json:"id"`
	IDPasien       int    `json:"id_pasien"`
	IDDokter       int    `json:"id_dokter"`
	Tanggal        string `json:"tanggal"`
	LokasiKelainan string `json:"lokasiKelainan"`
	JenisKelainan  string `json:"jenisKelainan"`
	Ukuran         string `json:"ukuran"`
	Warna          string `json:"warna"`
	NyeriTekan     bool   `json:"nyeriTekan"`
	Konsistensi    string `json:"konsistensi"`
	Mobilitas      string `json:"mobilitas"`
	TandaRadang    bool   `json:"tandaRadang"`
	Fluktuasi      bool   `json:"fluktuasi"`
	Catatan        string `json:"catatan"`
}

type LokalisBedahUpdateRequest struct {
	LokasiKelainan *string `json:"lokasiKelainan"`
	JenisKelainan  *string `json:"jenisKelainan"`
	Ukuran         *string `json:"ukuran"`
	Warna          *string `json:"warna"`
	NyeriTekan     *bool   `json:"nyeriTekan"`
	Konsistensi    *string `json:"konsistensi"`
	Mobilitas      *string `json:"mobilitas"`
	TandaRadang    *bool   `json:"tandaRadang"`
	Fluktuasi      *bool   `json:"fluktuasi"`
	Catatan        *string `json:"catatan"`
	Tanggal        *string `json:"tanggal"` // RFC3339
}
