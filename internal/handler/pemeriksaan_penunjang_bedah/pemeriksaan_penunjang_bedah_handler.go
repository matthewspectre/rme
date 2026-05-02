package pemeriksaan_penunjang_bedah

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/pemeriksaan_penunjang_bedah"
	usecase "rme/internal/usecase/pemeriksaan_penunjang_bedah"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req PemeriksaanPenunjangBedahCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var jadwal time.Time
	if req.JadwalBedah != "" {
		t, err := time.Parse(time.RFC3339, req.JadwalBedah)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jadwalBedah format, use RFC3339"})
			return
		}
		jadwal = t
	}

	d := &entity.PemeriksaanPenunjangBedah{
		IDPasien:      req.IDPasien,
		IDDokter:      req.IDDokter,
		ButuhUSG:      req.ButuhUSG,
		ButuhRontgen:  req.ButuhRontgen,
		ButuhCTScan:   req.ButuhCTScan,
		ButuhBiopsi:   req.ButuhBiopsi,
		StatusOperasi: req.StatusOperasi,
		JenisTindakan: req.JenisTindakan,
		Prioritas:     req.Prioritas,
		CatatanBedah:  req.CatatanBedah,
		JadwalBedah:   jadwal,
		Visible:       1,
	}

	if err := h.uc.Create(d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": d.ID})
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
	resp := PemeriksaanPenunjangBedahResponse{
		ID:            d.ID,
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		ButuhUSG:      d.ButuhUSG,
		ButuhRontgen:  d.ButuhRontgen,
		ButuhCTScan:   d.ButuhCTScan,
		ButuhBiopsi:   d.ButuhBiopsi,
		StatusOperasi: d.StatusOperasi,
		JenisTindakan: d.JenisTindakan,
		Prioritas:     d.Prioritas,
		CatatanBedah:  d.CatatanBedah,
		JadwalBedah:   d.JadwalBedah.Format(time.RFC3339),
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAll(c *gin.Context) {
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
	resp := make([]PemeriksaanPenunjangBedahResponse, 0, len(list))
	for _, d := range list {
		jadwal := ""
		if !d.JadwalBedah.IsZero() {
			jadwal = d.JadwalBedah.Format(time.RFC3339)
		}
		resp = append(resp, PemeriksaanPenunjangBedahResponse{
			ID:            d.ID,
			IDPasien:      d.IDPasien,
			IDDokter:      d.IDDokter,
			ButuhUSG:      d.ButuhUSG,
			ButuhRontgen:  d.ButuhRontgen,
			ButuhCTScan:   d.ButuhCTScan,
			ButuhBiopsi:   d.ButuhBiopsi,
			StatusOperasi: d.StatusOperasi,
			JenisTindakan: d.JenisTindakan,
			Prioritas:     d.Prioritas,
			CatatanBedah:  d.CatatanBedah,
			JadwalBedah:   jadwal,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// Update supports partial updates
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// handle jadwalBedah if provided
	if v, ok := payload["jadwalBedah"].(string); ok && v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			payload["jadwal_bedah"] = t
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid jadwalBedah format, use RFC3339"})
			return
		}
		delete(payload, "jadwalBedah")
	}
	if _, ok := payload["date_update"]; !ok {
		payload["date_update"] = time.Now()
	}
	// whitelist fields
	allowed := map[string]bool{
		"butuh_usg":      true,
		"butuh_rontgen":  true,
		"butuh_ctscan":   true,
		"butuh_biopsi":   true,
		"status_operasi": true,
		"jenis_tindakan": true,
		"prioritas":      true,
		"catatan_bedah":  true,
		"jadwal_bedah":   true,
		"id_dokter":      true,
		"date_update":    true,
	}
	updates := make(map[string]interface{})
	for k, v := range payload {
		if allowed[k] {
			updates[k] = v
		}
	}
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no updatable fields provided"})
		return
	}
	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
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
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}

// HideByPatient handles PATCH /pemeriksaan_penunjang_bedah/patient/:id_pasien/hide
func (h *Handler) HideByPatient(c *gin.Context) {
	idParam := c.Param("id_pasien")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_pasien is required"})
		return
	}
	idPasien, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_pasien"})
		return
	}
	if err := h.uc.HideByPatient(idPasien); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}
