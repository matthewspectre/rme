package diagnosis

import (
	"encoding/json"
	"errors"
	"fmt"

	entity "rme/internal/entity/diagnosis"
	model "rme/internal/model/diagnosis"
	icdModel "rme/internal/model/icd"

	"gorm.io/gorm"
)

type mysqlRepo struct {
	DB *gorm.DB
}

func NewMySQLRepo(db *gorm.DB) *mysqlRepo {
	return &mysqlRepo{DB: db}
}

func (m *mysqlRepo) Create(d *entity.Diagnosis) error {
	mm := model.DiagnosisModel{
		IDPasien:     d.IDPasien,
		IDDokter:     d.IDDokter,
		Tanggal:      d.Tanggal,
		KodeIcdUtama: d.KodeIcdUtama,
		Status:       d.Status,
		Catatan:      d.Catatan,
		Visible:      1,
		DateMake:     d.Tanggal,
		DateUpdate:   d.Tanggal,
	}

	if b, err := json.Marshal(d.KodeIcdSekunder); err == nil {
		mm.KodeIcdSekunder = string(b)
	}
	if b, err := json.Marshal(d.DiagnosisBanding); err == nil {
		mm.DiagnosisBanding = string(b)
	}
	if b, err := json.Marshal(d.DasarDiagnosis); err == nil {
		mm.DasarDiagnosis = string(b)
	}

	if err := m.DB.Create(&mm).Error; err != nil {
		return err
	}
	d.IDDiagnosis = mm.IDDiagnosis
	return nil
}

func (m *mysqlRepo) GetByID(id int) (*entity.Diagnosis, error) {
	var mm model.DiagnosisModel
	if err := m.DB.Where("id_diagnosis = ? AND visible = 1", id).First(&mm).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return m.mapModelToEntity(&mm)
}

func (m *mysqlRepo) GetAll(idDokter *int, idPasien *int) ([]*entity.Diagnosis, error) {
	var list []model.DiagnosisModel
	q := m.DB.Where("visible = 1")
	if idDokter != nil {
		q = q.Where("id_dokter = ?", *idDokter)
	}
	if idPasien != nil {
		q = q.Where("id_pasien = ?", *idPasien)
	}
	if err := q.Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*entity.Diagnosis, 0, len(list))
	for i := range list {
		e, err := m.mapModelToEntity(&list[i])
		if err != nil {
			return nil, err
		}
		out = append(out, e)
	}
	return out, nil
}

func (m *mysqlRepo) mapModelToEntity(mm *model.DiagnosisModel) (*entity.Diagnosis, error) {
	var sec []string
	var band []string
	var dasar []string
	if mm.KodeIcdSekunder != "" {
		_ = json.Unmarshal([]byte(mm.KodeIcdSekunder), &sec)
	}
	if mm.DiagnosisBanding != "" {
		_ = json.Unmarshal([]byte(mm.DiagnosisBanding), &band)
	}
	if mm.DasarDiagnosis != "" {
		_ = json.Unmarshal([]byte(mm.DasarDiagnosis), &dasar)
	}

	e := &entity.Diagnosis{
		IDDiagnosis:      mm.IDDiagnosis,
		IDPasien:         mm.IDPasien,
		IDDokter:         mm.IDDokter,
		Tanggal:          mm.Tanggal,
		KodeIcdUtama:     mm.KodeIcdUtama,
		KodeIcdSekunder:  sec,
		DiagnosisBanding: band,
		Status:           mm.Status,
		DasarDiagnosis:   dasar,
		Catatan:          mm.Catatan,
		Visible:          mm.Visible,
	}

	// Fetch nama for utama and sekunder by joining icd10 table
	// utama
	if e.KodeIcdUtama != "" {
		var icd icdModel.IcdModel
		if err := m.DB.Where("kode = ?", e.KodeIcdUtama).First(&icd).Error; err == nil {
			// attach name by replacing kode with "kode|nama" convention? We will keep kode fields and resolve names in usecase/handler
			// Alternatively, store mapping in entity by convention — leave names to higher layer.
			_ = icd
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	return e, nil
}

// helper to bulk fetch ICD names — used by usecase/handler
func (m *mysqlRepo) GetIcdNamesForCodes(codes []string) (map[string]string, error) {
	out := map[string]string{}
	if len(codes) == 0 {
		return out, nil
	}
	var icds []icdModel.IcdModel
	if err := m.DB.Where("kode IN (?)", codes).Find(&icds).Error; err != nil {
		return nil, err
	}
	for _, v := range icds {
		out[v.Kode] = v.Nama
	}
	return out, nil
}

// ensure mysqlRepo implements Repository
var _ = fmt.Sprintf

func (m *mysqlRepo) Update(id int, updates map[string]interface{}) error {
	dbUpdates := map[string]interface{}{}

	// map known update keys to columns; marshal slices to JSON strings
	for k, v := range updates {
		switch k {
		case "kode_icd_utama":
			if s, ok := v.(string); ok {
				dbUpdates["kode_icd_utama"] = s
			}
		case "kode_icd_sekunder":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["kode_icd_sekunder"] = string(b)
				}
			}
		case "diagnosis_banding":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["diagnosis_banding"] = string(b)
				}
			}
		case "dasar_diagnosis":
			if arr, ok := v.([]string); ok {
				if b, err := json.Marshal(arr); err == nil {
					dbUpdates["dasar_diagnosis"] = string(b)
				}
			}
		case "status":
			if s, ok := v.(string); ok {
				dbUpdates["status"] = s
			}
		case "catatan":
			if s, ok := v.(string); ok {
				dbUpdates["catatan"] = s
			}
		case "tanggal":
			if t, ok := v.(string); ok {
				dbUpdates["tanggal"] = t
			}
		}
	}

	// always update date_update
	dbUpdates["date_update"] = gorm.Expr("NOW()")

	if err := m.DB.Model(&model.DiagnosisModel{}).Where("id_diagnosis = ? AND visible = 1", id).Updates(dbUpdates).Error; err != nil {
		return err
	}
	return nil
}

func (m *mysqlRepo) Hide(id int) error {
	if err := m.DB.Model(&model.DiagnosisModel{}).
		Where("id_diagnosis = ? AND visible = 1", id).
		Updates(map[string]interface{}{"visible": 0, "date_update": gorm.Expr("NOW()")}).Error; err != nil {
		return err
	}
	return nil
}
