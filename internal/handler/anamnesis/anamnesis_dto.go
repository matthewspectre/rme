package anamnesis

// Request/response DTOs for anamnesis handlers.

import "time"

// AnamnesisCreateRequest merepresentasikan payload JSON untuk membuat data anamnesis baru.
type AnamnesisCreateRequest struct {
	IDPasien              int    `json:"idPasien"`
	IDDokter              int    `json:"idDokter"`
	Text                  string `json:"text"`
	DateMake              string `json:"dateMake"`
	DateUpdate            string `json:"dateUpdate"`
	IDDataKlinik          int    `json:"idDataKlinik"`
	RiwayatPengobatan     string `json:"riwayatPengobatan"`
	RiwayatKeluarga       string `json:"riwayatKeluarga"`
	RiwayatPenyakitDahulu string `json:"riwayatPenyakitDahulu"`
	RiwayatPenyakitLain   string `json:"riwayatPenyakitLain"`
	StatusKehamilan       string `json:"statusKehamilan"`
	KeluhanTambahan       string `json:"keluhanTambahan"`
}

// AnamnesisUpdateRequest merepresentasikan payload untuk PATCH /anamnesis/:id
type AnamnesisUpdateRequest struct {
	Text                  *string `json:"text,omitempty"`
	RiwayatPengobatan     *string `json:"riwayat_pengobatan,omitempty"`
	RiwayatKeluarga       *string `json:"riwayat_keluarga,omitempty"`
	RiwayatPenyakitDahulu *string `json:"riwayat_penyakit_dahulu,omitempty"`
	RiwayatPenyakitLain   *string `json:"riwayat_penyakit_lain,omitempty"`
	StatusKehamilan       *string `json:"status_kehamilan,omitempty"`
	KeluhanTambahan       *string `json:"keluhan_tambahan,omitempty"`
}

// AnamnesisResponse merepresentasikan data anamnesis yang dikembalikan ke client.
type AnamnesisResponse struct {
	ID                    int       `json:"id_anamnesis"`
	IDPasien              int       `json:"id_pasien"`
	IDDokter              int       `json:"id_dokter"`
	NamaPasien            string    `json:"nama_pasien"`
	Text                  string    `json:"text"`
	DateMake              time.Time `json:"date_make"`
	DateUpdate            time.Time `json:"date_update"`
	IDDataKlinik          int       `json:"id_data_klinik"`
	RiwayatPengobatan     string    `json:"riwayat_pengobatan"`
	RiwayatKeluarga       string    `json:"riwayat_keluarga"`
	RiwayatPenyakitDahulu string    `json:"riwayat_penyakit_dahulu"`
	RiwayatPenyakitLain   string    `json:"riwayat_penyakit_lain"`
	StatusKehamilan       string    `json:"status_kehamilan"`
	KeluhanTambahan       string    `json:"keluhan_tambahan"`
	Visible               int       `json:"visible"`
}
