package pemeriksaan_fungsi_organ

import (
	"errors"
	entity "rme/internal/entity/pemeriksaan_fungsi_organ"
	model "rme/internal/model/pemeriksaan_fungsi_organ"
	"time"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.PemeriksaanFungsiOrgan) error {
	mm := model.PemeriksaanFungsiOrganModel{
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		Tanggal:       d.Tanggal,
		GangguanBAB:   d.GangguanBAB,
		GangguanBAK:   d.GangguanBAK,
		MualMuntah:    d.MualMuntah,
		Demam:         d.Demam,
		Perdarahan:    d.Perdarahan,
		PenurunanBB:   d.PenurunanBB,
		GangguanGerak: d.GangguanGerak,
		Catatan:       d.Catatan,
		Visible:       1,
		DateMake:      d.Tanggal,
		DateUpdate:    d.Tanggal,
	}
	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.ID = mm.ID
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.PemeriksaanFungsiOrgan, error) {
	var mm model.PemeriksaanFungsiOrganModel
	if err := m.DB.Where("id = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	e := &entity.PemeriksaanFungsiOrgan{
		ID:            mm.ID,
		IDPasien:      mm.IDPasien,
		IDDokter:      mm.IDDokter,
		Tanggal:       mm.Tanggal,
		GangguanBAB:   mm.GangguanBAB,
		GangguanBAK:   mm.GangguanBAK,
		MualMuntah:    mm.MualMuntah,
		Demam:         mm.Demam,
		Perdarahan:    mm.Perdarahan,
		PenurunanBB:   mm.PenurunanBB,
		GangguanGerak: mm.GangguanGerak,
		Catatan:       mm.Catatan,
		Visible:       mm.Visible,
	}
	return e, nil
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.PemeriksaanFungsiOrgan, error) {
	var list []model.PemeriksaanFungsiOrganModel
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
	out := make([]*entity.PemeriksaanFungsiOrgan, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.PemeriksaanFungsiOrgan{
			ID:            mm.ID,
			IDPasien:      mm.IDPasien,
			IDDokter:      mm.IDDokter,
			Tanggal:       mm.Tanggal,
			GangguanBAB:   mm.GangguanBAB,
			GangguanBAK:   mm.GangguanBAK,
			MualMuntah:    mm.MualMuntah,
			Demam:         mm.Demam,
			Perdarahan:    mm.Perdarahan,
			PenurunanBB:   mm.PenurunanBB,
			GangguanGerak: mm.GangguanGerak,
			Catatan:       mm.Catatan,
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
		case "tanggal":
			switch val := v.(type) {
			case string:
				if t, err := time.Parse(time.RFC3339, val); err == nil {
					dbUpdates["tanggal"] = t
				} else {
					dbUpdates["tanggal"] = val
				}
			case time.Time:
				dbUpdates["tanggal"] = val
			default:
				dbUpdates["tanggal"] = v
			}
		default:
			dbUpdates[k] = v
		}
	}
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.PemeriksaanFungsiOrganModel{}).Where("id = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.PemeriksaanFungsiOrganModel{}).
		Where("id = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
