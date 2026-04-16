package poli

import (
	entity "rme/internal/entity/poli"
	poli_model "rme/internal/model/poli"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) GetAll(idDataKlinik *int) ([]*entity.Poli, error) {
	var models []poli_model.PoliModel
	q := r.db
	if idDataKlinik != nil {
		q = q.Where("id_data_klinik = ?", *idDataKlinik)
	}
	if err := q.Find(&models).Error; err != nil {
		return nil, err
	}
	var result []*entity.Poli
	for _, m := range models {
		result = append(result, &entity.Poli{
			ID:             m.ID,
			NamaPoli:       m.NamaPoli,
			Deskripsi:      m.Deskripsi,
			IDDataKlinik:   m.IDDataKlinik,
			Aktif:          m.Aktif,
			DibuatPada:     m.DibuatPada,
			DiperbaruiPada: m.DiperbaruiPada,
		})
	}
	return result, nil
}
