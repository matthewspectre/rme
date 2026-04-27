package lokalis_bedah

import (
	"errors"
	"time"

	entity "rme/internal/entity/lokalis_bedah"
	model "rme/internal/model/lokalis_bedah"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.LokalisBedah) error {
	mm := model.LokalisBedahModel{
		IDPasien:       d.IDPasien,
		IDDokter:       d.IDDokter,
		Tanggal:        d.Tanggal,
		LokasiKelainan: d.LokasiKelainan,
		JenisKelainan:  d.JenisKelainan,
		Ukuran:         d.Ukuran,
		Warna:          d.Warna,
		NyeriTekan:     d.NyeriTekan,
		Konsistensi:    d.Konsistensi,
		Mobilitas:      d.Mobilitas,
		TandaRadang:    d.TandaRadang,
		Fluktuasi:      d.Fluktuasi,
		Catatan:        d.Catatan,
		Visible:        1,
		DateMake:       d.Tanggal,
		DateUpdate:     d.Tanggal,
	}
	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.ID = mm.ID
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.LokalisBedah, error) {
	var mm model.LokalisBedahModel
	if err := m.DB.Where("id = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	e := &entity.LokalisBedah{
		ID:             mm.ID,
		IDPasien:       mm.IDPasien,
		IDDokter:       mm.IDDokter,
		Tanggal:        mm.Tanggal,
		LokasiKelainan: mm.LokasiKelainan,
		JenisKelainan:  mm.JenisKelainan,
		Ukuran:         mm.Ukuran,
		Warna:          mm.Warna,
		NyeriTekan:     mm.NyeriTekan,
		Konsistensi:    mm.Konsistensi,
		Mobilitas:      mm.Mobilitas,
		TandaRadang:    mm.TandaRadang,
		Fluktuasi:      mm.Fluktuasi,
		Catatan:        mm.Catatan,
		Visible:        mm.Visible,
	}
	return e, nil
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error) {
	var list []model.LokalisBedahModel
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
	out := make([]*entity.LokalisBedah, 0, len(list))
	for i := range list {
		mm := list[i]
		out = append(out, &entity.LokalisBedah{
			ID:             mm.ID,
			IDPasien:       mm.IDPasien,
			IDDokter:       mm.IDDokter,
			Tanggal:        mm.Tanggal,
			LokasiKelainan: mm.LokasiKelainan,
			JenisKelainan:  mm.JenisKelainan,
			Ukuran:         mm.Ukuran,
			Warna:          mm.Warna,
			NyeriTekan:     mm.NyeriTekan,
			Konsistensi:    mm.Konsistensi,
			Mobilitas:      mm.Mobilitas,
			TandaRadang:    mm.TandaRadang,
			Fluktuasi:      mm.Fluktuasi,
			Catatan:        mm.Catatan,
			Visible:        mm.Visible,
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
			// accept RFC3339 string or time.Time
			switch val := v.(type) {
			case string:
				if t, err := time.Parse(time.RFC3339, val); err == nil {
					dbUpdates["tanggal"] = t
				} else {
					// fallback to raw value
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

	if err := m.DB.Model(&model.LokalisBedahModel{}).Where("id = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.LokalisBedahModel{}).
		Where("id = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
