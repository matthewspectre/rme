package pemeriksaanvital

import (
	"context"
	"time"

	entity "rme/internal/entity/pemeriksaan_vital"
	model "rme/internal/model/pemeriksaan_vital"
	repo "rme/internal/repository/pemeriksaan_vital"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.PemeriksaanVital) *model.PemeriksaanVitalModel {
	if e == nil {
		return nil
	}
	return &model.PemeriksaanVitalModel{
		ID:             e.ID,
		IDPasien:       e.IDPasien,
		IDDokter:       e.IDDokter,
		DateMake:       e.DateMake,
		DateUpdate:     e.DateUpdate,
		TekananDarah:   e.TekananDarah,
		DenyutNadi:     e.DenyutNadi,
		SuhuTubuh:      e.SuhuTubuh,
		FrekuensiNapas: e.FrekuensiNapas,
		BeratBadan:     e.BeratBadan,
		TinggiBadan:    e.TinggiBadan,
		Visible:        e.Visible,
	}
}

func toEntity(m *model.PemeriksaanVitalModel) *entity.PemeriksaanVital {
	if m == nil {
		return nil
	}
	return &entity.PemeriksaanVital{
		ID:             m.ID,
		IDPasien:       m.IDPasien,
		IDDokter:       m.IDDokter,
		DateMake:       m.DateMake,
		DateUpdate:     m.DateUpdate,
		TekananDarah:   m.TekananDarah,
		DenyutNadi:     m.DenyutNadi,
		SuhuTubuh:      m.SuhuTubuh,
		FrekuensiNapas: m.FrekuensiNapas,
		BeratBadan:     m.BeratBadan,
		TinggiBadan:    m.TinggiBadan,
		Visible:        m.Visible,
	}
}

type joinedRow struct {
	ID             int       `gorm:"column:id_pemeriksaan_vital"`
	IDPasien       int       `gorm:"column:id_pasien"`
	IDDokter       int       `gorm:"column:id_dokter"`
	DateMake       time.Time `gorm:"column:date_make"`
	DateUpdate     time.Time `gorm:"column:date_update"`
	TekananDarah   string    `gorm:"column:tekanan_darah"`
	DenyutNadi     int       `gorm:"column:denyut_nadi"`
	SuhuTubuh      float64   `gorm:"column:suhu_tubuh"`
	FrekuensiNapas int       `gorm:"column:frekuensi_napas"`
	BeratBadan     float64   `gorm:"column:berat_badan"`
	TinggiBadan    float64   `gorm:"column:tinggi_badan"`
	Visible        int       `gorm:"column:visible"`
	NamaPasien     string    `gorm:"column:nama_pasien"`
	NamaDokter     string    `gorm:"column:nama_dokter"`
}

func rowToEntity(rj *joinedRow) *entity.PemeriksaanVital {
	if rj == nil {
		return nil
	}
	return &entity.PemeriksaanVital{
		ID:             rj.ID,
		IDPasien:       rj.IDPasien,
		IDDokter:       rj.IDDokter,
		DateMake:       rj.DateMake,
		DateUpdate:     rj.DateUpdate,
		TekananDarah:   rj.TekananDarah,
		DenyutNadi:     rj.DenyutNadi,
		SuhuTubuh:      rj.SuhuTubuh,
		FrekuensiNapas: rj.FrekuensiNapas,
		BeratBadan:     rj.BeratBadan,
		TinggiBadan:    rj.TinggiBadan,
		Visible:        rj.Visible,
		NamaPasien:     rj.NamaPasien,
		NamaDokter:     rj.NamaDokter,
	}
}

func (r *RepositoryMySQL) Create(data *entity.PemeriksaanVital) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

func (r *RepositoryMySQL) GetByID(id int) (*entity.PemeriksaanVital, error) {
	ctx := context.Background()
	var jr joinedRow
	query := `SELECT pv.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
			  FROM pemeriksaan_vital pv
			  LEFT JOIN patients p ON p.id = pv.id_pasien
			  LEFT JOIN dokter d ON d.id = pv.id_dokter
			  WHERE pv.id_pemeriksaan_vital = ?`
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&jr).Error; err != nil {
		return nil, err
	}
	return rowToEntity(&jr), nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*entity.PemeriksaanVital, error) {
	ctx := context.Background()
	args := make([]interface{}, 0)
	query := `SELECT pv.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
			  FROM pemeriksaan_vital pv
			  LEFT JOIN patients p ON p.id = pv.id_pasien
			  LEFT JOIN dokter d ON d.id = pv.id_dokter
			  WHERE pv.visible = 1`
	if idPasien != nil {
		query += " AND pv.id_pasien = ?"
		args = append(args, *idPasien)
	}
	if idDokter != nil {
		query += " AND pv.id_dokter = ?"
		args = append(args, *idDokter)
	}
	query += " ORDER BY pv.date_make DESC"

	var rows []joinedRow
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.PemeriksaanVital, 0, len(rows))
	for i := range rows {
		res = append(res, rowToEntity(&rows[i]))
	}
	return res, nil
}

func (r *RepositoryMySQL) Update(id int, updates map[string]interface{}) error {
	ctx := context.Background()
	if updates == nil || len(updates) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).Model(&model.PemeriksaanVitalModel{}).Where("id_pemeriksaan_vital = ?", id).Updates(updates).Error
}
