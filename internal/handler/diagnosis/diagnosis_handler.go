package diagnosis

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/diagnosis"
	usecase "rme/internal/usecase/diagnosis"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req DiagnosisCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tanggal time.Time
	if req.Tanggal != "" {
		t, err := time.Parse(time.RFC3339, req.Tanggal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format, use RFC3339"})
			return
		}
		tanggal = t
	} else {
		tanggal = time.Now()
	}

	d := &entity.Diagnosis{
		IDPasien:         req.IDPasien,
		IDDokter:         req.IDDokter,
		Tanggal:          tanggal,
		KodeIcdUtama:     req.DiagnosisUtama.KodeIcd,
		DiagnosisBanding: req.DiagnosisBanding,
		Status:           req.Status,
		DasarDiagnosis:   req.DasarDiagnosis,
		Catatan:          req.Catatan,
		Visible:          1,
	}
	// sekunder codes
	if len(req.DiagnosisSekunder) > 0 {
		d.KodeIcdSekunder = make([]string, 0, len(req.DiagnosisSekunder))
		for _, s := range req.DiagnosisSekunder {
			d.KodeIcdSekunder = append(d.KodeIcdSekunder, s.KodeIcd)
		}
	}

	if err := h.uc.Create(d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id_diagnosis": d.IDDiagnosis})
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	d, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	resp := DiagnosisResponse{
		IDDiagnosis: d.IDDiagnosis,
		IDPasien:    d.IDPasien,
		IDDokter:    d.IDDokter,
		Tanggal:     d.Tanggal,
		DiagnosisUtama: DiagnosisCodeResponse{
			KodeIcd: d.KodeIcdUtama,
			Nama:    d.NamaIcdUtama,
		},
		DiagnosisBanding: d.DiagnosisBanding,
		Status:           d.Status,
		DasarDiagnosis:   d.DasarDiagnosis,
		Catatan:          d.Catatan,
	}
	// sekunder
	for i, code := range d.KodeIcdSekunder {
		name := ""
		if i < len(d.NamaIcdSekunder) {
			name = d.NamaIcdSekunder[i]
		}
		resp.DiagnosisSekunder = append(resp.DiagnosisSekunder, DiagnosisCodeResponse{KodeIcd: code, Nama: name})
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAll(c *gin.Context) {
	// support optional query params: idDokter and idPasien
	var idDokterPtr *int
	var idPasienPtr *int
	idDokterStr := c.Query("idDokter")
	idPasienStr := c.Query("idPasien")
	if idDokterStr != "" {
		idd, err := strconv.Atoi(idDokterStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
			return
		}
		idDokterPtr = &idd
	}
	if idPasienStr != "" {
		idp, err := strconv.Atoi(idPasienStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
			return
		}
		idPasienPtr = &idp
	}

	list, err := h.uc.GetAll(idDokterPtr, idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	respList := make([]DiagnosisResponse, 0, len(list))
	for _, d := range list {
		dr := DiagnosisResponse{
			IDDiagnosis:      d.IDDiagnosis,
			IDPasien:         d.IDPasien,
			IDDokter:         d.IDDokter,
			Tanggal:          d.Tanggal,
			DiagnosisUtama:   DiagnosisCodeResponse{KodeIcd: d.KodeIcdUtama, Nama: d.NamaIcdUtama},
			DiagnosisBanding: d.DiagnosisBanding,
			Status:           d.Status,
			DasarDiagnosis:   d.DasarDiagnosis,
			Catatan:          d.Catatan,
		}
		for i, code := range d.KodeIcdSekunder {
			name := ""
			if i < len(d.NamaIcdSekunder) {
				name = d.NamaIcdSekunder[i]
			}
			dr.DiagnosisSekunder = append(dr.DiagnosisSekunder, DiagnosisCodeResponse{KodeIcd: code, Nama: name})
		}
		respList = append(respList, dr)
	}
	c.JSON(http.StatusOK, respList)
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req DiagnosisUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.DiagnosisUtama != nil {
		updates["kode_icd_utama"] = req.DiagnosisUtama.KodeIcd
	}
	if req.DiagnosisSekunder != nil {
		sec := make([]string, 0, len(*req.DiagnosisSekunder))
		for _, s := range *req.DiagnosisSekunder {
			sec = append(sec, s.KodeIcd)
		}
		updates["kode_icd_sekunder"] = sec
	}
	if req.DiagnosisBanding != nil {
		updates["diagnosis_banding"] = *req.DiagnosisBanding
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}
	if req.DasarDiagnosis != nil {
		updates["dasar_diagnosis"] = *req.DasarDiagnosis
	}
	if req.Catatan != nil {
		updates["catatan"] = *req.Catatan
	}
	if req.Tanggal != nil {
		updates["tanggal"] = *req.Tanggal
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "diagnosis updated"})
}

func (h *Handler) Hide(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.uc.Hide(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "diagnosis hidden"})
}
