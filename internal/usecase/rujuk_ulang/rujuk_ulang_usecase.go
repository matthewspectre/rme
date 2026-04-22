package rujukulang

import (
	entity "rme/internal/entity/rujuk_ulang"
	repo "rme/internal/repository/rujuk_ulang"
)

type Usecase interface {
	Create(data *entity.RujukUlang) error
	GetByID(id int) (*entity.RujukUlang, error)
	GetAll(idPasien *int) ([]*entity.RujukUlang, error)
	Update(id int, updates map[string]interface{}) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(data *entity.RujukUlang) error {
	return u.repo.Create(data)
}

func (u *usecase) GetByID(id int) (*entity.RujukUlang, error) {
	return u.repo.GetByID(id)
}

func (u *usecase) GetAll(idPasien *int) ([]*entity.RujukUlang, error) {
	return u.repo.GetAll(idPasien)
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}
