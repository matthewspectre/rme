package diagnosis

import "time"

type Diagnosis struct {
	IDDiagnosis      int
	IDPasien         int
	IDDokter         int
	Tanggal          time.Time
	KodeIcdUtama     string
	KodeIcdSekunder  []string
	DiagnosisBanding []string
	NamaIcdUtama     string
	NamaIcdSekunder  []string
	Status           string
	DasarDiagnosis   []string
	Catatan          string
	Visible          int
}
