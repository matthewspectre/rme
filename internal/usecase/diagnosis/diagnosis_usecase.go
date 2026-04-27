package diagnosis

import (
	"time"

	entity "rme/internal/entity/diagnosis"
	repo "rme/internal/repository/diagnosis"
)

type Usecase interface {
	Create(d *entity.Diagnosis) error
	GetByID(id int) (*entity.Diagnosis, error)
	GetAll(idDokter *int, idPasien *int) ([]*entity.Diagnosis, error)
	Update(id int, updates map[string]interface{}) error
	Hide(id int) error
}

type usecase struct {
	repo repo.Repository
}

func NewUsecase(r repo.Repository) Usecase {
	return &usecase{repo: r}
}

func (u *usecase) Create(d *entity.Diagnosis) error {
	if d.Tanggal.IsZero() {
		d.Tanggal = time.Now()
	}
	return u.repo.Create(d)
}

func (u *usecase) GetByID(id int) (*entity.Diagnosis, error) {
	d, err := u.repo.GetByID(id)
	if err != nil || d == nil {
		return d, err
	}
	// resolve names
	codes := []string{}
	if d.KodeIcdUtama != "" {
		codes = append(codes, d.KodeIcdUtama)
	}
	for _, c := range d.KodeIcdSekunder {
		codes = append(codes, c)
	}
	names, err := u.repo.GetIcdNamesForCodes(codes)
	if err != nil {
		return nil, err
	}
	if d.KodeIcdUtama != "" {
		if name, ok := names[d.KodeIcdUtama]; ok {
			d.NamaIcdUtama = name
		}
	}
	// fill sekunder names in same order
	d.NamaIcdSekunder = make([]string, 0, len(d.KodeIcdSekunder))
	for _, c := range d.KodeIcdSekunder {
		if name, ok := names[c]; ok {
			d.NamaIcdSekunder = append(d.NamaIcdSekunder, name)
		} else {
			d.NamaIcdSekunder = append(d.NamaIcdSekunder, "")
		}
	}
	return d, nil
}

func (u *usecase) GetAll(idDokter *int, idPasien *int) ([]*entity.Diagnosis, error) {
	list, err := u.repo.GetAll(idDokter, idPasien)
	if err != nil {
		return nil, err
	}
	// collect codes
	var allCodes []string
	for _, d := range list {
		if d.KodeIcdUtama != "" {
			allCodes = append(allCodes, d.KodeIcdUtama)
		}
		for _, c := range d.KodeIcdSekunder {
			allCodes = append(allCodes, c)
		}
	}
	names, err := u.repo.GetIcdNamesForCodes(allCodes)
	if err != nil {
		return nil, err
	}
	// populate names
	for _, d := range list {
		if d.KodeIcdUtama != "" {
			if name, ok := names[d.KodeIcdUtama]; ok {
				d.NamaIcdUtama = name
			}
		}
		d.NamaIcdSekunder = make([]string, 0, len(d.KodeIcdSekunder))
		for _, c := range d.KodeIcdSekunder {
			if name, ok := names[c]; ok {
				d.NamaIcdSekunder = append(d.NamaIcdSekunder, name)
			} else {
				d.NamaIcdSekunder = append(d.NamaIcdSekunder, "")
			}
		}
	}
	return list, nil
}

func (u *usecase) Update(id int, updates map[string]interface{}) error {
	return u.repo.Update(id, updates)
}

func (u *usecase) Hide(id int) error {
	return u.repo.Hide(id)
}
