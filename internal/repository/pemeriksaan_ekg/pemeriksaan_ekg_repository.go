package pemeriksaan_ekg

import "rme/internal/entity/pemeriksaan_ekg"

type Repository interface {
	Create(p *pemeriksaan_ekg.PemeriksaanEkg) (int, error)
	GetAll(idPasien *int, idDokter *int) ([]*PemeriksaanEkgWithNames, error)
	GetByID(id int) (*PemeriksaanEkgWithNames, error)
	Update(id int, updates map[string]interface{}) (*PemeriksaanEkgWithNames, error)
	Delete(id int) error
}

type PemeriksaanEkgWithNames struct {
	ID                 int
	IDPasien           int
	NamaPasien         string
	IDDokter           int
	NamaDokter         string
	Visible            int
	DetakJantung       *int
	Irama              string
	PRInterval         *float64
	QRSDuration        *float64
	QTTcInterval       *float64
	AxisJantung        string
	STElevationDepress string
	TWaveAbnormality   string
	InterpretasiDokter string
	DateMake           string
}
