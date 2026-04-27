package pemeriksaan_fungsi_organ

import "time"

type PemeriksaanFungsiOrgan struct {
	ID            int
	IDPasien      int
	IDDokter      int
	Tanggal       time.Time
	GangguanBAB   bool
	GangguanBAK   bool
	MualMuntah    bool
	Demam         bool
	Perdarahan    bool
	PenurunanBB   bool
	GangguanGerak bool
	Catatan       string
	Visible       int
	NamaPasien    string
}
