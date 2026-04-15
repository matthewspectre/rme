package user

// MySQL implementation of user repository.

import (
	"context"
	"errors"

	entity "rme/internal/entity/user"
	model "rme/internal/model/user"
	repo "rme/internal/repository/user"

	"gorm.io/gorm"
)

type RepositoryMySQL struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repo.Repository {
	return &RepositoryMySQL{db: db}
}

func toEntity(m *model.UserModel) *entity.User {
	if m == nil {
		return nil
	}
	return &entity.User{
		ID:       m.ID,
		Username: m.Username,
		Password: m.Password,
		FullName: m.FullName,
		Role:     m.Role,
	}
}

// FindByUsername mencari user berdasarkan username.
func (r *RepositoryMySQL) FindByUsername(username string) (*entity.User, error) {
	ctx := context.Background()
	var m model.UserModel
	if err := r.db.WithContext(ctx).
		Where("username = ?", username).
		First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return toEntity(&m), nil
}
