package pemeriksaanvital

import "time"

// DTO untuk create / response pemeriksaan vital
type PemeriksaanVitalCreateRequest struct {
	IDPasien       int     `json:"idPasien"`
	IDDokter       int     `json:"idDokter"`
	DateMake       string  `json:"dateMake"`
	DateUpdate     string  `json:"dateUpdate"`
	TekananDarah   string  `json:"tekanan_darah"`
	DenyutNadi     int     `json:"denyut_nadi"`
	SuhuTubuh      float64 `json:"suhu_tubuh"`
	FrekuensiNapas int     `json:"frekuensi_napas"`
	BeratBadan     float64 `json:"berat_badan"`
	TinggiBadan    float64 `json:"tinggi_badan"`
}

type PemeriksaanVitalResponse struct {
	ID             int       `json:"id_pemeriksaan_vital"`
	IDPasien       int       `json:"id_pasien"`
	IDDokter       int       `json:"id_dokter"`
	DateMake       time.Time `json:"date_make"`
	DateUpdate     time.Time `json:"date_update"`
	TekananDarah   string    `json:"tekanan_darah"`
	DenyutNadi     int       `json:"denyut_nadi"`
	SuhuTubuh      float64   `json:"suhu_tubuh"`
	FrekuensiNapas int       `json:"frekuensi_napas"`
	BeratBadan     float64   `json:"berat_badan"`
	TinggiBadan    float64   `json:"tinggi_badan"`
	Visible        int       `json:"visible"`
	NamaPasien     string    `json:"nama_pasien"`
	NamaDokter     string    `json:"nama_dokter"`
}
