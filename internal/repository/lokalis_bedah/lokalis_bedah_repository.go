package lokalis_bedah

import entity "rme/internal/entity/lokalis_bedah"

type Repository interface {
	Create(d *entity.LokalisBedah) error
	GetByID(id int) (*entity.LokalisBedah, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.LokalisBedah, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
}
