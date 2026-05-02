package pemeriksaan_penunjang_bedah

import (
	"errors"
	"time"

	entity "rme/internal/entity/pemeriksaan_penunjang_bedah"
	model "rme/internal/model/pemeriksaan_penunjang_bedah"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.PemeriksaanPenunjangBedah) error {
	mm := model.PemeriksaanPenunjangBedahModel{
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		ButuhUSG:      d.ButuhUSG,
		ButuhRontgen:  d.ButuhRontgen,
		ButuhCTScan:   d.ButuhCTScan,
		ButuhBiopsi:   d.ButuhBiopsi,
		StatusOperasi: d.StatusOperasi,
		JenisTindakan: d.JenisTindakan,
		Prioritas:     d.Prioritas,
		CatatanBedah:  d.CatatanBedah,
		JadwalBedah:   d.JadwalBedah,
		Visible:       1,
		DateMake:      time.Now(),
		DateUpdate:    time.Now(),
	}
	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.ID = mm.ID
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.PemeriksaanPenunjangBedah, error) {
	var mm model.PemeriksaanPenunjangBedahModel
	if err := m.DB.Where("id = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	e := &entity.PemeriksaanPenunjangBedah{
		ID:            mm.ID,
		IDPasien:      mm.IDPasien,
		IDDokter:      mm.IDDokter,
		ButuhUSG:      mm.ButuhUSG,
		ButuhRontgen:  mm.ButuhRontgen,
		ButuhCTScan:   mm.ButuhCTScan,
		ButuhBiopsi:   mm.ButuhBiopsi,
		StatusOperasi: mm.StatusOperasi,
		JenisTindakan: mm.JenisTindakan,
		Prioritas:     mm.Prioritas,
		CatatanBedah:  mm.CatatanBedah,
		JadwalBedah:   mm.JadwalBedah,
		Visible:       mm.Visible,
	}
	return e, nil
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanPenunjangBedah, error) {
	var list []model.PemeriksaanPenunjangBedahModel
	q := m.DB.Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.PemeriksaanPenunjangBedah, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanPenunjangBedah{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			ButuhUSG:      mm.ButuhUSG,
			ButuhRontgen:  mm.ButuhRontgen,
			ButuhCTScan:   mm.ButuhCTScan,
			ButuhBiopsi:   mm.ButuhBiopsi,
			StatusOperasi: mm.StatusOperasi,
			JenisTindakan: mm.JenisTindakan,
			Prioritas:     mm.Prioritas,
			CatatanBedah:  mm.CatatanBedah,
			JadwalBedah:   mm.JadwalBedah,
			Visible:       mm.Visible,
		})
	}
	return out, nil
}

func (m *mysqlRepo) Update(id int, updates map[string]interface{}) error {
	if len(updates) == 0 {
		return nil
	}
	dbUpdates := map[string]interface{}{}
	for k, v := range updates {
		switch k {
		case "jadwalBedah", "jadwal_bedah":
			switch val := v.(type) {
			case string:
				if t, err := time.Parse(time.RFC3339, val); err == nil {
					dbUpdates["jadwal_bedah"] = t
				} else {
					dbUpdates["jadwal_bedah"] = val
				}
			case time.Time:
				dbUpdates["jadwal_bedah"] = val
			default:
				dbUpdates["jadwal_bedah"] = v
			}
		case "statusOperasi", "status_operasi":
			dbUpdates["status_operasi"] = v
		case "jenisTindakan", "jenis_tindakan":
			dbUpdates["jenis_tindakan"] = v
		case "prioritas":
			dbUpdates["prioritas"] = v
		case "catatanBedah", "catatan_bedah":
			dbUpdates["catatan_bedah"] = v
		case "butuhUSG", "butuh_usg":
			dbUpdates["butuh_usg"] = v
		case "butuhRontgen", "butuh_rontgen":
			dbUpdates["butuh_rontgen"] = v
		case "butuhCTScan", "butuh_ctscan":
			dbUpdates["butuh_ctscan"] = v
		case "butuhBiopsi", "butuh_biopsi":
			dbUpdates["butuh_biopsi"] = v
		case "id_dokter":
			dbUpdates["id_dokter"] = v
		default:
			dbUpdates[k] = v
		}
	}
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).Where("id = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).
		Where("id = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) HideByPatient(idPasien int) error {
	if err := m.DB.Model(&model.PemeriksaanPenunjangBedahModel{}).
		Where("id_pasien = ? AND visible = 1", idPasien).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
