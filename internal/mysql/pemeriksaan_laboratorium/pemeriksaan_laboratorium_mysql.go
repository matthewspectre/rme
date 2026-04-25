package pemeriksaan_laboratorium

import (
	"time"

	entity "rme/internal/entity/pemeriksaan_laboratorium"
	model "rme/internal/model/pemeriksaan_laboratorium"
	repo "rme/internal/repository/pemeriksaan_laboratorium"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) Create(p *entity.PemeriksaanLaboratorium) (int, error) {
	m := model.PemeriksaanLaboratoriumModel{
		IDPasien:     p.IDPasien,
		IDDokter:     p.IDDokter,
		Hb:           p.Hb,
		Ht:           p.Ht,
		Leukosit:     p.Leukosit,
		Trombosit:    p.Trombosit,
		GulaPuasa:    p.GulaPuasa,
		GulaSewaktu:  p.GulaSewaktu,
		HbA1c:        p.HbA1c,
		Kolesterol:   p.Kolesterol,
		HDL:          p.HDL,
		LDL:          p.LDL,
		Trigliserida: p.Trigliserida,
		SGOT:         p.SGOT,
		SGPT:         p.SGPT,
		Ureum:        p.Ureum,
		Kreatinin:    p.Kreatinin,
		AsamUrat:     p.AsamUrat,
		Natrium:      p.Natrium,
		Kalium:       p.Kalium,
		Klorida:      p.Klorida,
		Waktu:        time.Now().Format("2006-01-02 15:04:05"),
		Visible:      1,
	}
	if err := r.db.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanLaboratoriumWithNames, error) {
	var rows []repo.PemeriksaanLaboratoriumWithNames
	q := r.db.Table("pemeriksaan_laboratorium as pl").
		Select("pl.id, pl.id_pasien, p.name as nama_pasien, pl.id_dokter, d.nama_dokter as nama_dokter, pl.hb, pl.ht, pl.leukosit, pl.trombosit, pl.gula_puasa, pl.gula_sewaktu, pl.hba1c AS hb_a1c, pl.kolesterol_total AS kolesterol, pl.hdl, pl.ldl, pl.trigliserida, pl.sgot, pl.sgpt, pl.ureum, pl.kreatinin, pl.asam_urat, pl.natrium, pl.kalium, pl.klorida, pl.visible, pl.waktu").
		Joins("JOIN patients p ON pl.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pl.id_dokter = d.id").
		Where("pl.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pl.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pl.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanLaboratoriumWithNames, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func (r *RepositoryMySQL) GetByID(id int) (*repo.PemeriksaanLaboratoriumWithNames, error) {
	var row repo.PemeriksaanLaboratoriumWithNames
	q := r.db.Table("pemeriksaan_laboratorium as pl").
		Select("pl.id, pl.id_pasien, p.name as nama_pasien, pl.id_dokter, d.nama_dokter as nama_dokter, pl.hb, pl.ht, pl.leukosit, pl.trombosit, pl.gula_puasa, pl.gula_sewaktu, pl.hba1c AS hb_a1c, pl.kolesterol_total AS kolesterol, pl.hdl, pl.ldl, pl.trigliserida, pl.sgot, pl.sgpt, pl.ureum, pl.kreatinin, pl.asam_urat, pl.natrium, pl.kalium, pl.klorida, pl.visible, pl.waktu").
		Joins("JOIN patients p ON pl.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pl.id_dokter = d.id").
		Where("pl.id = ?", id)
	if err := q.Scan(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *RepositoryMySQL) Update(id int, updates map[string]interface{}) (*repo.PemeriksaanLaboratoriumWithNames, error) {
	mapped := make(map[string]interface{})
	for k, v := range updates {
		switch k {
		case "hb":
			mapped["hb"] = v
		case "ht":
			mapped["ht"] = v
		case "visible":
			mapped["visible"] = v
		default:
			mapped[k] = v
		}
	}
	if err := r.db.Model(&model.PemeriksaanLaboratoriumModel{}).Where("id = ?", id).Updates(mapped).Error; err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *RepositoryMySQL) Delete(id int) error {
	return r.db.Exec("DELETE FROM pemeriksaan_laboratorium WHERE id = ?", id).Error
}
