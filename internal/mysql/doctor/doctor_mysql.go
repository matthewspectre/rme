package doctor

// MySQL implementation of doctor repository.

import (
	"context"

	entity "rme/internal/entity/doctor"
	model "rme/internal/model/doctor"
	repo "rme/internal/repository/doctor"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.Doctor) *model.DoctorModel {
	if e == nil {
		return nil
	}
	return &model.DoctorModel{
		ID:           e.ID,
		IDUser:       e.IDUser,
		NamaDokter:   e.Nama,
		Poli:         e.Poli,
		NomorTelepon: e.NomorTelepon,
		IDDataKlinik: e.IDDataKlinik,
	}
}

func toEntity(m *model.DoctorModel) *entity.Doctor {
	if m == nil {
		return nil
	}
	return &entity.Doctor{
		ID:           m.ID,
		IDUser:       m.IDUser,
		Nama:         m.NamaDokter,
		Poli:         m.Poli,
		NomorTelepon: m.NomorTelepon,
		IDDataKlinik: m.IDDataKlinik,
	}
}

func (r *RepositoryMySQL) Create(data *entity.Doctor) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

func (r *RepositoryMySQL) GetAll() ([]*entity.Doctor, error) {
	ctx := context.Background()
	var mdls []model.DoctorModel
	if err := r.db.WithContext(ctx).Find(&mdls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Doctor, 0, len(mdls))
	for i := range mdls {
		res = append(res, toEntity(&mdls[i]))
	}
	return res, nil
}

func (r *RepositoryMySQL) GetAllByIDDataKlinik(idDataKlinik int) ([]*entity.Doctor, error) {
	ctx := context.Background()
	var mdls []model.DoctorModel
	if err := r.db.WithContext(ctx).
		Where("id_data_klinik = ?", idDataKlinik).
		Find(&mdls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Doctor, 0, len(mdls))
	for i := range mdls {
		res = append(res, toEntity(&mdls[i]))
	}
	return res, nil
}
