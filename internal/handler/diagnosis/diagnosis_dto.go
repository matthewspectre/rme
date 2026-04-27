package diagnosis

import "time"

type DiagnosisCreateRequest struct {
	IDPasien          int                    `json:"id_pasien"`
	IDDokter          int                    `json:"id_dokter"`
	Tanggal           string                 `json:"tanggal"`
	DiagnosisUtama    DiagnosisCodeRequest   `json:"diagnosis_utama"`
	DiagnosisSekunder []DiagnosisCodeRequest `json:"diagnosis_sekunder"`
	DiagnosisBanding  []string               `json:"diagnosis_banding"`
	Status            string                 `json:"status"`
	DasarDiagnosis    []string               `json:"dasar_diagnosis"`
	Catatan           string                 `json:"catatan"`
}

type DiagnosisCodeRequest struct {
	KodeIcd string `json:"kode_icd"`
}

type DiagnosisUpdateRequest struct {
	DiagnosisUtama    *DiagnosisCodeRequest   `json:"diagnosis_utama"`
	DiagnosisSekunder *[]DiagnosisCodeRequest `json:"diagnosis_sekunder"`
	DiagnosisBanding  *[]string               `json:"diagnosis_banding"`
	Status            *string                 `json:"status"`
	DasarDiagnosis    *[]string               `json:"dasar_diagnosis"`
	Catatan           *string                 `json:"catatan"`
	Tanggal           *string                 `json:"tanggal"` // RFC3339
}

type DiagnosisResponse struct {
	IDDiagnosis       int                     `json:"id_diagnosis"`
	IDPasien          int                     `json:"id_pasien"`
	IDDokter          int                     `json:"id_dokter"`
	Tanggal           time.Time               `json:"tanggal"`
	DiagnosisUtama    DiagnosisCodeResponse   `json:"diagnosis_utama"`
	DiagnosisSekunder []DiagnosisCodeResponse `json:"diagnosis_sekunder"`
	DiagnosisBanding  []string                `json:"diagnosis_banding"`
	Status            string                  `json:"status"`
	DasarDiagnosis    []string                `json:"dasar_diagnosis"`
	Catatan           string                  `json:"catatan"`
}

type DiagnosisCodeResponse struct {
	KodeIcd string `json:"kode_icd"`
	Nama    string `json:"nama"`
}
