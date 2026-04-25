package pemeriksaan_ekg

import "time"

type PemeriksaanEkg struct {
	ID                 int
	IDPasien           int
	IDDokter           int
	DetakJantung       *int   // heart rate (HR)
	Irama              string // rhythm (normal/aritmia)
	PRInterval         *float64
	QRSDuration        *float64
	QTTcInterval       *float64
	AxisJantung        string
	STElevationDepress string
	TWaveAbnormality   string
	InterpretasiDokter string
	DateMake           time.Time
	DateUpdate         time.Time
	Visible            int
	NamaPasien         string
	NamaDokter         string
}
