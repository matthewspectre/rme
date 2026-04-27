package pemeriksaan_fungsi_organ

type PemeriksaanFungsiOrganCreateRequest struct {
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	Tanggal       string `json:"tanggal"` // RFC3339 optional
	GangguanBAB   bool   `json:"gangguanBAB"`
	GangguanBAK   bool   `json:"gangguanBAK"`
	MualMuntah    bool   `json:"mualMuntah"`
	Demam         bool   `json:"demam"`
	Perdarahan    bool   `json:"perdarahan"`
	PenurunanBB   bool   `json:"penurunanBB"`
	GangguanGerak bool   `json:"gangguanGerak"`
	Catatan       string `json:"catatan"`
}

type PemeriksaanFungsiOrganResponse struct {
	ID            int    `json:"id"`
	IDPasien      int    `json:"id_pasien"`
	IDDokter      int    `json:"id_dokter"`
	Tanggal       string `json:"tanggal"`
	GangguanBAB   bool   `json:"gangguanBAB"`
	GangguanBAK   bool   `json:"gangguanBAK"`
	MualMuntah    bool   `json:"mualMuntah"`
	Demam         bool   `json:"demam"`
	Perdarahan    bool   `json:"perdarahan"`
	PenurunanBB   bool   `json:"penurunanBB"`
	GangguanGerak bool   `json:"gangguanGerak"`
	Catatan       string `json:"catatan"`
}
