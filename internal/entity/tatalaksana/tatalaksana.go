package tatalaksana

import "time"

// Domain entity untuk tatalaksana.
type Tatalaksana struct {
	ID         int
	IDPasien   int
	IDDokter   int
	DateMake   time.Time
	DateUpdate time.Time
	NamaObat   string
	Dosis      string
	Frekuensi  string
	Durasi     string
	CaraPakai  string
	Catatan    string
	Visible    int
	NamaPasien string
	NamaDokter string
}
