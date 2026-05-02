package pemeriksaan_penunjang_bedah

import "time"

type PemeriksaanPenunjangBedah struct {
	ID            int
	IDPasien      int
	IDDokter      int
	ButuhUSG      bool
	ButuhRontgen  bool
	ButuhCTScan   bool
	ButuhBiopsi   bool
	StatusOperasi string
	JenisTindakan string
	Prioritas     string
	CatatanBedah  string
	JadwalBedah   time.Time
	Visible       int
	DateMake      time.Time
	DateUpdate    time.Time
}
