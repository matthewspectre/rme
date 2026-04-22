package tatalaksana

import "time"

// DTO untuk create / response tatalaksana
type TatalaksanaCreateRequest struct {
	IDPasien   int    `json:"idPasien"`
	IDDokter   int    `json:"idDokter"`
	DateMake   string `json:"dateMake"`
	DateUpdate string `json:"dateUpdate"`
	NamaObat   string `json:"nama_obat"`
	Dosis      string `json:"dosis"`
	Frekuensi  string `json:"frekuensi"`
	Durasi     string `json:"durasi"`
	CaraPakai  string `json:"cara_pakai"`
	Catatan    string `json:"catatan"`
}

type TatalaksanaResponse struct {
	ID         int       `json:"id_tatalaksana"`
	IDPasien   int       `json:"id_pasien"`
	IDDokter   int       `json:"id_dokter"`
	DateMake   time.Time `json:"date_make"`
	DateUpdate time.Time `json:"date_update"`
	NamaObat   string    `json:"nama_obat"`
	Dosis      string    `json:"dosis"`
	Frekuensi  string    `json:"frekuensi"`
	Durasi     string    `json:"durasi"`
	CaraPakai  string    `json:"cara_pakai"`
	Catatan    string    `json:"catatan"`
	Visible    int       `json:"visible"`
	NamaPasien string    `json:"nama_pasien"`
	NamaDokter string    `json:"nama_dokter"`
}
