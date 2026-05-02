package pemeriksaan_penunjang_bedah

import entity "rme/internal/entity/pemeriksaan_penunjang_bedah"

type Repository interface {
	Create(d *entity.PemeriksaanPenunjangBedah) error
	GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
	HideByPatient(idPasien int) error
}
