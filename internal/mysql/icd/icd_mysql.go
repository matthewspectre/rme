package icd

import (
	repo "rme/internal/repository/icd"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RepositoryMySQL {
	return &RepositoryMySQL{db: db}
}

func (r *RepositoryMySQL) GetAll() ([]*repo.IcdRow, error) {
	var rows []repo.IcdRow
	q := r.db.Table("icd10").Select("id, kode, nama")
	if err := q.Scan(&rows).Error; err != nil {
		return nil, err
	}
	res := make([]*repo.IcdRow, 0, len(rows))
	for i := range rows {
		res = append(res, &rows[i])
	}
	return res, nil
}
