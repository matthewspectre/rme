package rujukulang

import "time"

// DTO untuk create / response rujuk ulang
type RujukUlangCreateRequest struct {
	IDPasien           int    `json:"idPasien"`
	IDDokter           int    `json:"idDokter"`
	PoliAsal           int    `json:"poliAsal"`
	PoliTujuan         int    `json:"poliTujuan"`
	DateMake           string `json:"dateMake"`
	DateUpdate         string `json:"dateUpdate"`
	DiagnosisSementara string `json:"diagnosis_sementara"`
	Catatan            string `json:"catatan"`
}

type RujukUlangResponse struct {
	ID                 int       `json:"id_rujuk_ulang"`
	IDPasien           int       `json:"id_pasien"`
	PoliAsal           int       `json:"poli_asal"`
	PoliTujuan         int       `json:"poli_tujuan"`
	DiagnosisSementara string    `json:"diagnosis_sementara"`
	Catatan            string    `json:"catatan"`
	IDDokter           int       `json:"id_dokter"`
	DateMake           time.Time `json:"date_make"`
	DateUpdate         time.Time `json:"date_update"`
	Visible            int       `json:"visible"`
	NamaPasien         string    `json:"nama_pasien"`
	NamaPoliAsal       string    `json:"nama_poli_asal"`
	NamaPoliTujuan     string    `json:"nama_poli_tujuan"`
	NamaDokter         string    `json:"nama_dokter"`
}
