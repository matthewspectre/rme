package anamnesis

// MySQL implementation of anamnesis repository.

import (
	"context"

	entity "rme/internal/entity/anamnesis"
	model "rme/internal/model/anamnesis"
	repo "rme/internal/repository/anamnesis"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toModel(e *entity.Anamnesis) *model.AnamnesisModel {
	if e == nil {
		return nil
	}
	return &model.AnamnesisModel{
		IDPasien:                e.IDPasien,
		IDDokter:                e.IDDokter,
		Text:                    e.Text,
		DateMake:                e.DateMake,
		DateUpdate:              e.DateUpdate,
		IDDataKlinik:            e.IDDataKlinik,
		RiwayatPengobatan:       e.RiwayatPengobatan,
		RiwayatKeluarga:         e.RiwayatKeluarga,
		RiwayatPekerjaan:        e.RiwayatPekerjaan,
		RiwayatAutoanamnesis:    e.RiwayatAutoanamnesis,
		RiwayatPenyakitDahulu:   e.RiwayatPenyakitDahulu,
		RiwayatPenyakitSekarang: e.RiwayatPenyakitSekarang,
		RiwayatPenyakitLain:     e.RiwayatPenyakitLain,
		HubunganPasien:          e.HubunganPasien,
		RiwayatAnestesiBedah:    e.RiwayatAnestesiBedah,
		RiwayatKeluhanUtama:     e.RiwayatKeluhanUtama,
		StatusKehamilan:         e.StatusKehamilan,
		KeluhanTambahan:         e.KeluhanTambahan,
		Catatan:                 e.Catatan,
		Visible:                 e.Visible,
	}
}

func toEntity(m *model.AnamnesisModel) *entity.Anamnesis {
	if m == nil {
		return nil
	}
	return &entity.Anamnesis{
		IDPasien:                m.IDPasien,
		IDDokter:                m.IDDokter,
		Text:                    m.Text,
		DateMake:                m.DateMake,
		DateUpdate:              m.DateUpdate,
		IDDataKlinik:            m.IDDataKlinik,
		RiwayatPengobatan:       m.RiwayatPengobatan,
		RiwayatKeluarga:         m.RiwayatKeluarga,
		RiwayatPekerjaan:        m.RiwayatPekerjaan,
		RiwayatAutoanamnesis:    m.RiwayatAutoanamnesis,
		RiwayatPenyakitDahulu:   m.RiwayatPenyakitDahulu,
		RiwayatPenyakitSekarang: m.RiwayatPenyakitSekarang,
		RiwayatPenyakitLain:     m.RiwayatPenyakitLain,
		HubunganPasien:          m.HubunganPasien,
		RiwayatAnestesiBedah:    m.RiwayatAnestesiBedah,
		RiwayatKeluhanUtama:     m.RiwayatKeluhanUtama,
		StatusKehamilan:         m.StatusKehamilan,
		KeluhanTambahan:         m.KeluhanTambahan,
		Catatan:                 m.Catatan,
		Visible:                 m.Visible,
	}
}

// Create menyimpan data anamnesis baru.
func (r *RepositoryMySQL) Create(data *entity.Anamnesis) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

// GetByID mengambil satu data anamnesis berdasarkan ID pasien dan idPoli.
func (r *RepositoryMySQL) GetByID(idPasien int, idPoli int) (*entity.Anamnesis, error) {
	ctx := context.Background()
	var mdl model.AnamnesisModel
	if err := r.db.WithContext(ctx).
		Where("id_pasien = ? AND id_data_klinik = ? AND visible = ?", idPasien, idPoli, 1).
		First(&mdl).Error; err != nil {
		return nil, err
	}
	return toEntity(&mdl), nil
}

// GetAll mengambil semua data anamnesis yang masih visible dan sesuai poli.
func (r *RepositoryMySQL) GetAll(idPoli int) ([]*entity.Anamnesis, error) {
	ctx := context.Background()
	var mdls []model.AnamnesisModel
	if err := r.db.WithContext(ctx).
		Where("id_data_klinik = ? AND visible = ?", idPoli, 1).
		Find(&mdls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Anamnesis, 0, len(mdls))
	for i := range mdls {
		res = append(res, toEntity(&mdls[i]))
	}
	return res, nil
}
