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
		ID:                    e.ID,
		IDPasien:              e.IDPasien,
		IDDokter:              e.IDDokter,
		Text:                  e.Text,
		DateMake:              e.DateMake,
		DateUpdate:            e.DateUpdate,
		IDDataKlinik:          e.IDDataKlinik,
		RiwayatPengobatan:     e.RiwayatPengobatan,
		RiwayatKeluarga:       e.RiwayatKeluarga,
		RiwayatPenyakitDahulu: e.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   e.RiwayatPenyakitLain,
		StatusKehamilan:       e.StatusKehamilan,
		KeluhanTambahan:       e.KeluhanTambahan,
		Visible:               e.Visible,
	}
}

func toEntity(m *model.AnamnesisModel) *entity.Anamnesis {
	if m == nil {
		return nil
	}
	return &entity.Anamnesis{
		ID:                    m.ID,
		IDPasien:              m.IDPasien,
		IDDokter:              m.IDDokter,
		Text:                  m.Text,
		DateMake:              m.DateMake,
		DateUpdate:            m.DateUpdate,
		IDDataKlinik:          m.IDDataKlinik,
		RiwayatPengobatan:     m.RiwayatPengobatan,
		RiwayatKeluarga:       m.RiwayatKeluarga,
		RiwayatPenyakitDahulu: m.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   m.RiwayatPenyakitLain,
		StatusKehamilan:       m.StatusKehamilan,
		KeluhanTambahan:       m.KeluhanTambahan,
		Visible:               m.Visible,
	}
}

// Create menyimpan data anamnesis baru.
func (r *RepositoryMySQL) Create(data *entity.Anamnesis) error {
	mdl := toModel(data)
	return r.db.Create(mdl).Error
}

// GetByID mengambil satu data anamnesis (ambil baris pertama).
// Parameter `idPasien` diterima tetapi tidak dipakai di query.
func (r *RepositoryMySQL) GetByID(idPasien int) (*entity.Anamnesis, error) {
	ctx := context.Background()
	type modelWithName struct {
		model.AnamnesisModel
		NamaPasien string `gorm:"column:nama_pasien"`
	}
	var aw modelWithName
	if err := r.db.WithContext(ctx).
		Table("anamnesis a").
		Select("a.*, p.name AS nama_pasien").
		Joins("LEFT JOIN patients p ON p.id = a.id_pasien").
		First(&aw).Error; err != nil {
		return nil, err
	}
	ent := toEntity(&aw.AnamnesisModel)
	ent.NamaPasien = aw.NamaPasien
	return ent, nil
}

// GetAll mengambil semua data anamnesis yang masih visible.
// Jika `idDokter` atau `idPasien` tidak nil, hasil akan difilter berdasarkan kolom terkait.
func (r *RepositoryMySQL) GetAll(idDokter *int, idPasien *int) ([]*entity.Anamnesis, error) {
	ctx := context.Background()
	type modelWithName struct {
		model.AnamnesisModel
		NamaPasien string `gorm:"column:nama_pasien"`
	}
	var awls []modelWithName
	q := r.db.WithContext(ctx).Table("anamnesis a").Select("a.*, p.name AS nama_pasien").Joins("LEFT JOIN patients p ON p.id = a.id_pasien").Where("a.visible = ?", 1)
	if idDokter != nil {
		q = q.Where("a.id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("a.id_pasien = ?", *idPasien)
	}
	if err := q.Find(&awls).Error; err != nil {
		return nil, err
	}
	res := make([]*entity.Anamnesis, 0, len(awls))
	for i := range awls {
		ent := toEntity(&awls[i].AnamnesisModel)
		ent.NamaPasien = awls[i].NamaPasien
		res = append(res, ent)
	}
	return res, nil
}

// Update memperbarui kolom pada baris anamnesis yang ditentukan oleh id_anamnesis.
func (r *RepositoryMySQL) Update(idAnamnesis int, updates map[string]interface{}) error {
	ctx := context.Background()
	if updates == nil || len(updates) == 0 {
		return nil
	}
	// Pastikan kolom yang diupdate sesuai nama kolom di DB (snake_case)
	return r.db.WithContext(ctx).
		Model(&model.AnamnesisModel{}).
		Where("id_anamnesis = ?", idAnamnesis).
		Updates(updates).Error
}
