package patient

// MySQL implementation of patient repository.

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	entity "rme/internal/entity/patient"
	model "rme/internal/model/patient"
	repo "rme/internal/repository/patient"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.Patient) *model.PatientModel {
	if e == nil {
		return nil
	}
	return &model.PatientModel{
		ID:             e.ID,
		Name:           e.Name,
		AdmissionDate:  e.AdmissionDate,
		NoRekamMedis:   e.NoRekamMedis,
		NIK:            e.NIK,
		Gender:         e.Gender,
		BloodType:      e.BloodType,
		BirthPlaceDate: e.BirthPlaceDate,
		Phone:          e.Phone,
		Address:        e.Address,
		Category:       e.Category,
		Job:            e.Job,
		IDDataKlinik:   e.IDDataKlinik,
	}
}

func toEntity(m *model.PatientModel) *entity.Patient {
	if m == nil {
		return nil
	}
	return &entity.Patient{
		ID:             m.ID,
		Name:           m.Name,
		AdmissionDate:  m.AdmissionDate,
		NoRekamMedis:   m.NoRekamMedis,
		NIK:            m.NIK,
		Gender:         m.Gender,
		BloodType:      m.BloodType,
		BirthPlaceDate: m.BirthPlaceDate,
		Phone:          m.Phone,
		Address:        m.Address,
		Category:       m.Category,
		Job:            m.Job,
		IDDataKlinik:   m.IDDataKlinik,
	}
}

// Create menyimpan data pasien baru.
func (r *RepositoryMySQL) Create(data *entity.Patient) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

// CreateWithGeneratedRM inserts patient and generates `no_rekam_medis` if empty.
// Format: RM-<year>-<seq> where seq is zero-padded 3 digits and resets each year.
func (r *RepositoryMySQL) CreateWithGeneratedRM(data *entity.Patient) error {
	if data == nil {
		return fmt.Errorf("data is nil")
	}

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	// Only generate if not provided
	if data.NoRekamMedis == "" {
		year := time.Now().Year()
		var max sql.NullInt64
		raw := "SELECT COALESCE(MAX(CAST(RIGHT(no_rekam_medis,3) AS UNSIGNED)),0) AS maxseq FROM patients WHERE no_rekam_medis LIKE CONCAT('RM-', ?, '-%') FOR UPDATE"
		if err := tx.Raw(raw, year).Row().Scan(&max); err != nil {
			tx.Rollback()
			return err
		}
		next := int(max.Int64) + 1
		data.NoRekamMedis = fmt.Sprintf("RM-%d-%03d", year, next)
	}

	mdl := toModel(data)
	if err := tx.Create(mdl).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// Update memperbarui data pasien yang sudah ada berdasarkan primary key.
func (r *RepositoryMySQL) Update(data *entity.Patient) error {
	mdl := toModel(data)
	ctx := context.Background()
	return r.db.WithContext(ctx).Save(mdl).Error
}

// GetByNIK mengambil satu data pasien berdasarkan NIK.
func (r *RepositoryMySQL) GetByNIK(nik string) (*entity.Patient, error) {
	ctx := context.Background()
	var mdl model.PatientModel
	if err := r.db.WithContext(ctx).
		Where("nik = ?", nik).
		First(&mdl).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toEntity(&mdl), nil
}

// GetAll mengambil semua data pasien.
func (r *RepositoryMySQL) GetAll() ([]*entity.Patient, error) {
	ctx := context.Background()
	var mdls []model.PatientModel
	if err := r.db.WithContext(ctx).Find(&mdls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Patient, 0, len(mdls))
	for i := range mdls {
		res = append(res, toEntity(&mdls[i]))
	}
	return res, nil
}

// GetAllByIDDataKlinik mengambil data pasien berdasarkan id_data_klinik.
func (r *RepositoryMySQL) GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Patient, error) {
	ctx := context.Background()
	var mdls []model.PatientModel
	if err := r.db.WithContext(ctx).
		Where("id_data_klinik = ?", idDataKlinik).
		Find(&mdls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Patient, 0, len(mdls))
	for i := range mdls {
		res = append(res, toEntity(&mdls[i]))
	}
	return res, nil
}
