package pemeriksaan_ekg

import (
	"context"
	"time"

	entity "rme/internal/entity/pemeriksaan_ekg"
	model "rme/internal/model/pemeriksaan_ekg"
	repo "rme/internal/repository/pemeriksaan_ekg"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) Create(p *entity.PemeriksaanEkg) (int, error) {
	m := model.PemeriksaanEkgModel{
		IDPasien:           p.IDPasien,
		IDDokter:           p.IDDokter,
		DetakJantung:       p.DetakJantung,
		Irama:              p.Irama,
		PRInterval:         p.PRInterval,
		QRSDuration:        p.QRSDuration,
		QTTcInterval:       p.QTTcInterval,
		AxisJantung:        p.AxisJantung,
		STElevationDepress: p.STElevationDepress,
		TWaveAbnormality:   p.TWaveAbnormality,
		InterpretasiDokter: p.InterpretasiDokter,
		DateMake:           p.DateMake,
		DateUpdate:         p.DateUpdate,
		Visible:            1,
	}
	if m.DateMake.IsZero() {
		m.DateMake = time.Now()
	}
	if m.DateUpdate.IsZero() {
		m.DateUpdate = m.DateMake
	}
	if err := r.db.Create(&m).Error; err != nil {
		return 0, err
	}
	return m.ID, nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*repo.PemeriksaanEkgWithNames, error) {
	ctx := context.Background()
	var rows []repo.PemeriksaanEkgWithNames
	q := r.db.WithContext(ctx).Table("pemeriksaan_ekg as pe").
		Select("pe.id, pe.id_pasien, p.name as nama_pasien, pe.id_dokter, d.nama_dokter as nama_dokter, pe.detak_jantung, pe.irama, pe.pr_interval, pe.qrs_duration, pe.qt_qtc_interval AS qt_tc_interval, pe.axis_jantung, pe.st_elevation_depression AS st_elevation_depress, pe.t_wave_abnormality, pe.interpretasi_dokter, pe.visible, pe.date_make").
		Joins("JOIN patients p ON pe.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pe.id_dokter = d.id").
		Where("pe.visible = ?", 1)

	if idPasien != nil {
		q = q.Where("pe.id_pasien = ?", *idPasien)
	}
	if idDokter != nil {
		q = q.Where("pe.id_dokter = ?", *idDokter)
	}

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.PemeriksaanEkgWithNames, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func (r *RepositoryMySQL) GetByID(id int) (*repo.PemeriksaanEkgWithNames, error) {
	ctx := context.Background()
	var row repo.PemeriksaanEkgWithNames
	q := r.db.WithContext(ctx).Table("pemeriksaan_ekg as pe").
		Select("pe.id, pe.id_pasien, p.name as nama_pasien, pe.id_dokter, d.nama_dokter as nama_dokter, pe.detak_jantung, pe.irama, pe.pr_interval, pe.qrs_duration, pe.qt_qtc_interval AS qt_tc_interval, pe.axis_jantung, pe.st_elevation_depression AS st_elevation_depress, pe.t_wave_abnormality, pe.interpretasi_dokter, pe.visible, pe.date_make").
		Joins("JOIN patients p ON pe.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON pe.id_dokter = d.id").
		Where("pe.id = ?", id)
	if err := q.Scan(&row).Error; err != nil {
		return nil, err
	}
	return &row, nil
}

func (r *RepositoryMySQL) Update(id int, updates map[string]interface{}) (*repo.PemeriksaanEkgWithNames, error) {
	if updates == nil || len(updates) == 0 {
		return r.GetByID(id)
	}
	if err := r.db.Model(&model.PemeriksaanEkgModel{}).Where("id = ?", id).Updates(updates).Error; err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *RepositoryMySQL) Delete(id int) error {
	return r.db.Exec("DELETE FROM pemeriksaan_ekg WHERE id = ?", id).Error
}
