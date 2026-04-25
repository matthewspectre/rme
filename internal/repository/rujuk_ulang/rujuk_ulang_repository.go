package rujukulang

import (
	entity "rme/internal/entity/rujuk_ulang"
)

type Repository interface {
	Create(data *entity.RujukUlang) error
	GetByID(id int) (*entity.RujukUlang, error)
	GetAll(idPasien *int) ([]*entity.RujukUlang, error)
	Update(id int, updates map[string]interface{}) error
	Delete(id int) error
}
