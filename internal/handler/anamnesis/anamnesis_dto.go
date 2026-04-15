package anamnesis

// Request/response DTOs for anamnesis handlers.

import "time"

// AnamnesisCreateRequest merepresentasikan payload JSON untuk membuat data anamnesis baru.
type AnamnesisCreateRequest struct {
	IDPasien                int    `json:"idPasien"`
	IDDokter                int    `json:"idDokter"`
	Text                    string `json:"text"`
	DateMake                string `json:"dateMake"`
	DateUpdate              string `json:"dateUpdate"`
	IDDataKlinik            int    `json:"idDataKlinik"`
	RiwayatPengobatan       string `json:"riwayatPengobatan"`
	RiwayatKeluarga         string `json:"riwayatKeluarga"`
	RiwayatPekerjaan        string `json:"riwayatPekerjaan"`
	RiwayatAutoanamnesis    string `json:"riwayatAutoanamnesis"`
	RiwayatPenyakitDahulu   string `json:"riwayatPenyakitDahulu"`
	RiwayatPenyakitSekarang string `json:"riwayatPenyakitSekarang"`
	RiwayatPenyakitLain     string `json:"riwayatPenyakitLain"`
	HubunganPasien          string `json:"hubunganPasien"`
	RiwayatAnestesiBedah    string `json:"riwayatAnestesiBedah"`
	RiwayatKeluhanUtama     string `json:"riwayatKeluhanUtama"`
	StatusKehamilan         string `json:"statusKehamilan"`
	KeluhanTambahan         string `json:"keluhanTambahan"`
	Catatan                 string `json:"catatan"`
}

// AnamnesisResponse merepresentasikan data anamnesis yang dikembalikan ke client.
type AnamnesisResponse struct {
	IDPasien                int       `json:"id_pasien"`
	IDDokter                int       `json:"id_dokter"`
	Text                    string    `json:"text"`
	DateMake                time.Time `json:"date_make"`
	DateUpdate              time.Time `json:"date_update"`
	IDDataKlinik            int       `json:"id_data_klinik"`
	RiwayatPengobatan       string    `json:"riwayat_pengobatan"`
	RiwayatKeluarga         string    `json:"riwayat_keluarga"`
	RiwayatPekerjaan        string    `json:"riwayat_pekerjaan"`
	RiwayatAutoanamnesis    string    `json:"riwayat_autoanamnesis"`
	RiwayatPenyakitDahulu   string    `json:"riwayat_penyakit_dahulu"`
	RiwayatPenyakitSekarang string    `json:"riwayat_penyakit_sekarang"`
	RiwayatPenyakitLain     string    `json:"riwayat_penyakit_lain"`
	HubunganPasien          string    `json:"hubungan_pasien"`
	RiwayatAnestesiBedah    string    `json:"riwayat_anestesi_bedah"`
	RiwayatKeluhanUtama     string    `json:"riwayat_keluhan_utama"`
	StatusKehamilan         string    `json:"status_kehamilan"`
	KeluhanTambahan         string    `json:"keluhan_tambahan"`
	Catatan                 string    `json:"catatan"`
	Visible                 int       `json:"visible"`
}
