package tatalaksana

import (
	"context"
	"time"

	entity "rme/internal/entity/tatalaksana"
	model "rme/internal/model/tatalaksana"
	repo "rme/internal/repository/tatalaksana"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.Tatalaksana) *model.TatalaksanaModel {
	if e == nil {
		return nil
	}
	return &model.TatalaksanaModel{
		ID:         e.ID,
		IDPasien:   e.IDPasien,
		IDDokter:   e.IDDokter,
		DateMake:   e.DateMake,
		DateUpdate: e.DateUpdate,
		NamaObat:   e.NamaObat,
		Dosis:      e.Dosis,
		Frekuensi:  e.Frekuensi,
		Durasi:     e.Durasi,
		CaraPakai:  e.CaraPakai,
		Catatan:    e.Catatan,
		Visible:    e.Visible,
	}
}

func toEntity(m *model.TatalaksanaModel) *entity.Tatalaksana {
	if m == nil {
		return nil
	}
	return &entity.Tatalaksana{
		ID:         m.ID,
		IDPasien:   m.IDPasien,
		IDDokter:   m.IDDokter,
		DateMake:   m.DateMake,
		DateUpdate: m.DateUpdate,
		NamaObat:   m.NamaObat,
		Dosis:      m.Dosis,
		Frekuensi:  m.Frekuensi,
		Durasi:     m.Durasi,
		CaraPakai:  m.CaraPakai,
		Catatan:    m.Catatan,
		Visible:    m.Visible,
	}
}

type joinedRow struct {
	ID         int       `gorm:"column:id_tatalaksana"`
	IDPasien   int       `gorm:"column:id_pasien"`
	IDDokter   int       `gorm:"column:id_dokter"`
	DateMake   time.Time `gorm:"column:date_make"`
	DateUpdate time.Time `gorm:"column:date_update"`
	NamaObat   string    `gorm:"column:nama_obat"`
	Dosis      string    `gorm:"column:dosis"`
	Frekuensi  string    `gorm:"column:frekuensi"`
	Durasi     string    `gorm:"column:durasi"`
	CaraPakai  string    `gorm:"column:cara_pakai"`
	Catatan    string    `gorm:"column:catatan"`
	Visible    int       `gorm:"column:visible"`
	NamaPasien string    `gorm:"column:nama_pasien"`
	NamaDokter string    `gorm:"column:nama_dokter"`
}

func rowToEntity(rj *joinedRow) *entity.Tatalaksana {
	if rj == nil {
		return nil
	}
	return &entity.Tatalaksana{
		ID:         rj.ID,
		IDPasien:   rj.IDPasien,
		IDDokter:   rj.IDDokter,
		DateMake:   rj.DateMake,
		DateUpdate: rj.DateUpdate,
		NamaObat:   rj.NamaObat,
		Dosis:      rj.Dosis,
		Frekuensi:  rj.Frekuensi,
		Durasi:     rj.Durasi,
		CaraPakai:  rj.CaraPakai,
		Catatan:    rj.Catatan,
		Visible:    rj.Visible,
		NamaPasien: rj.NamaPasien,
		NamaDokter: rj.NamaDokter,
	}
}

func (r *RepositoryMySQL) Create(data *entity.Tatalaksana) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

func (r *RepositoryMySQL) GetByID(id int) (*entity.Tatalaksana, error) {
	ctx := context.Background()
	var jr joinedRow
	query := `SELECT t.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
              FROM tatalaksana t
              LEFT JOIN patients p ON p.id = t.id_pasien
              LEFT JOIN dokter d ON d.id = t.id_dokter
              WHERE t.id_tatalaksana = ?`
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&jr).Error; err != nil {
		return nil, err
	}
	return rowToEntity(&jr), nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int, idDokter *int) ([]*entity.Tatalaksana, error) {
	ctx := context.Background()
	args := make([]interface{}, 0)
	query := `SELECT t.*, p.name AS nama_pasien, d.nama_dokter AS nama_dokter
              FROM tatalaksana t
              LEFT JOIN patients p ON p.id = t.id_pasien
              LEFT JOIN dokter d ON d.id = t.id_dokter
              WHERE t.visible = 1`
	if idPasien != nil {
		query += " AND t.id_pasien = ?"
		args = append(args, *idPasien)
	}
	if idDokter != nil {
		query += " AND t.id_dokter = ?"
		args = append(args, *idDokter)
	}
	query += " ORDER BY t.date_make DESC"

	var rows []joinedRow
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Tatalaksana, 0, len(rows))
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
	return r.db.WithContext(ctx).Model(&model.TatalaksanaModel{}).Where("id_tatalaksana = ?", id).Updates(updates).Error
}
