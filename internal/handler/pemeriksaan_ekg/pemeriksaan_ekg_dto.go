package pemeriksaan_ekg

// DTOs for handler
type CreateRequest struct {
	IDPasien           int      `json:"idPasien" binding:"required"`
	IDDokter           int      `json:"idDokter"`
	DateMake           string   `json:"dateMake"`
	DateUpdate         string   `json:"dateUpdate"`
	DetakJantung       *int     `json:"detakJantung"`
	Irama              string   `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	AxisJantung        string   `json:"axisJantung"`
	STElevationDepress string   `json:"st_elevation_depression"`
	TWaveAbnormality   string   `json:"t_wave_abnormality"`
	InterpretasiDokter string   `json:"interpretasi_dokter"`
}

type UpdateRequest struct {
	DetakJantung       *int     `json:"detakJantung"`
	Irama              *string  `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	AxisJantung        *string  `json:"axisJantung"`
	STElevationDepress *string  `json:"st_elevation_depression"`
	TWaveAbnormality   *string  `json:"t_wave_abnormality"`
	InterpretasiDokter *string  `json:"interpretasi_dokter"`
	Visible            *int     `json:"visible"`
}

type Response struct {
	ID                 int      `json:"id"`
	IDPasien           int      `json:"idPasien"`
	NamaPasien         string   `json:"namaPasien"`
	IDDokter           int      `json:"idDokter"`
	NamaDokter         string   `json:"namaDokter"`
	Visible            int      `json:"visible"`
	DetakJantung       *int     `json:"detakJantung"`
	Irama              string   `json:"irama"`
	PRInterval         *float64 `json:"prInterval"`
	QRSDuration        *float64 `json:"qrsDuration"`
	QTTcInterval       *float64 `json:"qt_qtc_interval"`
	AxisJantung        string   `json:"axisJantung"`
	STElevationDepress string   `json:"st_elevation_depression"`
	TWaveAbnormality   string   `json:"t_wave_abnormality"`
	InterpretasiDokter string   `json:"interpretasi_dokter"`
	DateMake           string   `json:"dateMake"`
}
