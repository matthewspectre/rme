package rujukulang

import (
	"context"
	"time"

	entity "rme/internal/entity/rujuk_ulang"
	model "rme/internal/model/rujuk_ulang"
	repo "rme/internal/repository/rujuk_ulang"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.RujukUlang) *model.RujukUlangModel {
	if e == nil {
		return nil
	}
	return &model.RujukUlangModel{
		ID:                 e.ID,
		IDPasien:           e.IDPasien,
		IDDokter:           e.IDDokter,
		PoliAsal:           e.PoliAsal,
		PoliTujuan:         e.PoliTujuan,
		DiagnosisSementara: e.DiagnosisSementara,
		Catatan:            e.Catatan,
		DateMake:           e.DateMake,
		DateUpdate:         e.DateUpdate,
		Visible:            e.Visible,
	}
}

func toEntity(m *model.RujukUlangModel) *entity.RujukUlang {
	if m == nil {
		return nil
	}
	return &entity.RujukUlang{
		ID:                 m.ID,
		IDPasien:           m.IDPasien,
		IDDokter:           m.IDDokter,
		PoliAsal:           m.PoliAsal,
		PoliTujuan:         m.PoliTujuan,
		DiagnosisSementara: m.DiagnosisSementara,
		Catatan:            m.Catatan,
		DateMake:           m.DateMake,
		DateUpdate:         m.DateUpdate,
		Visible:            m.Visible,
	}
}

type joinedRow struct {
	ID                 int       `gorm:"column:id_rujuk_ulang"`
	IDPasien           int       `gorm:"column:id_pasien"`
	PoliAsal           int       `gorm:"column:poli_asal"`
	PoliTujuan         int       `gorm:"column:poli_tujuan"`
	IDDokter           int       `gorm:"column:id_dokter"`
	DiagnosisSementara string    `gorm:"column:diagnosis_sementara"`
	Catatan            string    `gorm:"column:catatan"`
	DateMake           time.Time `gorm:"column:date_make"`
	DateUpdate         time.Time `gorm:"column:date_update"`
	Visible            int       `gorm:"column:visible"`
	NamaPasien         string    `gorm:"column:nama_pasien"`
	NamaPoliAsal       string    `gorm:"column:nama_poli_asal"`
	NamaPoliTujuan     string    `gorm:"column:nama_poli_tujuan"`
	NamaDokter         string    `gorm:"column:nama_dokter"`
}

func rowToEntity(rj *joinedRow) *entity.RujukUlang {
	if rj == nil {
		return nil
	}
	return &entity.RujukUlang{
		ID:                 rj.ID,
		IDPasien:           rj.IDPasien,
		PoliAsal:           rj.PoliAsal,
		PoliTujuan:         rj.PoliTujuan,
		IDDokter:           rj.IDDokter,
		DiagnosisSementara: rj.DiagnosisSementara,
		Catatan:            rj.Catatan,
		DateMake:           rj.DateMake,
		DateUpdate:         rj.DateUpdate,
		Visible:            rj.Visible,
		NamaPasien:         rj.NamaPasien,
		NamaPoliAsal:       rj.NamaPoliAsal,
		NamaPoliTujuan:     rj.NamaPoliTujuan,
		NamaDokter:         rj.NamaDokter,
	}
}

func (r *RepositoryMySQL) Create(data *entity.RujukUlang) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

func (r *RepositoryMySQL) GetByID(id int) (*entity.RujukUlang, error) {
	ctx := context.Background()
	var jr joinedRow
	query := `SELECT ru.*, p.name AS nama_pasien, pa.nama_poli AS nama_poli_asal, pt.nama_poli AS nama_poli_tujuan, d.nama_dokter AS nama_dokter
			  FROM rujuk_ulang ru
			  LEFT JOIN patients p ON p.id = ru.id_pasien
			  LEFT JOIN daftar_poli pa ON pa.id = ru.poli_asal
			  LEFT JOIN daftar_poli pt ON pt.id = ru.poli_tujuan
			  LEFT JOIN dokter d ON d.id = ru.id_dokter
			  WHERE ru.id_rujuk_ulang = ?`
	if err := r.db.WithContext(ctx).Raw(query, id).Scan(&jr).Error; err != nil {
		return nil, err
	}
	return rowToEntity(&jr), nil
}

func (r *RepositoryMySQL) GetAll(idPasien *int) ([]*entity.RujukUlang, error) {
	ctx := context.Background()
	args := make([]interface{}, 0)
	query := `SELECT ru.*, p.name AS nama_pasien, pa.nama_poli AS nama_poli_asal, pt.nama_poli AS nama_poli_tujuan, d.nama_dokter AS nama_dokter
			  FROM rujuk_ulang ru
			  LEFT JOIN patients p ON p.id = ru.id_pasien
			  LEFT JOIN daftar_poli pa ON pa.id = ru.poli_asal
			  LEFT JOIN daftar_poli pt ON pt.id = ru.poli_tujuan
			  LEFT JOIN dokter d ON d.id = ru.id_dokter
			  WHERE ru.visible = 1`
	if idPasien != nil {
		query += " AND ru.id_pasien = ?"
		args = append(args, *idPasien)
	}
	query += " ORDER BY ru.date_make DESC"

	var rows []joinedRow
	if err := r.db.WithContext(ctx).Raw(query, args...).Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.RujukUlang, 0, len(rows))
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
	return r.db.WithContext(ctx).Model(&model.RujukUlangModel{}).Where("id_rujuk_ulang = ?", id).Updates(updates).Error
}

func (r *RepositoryMySQL) Delete(id int) error {
	ctx := context.Background()
	return r.db.WithContext(ctx).Where("id_rujuk_ulang = ?", id).Delete(&model.RujukUlangModel{}).Error
}
