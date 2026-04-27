package pemeriksaan_fungsi_organ

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/pemeriksaan_fungsi_organ"
	usecase "rme/internal/usecase/pemeriksaan_fungsi_organ"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req PemeriksaanFungsiOrganCreateRequest
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

	d := &entity.PemeriksaanFungsiOrgan{
		IDPasien:      req.IDPasien,
		IDDokter:      req.IDDokter,
		Tanggal:       tanggal,
		GangguanBAB:   req.GangguanBAB,
		GangguanBAK:   req.GangguanBAK,
		MualMuntah:    req.MualMuntah,
		Demam:         req.Demam,
		Perdarahan:    req.Perdarahan,
		PenurunanBB:   req.PenurunanBB,
		GangguanGerak: req.GangguanGerak,
		Catatan:       req.Catatan,
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

	resp := PemeriksaanFungsiOrganResponse{
		ID:            d.ID,
		IDPasien:      d.IDPasien,
		IDDokter:      d.IDDokter,
		Tanggal:       d.Tanggal.Format(time.RFC3339),
		GangguanBAB:   d.GangguanBAB,
		GangguanBAK:   d.GangguanBAK,
		MualMuntah:    d.MualMuntah,
		Demam:         d.Demam,
		Perdarahan:    d.Perdarahan,
		PenurunanBB:   d.PenurunanBB,
		GangguanGerak: d.GangguanGerak,
		Catatan:       d.Catatan,
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
	resp := make([]PemeriksaanFungsiOrganResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, PemeriksaanFungsiOrganResponse{
			ID:            d.ID,
			IDPasien:      d.IDPasien,
			IDDokter:      d.IDDokter,
			Tanggal:       d.Tanggal.Format(time.RFC3339),
			GangguanBAB:   d.GangguanBAB,
			GangguanBAK:   d.GangguanBAK,
			MualMuntah:    d.MualMuntah,
			Demam:         d.Demam,
			Perdarahan:    d.Perdarahan,
			PenurunanBB:   d.PenurunanBB,
			GangguanGerak: d.GangguanGerak,
			Catatan:       d.Catatan,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// Update handles PATCH /pemeriksaan_fungsi_organ/:id
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

	// handle tanggal if provided as string
	if v, ok := payload["tanggal"].(string); ok && v != "" {
		if t, err := time.Parse(time.RFC3339, v); err == nil {
			payload["tanggal"] = t
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format, use RFC3339"})
			return
		}
	}

	// ensure date_update
	if _, ok := payload["date_update"]; !ok {
		payload["date_update"] = time.Now()
	}

	// Allowed fields for update
	allowed := map[string]bool{
		"id_dokter":      true,
		"tanggal":        true,
		"gangguan_bab":   true,
		"gangguan_bak":   true,
		"mual_muntah":    true,
		"demam":          true,
		"perdarahan":     true,
		"penurunan_bb":   true,
		"gangguan_gerak": true,
		"catatan":        true,
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

// Hide handles PATCH /pemeriksaan_fungsi_organ/:id/hide (soft delete)
func (h *Handler) Hide(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	updates := map[string]interface{}{"visible": 0, "date_update": time.Now()}
	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}
