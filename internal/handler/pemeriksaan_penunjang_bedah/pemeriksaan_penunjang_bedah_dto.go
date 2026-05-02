package pemeriksaan_penunjang_bedah

type PemeriksaanPenunjangBedahCreateRequest struct {
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	ButuhUSG      bool   `json:"butuhUSG"`
	ButuhRontgen  bool   `json:"butuhRontgen"`
	ButuhCTScan   bool   `json:"butuhCTScan"`
	ButuhBiopsi   bool   `json:"butuhBiopsi"`
	StatusOperasi string `json:"statusOperasi"`
	JenisTindakan string `json:"jenisTindakan"`
	Prioritas     string `json:"prioritas"`
	CatatanBedah  string `json:"catatanBedah"`
	JadwalBedah   string `json:"jadwalBedah"` // RFC3339 optional
}

type PemeriksaanPenunjangBedahResponse struct {
	ID            int    `json:"id"`
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	ButuhUSG      bool   `json:"butuhUSG"`
	ButuhRontgen  bool   `json:"butuhRontgen"`
	ButuhCTScan   bool   `json:"butuhCTScan"`
	ButuhBiopsi   bool   `json:"butuhBiopsi"`
	StatusOperasi string `json:"statusOperasi"`
	JenisTindakan string `json:"jenisTindakan"`
	Prioritas     string `json:"prioritas"`
	CatatanBedah  string `json:"catatanBedah"`
	JadwalBedah   string `json:"jadwalBedah"`
}
