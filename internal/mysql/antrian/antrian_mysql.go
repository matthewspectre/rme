package antrian

import (
	"database/sql"
	"time"

	entity "rme/internal/entity/antrian"
	model "rme/internal/model/antrian"
	repo "rme/internal/repository/antrian"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) Create(a *entity.Antrian) (int, error) {
	return r.CreateWithNumber(a)
}

// CreateWithNumber inserts antrian inside a transaction and returns assigned nomorAntrian.
func (r *RepositoryMySQL) CreateWithNumber(a *entity.Antrian) (int, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return 0, tx.Error
	}
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	var max sql.NullInt64
	raw := "SELECT COALESCE(MAX(nomor_antrian),0) as max FROM antrian_pasien WHERE id_poli = ? AND DATE(waktu) = CURRENT_DATE() FOR UPDATE"
	if err := tx.Raw(raw, a.IDPoli).Row().Scan(&max); err != nil {
		tx.Rollback()
		return 0, err
	}
	next := int(max.Int64) + 1

	m := model.AntrianModel{
		IDPasien:     a.IDPasien,
		IDDokter:     a.IDDokter,
		NomorAntrian: next,
		IDPoli:       a.IDPoli,
		Waktu:        time.Now().Format("2006-01-02 15:04:05"),
		Status:       1,
	}
	if err := tx.Create(&m).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	if err := tx.Commit().Error; err != nil {
		return 0, err
	}
	return next, nil
}

func (r *RepositoryMySQL) GetAll(idDokter *int) ([]*repo.AntrianWithNamaPasien, error) {
	var rowsStruct []repo.AntrianWithNamaPasien
	// Join ke tabel patients untuk ambil nama pasien
	q := r.db.Table("antrian_pasien as a").
		Select("a.id, a.id_pasien, p.name as nama_pasien, a.id_dokter as id_dokter, d.nama_dokter as nama_dokter, a.nomor_antrian, a.id_poli, a.waktu, a.status as status").
		Joins("JOIN patients p ON a.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON a.id_dokter = d.id")
	if idDokter != nil {
		q = q.Where("a.id_dokter = ?", *idDokter)
	}
	err := q.Scan(&rowsStruct).Error
	if err != nil {
		return nil, err
	}
	// Convert to slice of pointers
	var rows []*repo.AntrianWithNamaPasien
	for i := range rowsStruct {
		rows = append(rows, &rowsStruct[i])
	}
	return rows, nil
}

func (r *RepositoryMySQL) Update(id int, updates map[string]interface{}) (*repo.AntrianWithNamaPasien, error) {
	mapped := make(map[string]interface{})
	for k, v := range updates {
		switch k {
		case "idPasien":
			mapped["id_pasien"] = v
		case "idDokter":
			mapped["id_dokter"] = v
		case "idPoli":
			mapped["id_poli"] = v
		case "nomorAntrian":
			mapped["nomor_antrian"] = v
		case "status":
			mapped["status"] = v
		default:
			mapped[k] = v
		}
	}

	if err := r.db.Model(&model.AntrianModel{}).Where("id = ?", id).Updates(mapped).Error; err != nil {
		return nil, err
	}

	// Fetch updated row with joins (patients, dokter)
	var row repo.AntrianWithNamaPasien
	err := r.db.Table("antrian_pasien as a").
		Select("a.id, a.id_pasien, p.name as nama_pasien, a.id_dokter as id_dokter, d.nama_dokter as nama_dokter, a.nomor_antrian, a.id_poli, a.waktu, a.status as status").
		Joins("JOIN patients p ON a.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON a.id_dokter = d.id").
		Where("a.id = ?", id).
		Scan(&row).Error
	if err != nil {
		return nil, err
	}
	return &row, nil
}

// UpdateByPatient updates rows where id_pasien = ? and returns updated rows
func (r *RepositoryMySQL) UpdateByPatient(idPasien int, updates map[string]interface{}) ([]*repo.AntrianWithNamaPasien, error) {
	// Map keys from DTO-style to DB columns if necessary
	mapped := make(map[string]interface{})
	for k, v := range updates {
		switch k {
		case "idPasien":
			mapped["id_pasien"] = v
		case "idDokter":
			mapped["id_dokter"] = v
		case "idPoli":
			mapped["id_poli"] = v
		case "nomorAntrian":
			mapped["nomor_antrian"] = v
		case "status":
			mapped["status"] = v
		default:
			mapped[k] = v
		}
	}

	if err := r.db.Model(&model.AntrianModel{}).Where("id_pasien = ?", idPasien).Updates(mapped).Error; err != nil {
		return nil, err
	}

	var rows []repo.AntrianWithNamaPasien
	q := r.db.Table("antrian_pasien as a").
		Select("a.id, a.id_pasien, p.name as nama_pasien, a.id_dokter as id_dokter, d.nama_dokter as nama_dokter, a.nomor_antrian, a.id_poli, a.waktu, a.status as status").
		Joins("JOIN patients p ON a.id_pasien = p.id").
		Joins("LEFT JOIN dokter d ON a.id_dokter = d.id").
		Where("a.id_pasien = ?", idPasien)

	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.AntrianWithNamaPasien, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}

func (r *RepositoryMySQL) DeleteByPatient(idPasien int) error {
	// Use Exec raw delete to avoid potential GORM layer issues returning EOF
	return r.db.Exec("DELETE FROM antrian_pasien WHERE id_pasien = ?", idPasien).Error
}
