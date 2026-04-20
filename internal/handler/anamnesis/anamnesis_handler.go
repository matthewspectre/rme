package anamnesis

// HTTP handlers for anamnesis endpoints.

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	entity "rme/internal/entity/anamnesis"
	usecase "rme/internal/usecase/anamnesis"

	"github.com/gin-gonic/gin"
)

// Handler membungkus usecase dan menyediakan handler HTTP.
type Handler struct {
	uc usecase.Usecase
}

// NewHandler membuat instance Handler baru.
func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Create menangani POST /anamnesis (Gin handler)
func (h *Handler) Create(c *gin.Context) {
	var req AnamnesisCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// parse tanggal dari string (format: "2006-01-02 15:04:05")
	const layout = "2006-01-02 15:04:05"
	var dateMake, dateUpdate time.Time
	var err error

	if req.DateMake != "" {
		dateMake, err = time.Parse(layout, req.DateMake)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateMake format, use YYYY-MM-DD HH:MM:SS"})
			return
		}
	} else {
		dateMake = time.Now()
	}

	if req.DateUpdate != "" {
		dateUpdate, err = time.Parse(layout, req.DateUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateUpdate format, use YYYY-MM-DD HH:MM:SS"})
			return
		}
	} else {
		dateUpdate = dateMake
	}

	// mapping DTO → entity
	data := &entity.Anamnesis{
		IDPasien:              req.IDPasien,
		IDDokter:              req.IDDokter,
		Text:                  req.Text,
		DateMake:              dateMake,
		DateUpdate:            dateUpdate,
		IDDataKlinik:          req.IDDataKlinik,
		RiwayatPengobatan:     req.RiwayatPengobatan,
		RiwayatKeluarga:       req.RiwayatKeluarga,
		RiwayatPenyakitDahulu: req.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   req.RiwayatPenyakitLain,
		StatusKehamilan:       req.StatusKehamilan,
		KeluhanTambahan:       req.KeluhanTambahan,
		Visible:               1,
	}

	if err := h.uc.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "anamnesis created"})
}

// GetByID menangani GET /anamnesis/:id_pasien (Gin handler)
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id_pasien")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_pasien"})
		return
	}
	data, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	resp := AnamnesisResponse{
		ID:                    data.ID,
		IDPasien:              data.IDPasien,
		IDDokter:              data.IDDokter,
		NamaPasien:            data.NamaPasien,
		Text:                  data.Text,
		DateMake:              data.DateMake,
		DateUpdate:            data.DateUpdate,
		IDDataKlinik:          data.IDDataKlinik,
		RiwayatPengobatan:     data.RiwayatPengobatan,
		RiwayatKeluarga:       data.RiwayatKeluarga,
		RiwayatPenyakitDahulu: data.RiwayatPenyakitDahulu,
		RiwayatPenyakitLain:   data.RiwayatPenyakitLain,
		StatusKehamilan:       data.StatusKehamilan,
		KeluhanTambahan:       data.KeluhanTambahan,
		Visible:               data.Visible,
	}

	c.JSON(http.StatusOK, resp)
}

// GetAll menangani GET /anamnesis (Gin handler)
func (h *Handler) GetAll(c *gin.Context) {
	// optional filters: idDokter, idPasien
	idDokterStr := c.Query("idDokter")
	idPasienStr := c.Query("idPasien")
	var idDokterPtr *int
	var idPasienPtr *int
	if idDokterStr != "" {
		idDoc, convErr := strconv.Atoi(idDokterStr)
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
			return
		}
		idDokterPtr = &idDoc
	}
	if idPasienStr != "" {
		idP, convErr := strconv.Atoi(idPasienStr)
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
			return
		}
		idPasienPtr = &idP
	}

	list, err := h.uc.GetAll(idDokterPtr, idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]AnamnesisResponse, 0, len(list))
	for _, data := range list {
		resp = append(resp, AnamnesisResponse{
			ID:                    data.ID,
			NamaPasien:            data.NamaPasien,
			IDPasien:              data.IDPasien,
			IDDokter:              data.IDDokter,
			Text:                  data.Text,
			DateMake:              data.DateMake,
			DateUpdate:            data.DateUpdate,
			IDDataKlinik:          data.IDDataKlinik,
			RiwayatPengobatan:     data.RiwayatPengobatan,
			RiwayatKeluarga:       data.RiwayatKeluarga,
			RiwayatPenyakitDahulu: data.RiwayatPenyakitDahulu,
			RiwayatPenyakitLain:   data.RiwayatPenyakitLain,
			StatusKehamilan:       data.StatusKehamilan,
			KeluhanTambahan:       data.KeluhanTambahan,
			Visible:               data.Visible,
		})
	}

	c.JSON(http.StatusOK, resp)
}

// Update menangani PATCH /anamnesis/:id
func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req AnamnesisUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Text != nil {
		updates["text"] = *req.Text
	}
	if req.RiwayatPengobatan != nil {
		updates["riwayat_pengobatan"] = *req.RiwayatPengobatan
	}
	if req.RiwayatKeluarga != nil {
		updates["riwayat_keluarga"] = *req.RiwayatKeluarga
	}
	if req.RiwayatPenyakitDahulu != nil {
		updates["riwayat_penyakit_dahulu"] = *req.RiwayatPenyakitDahulu
	}
	if req.RiwayatPenyakitLain != nil {
		updates["riwayat_penyakit_lain"] = *req.RiwayatPenyakitLain
	}
	if req.StatusKehamilan != nil {
		updates["status_kehamilan"] = *req.StatusKehamilan
	}
	if req.KeluhanTambahan != nil {
		updates["keluhan_tambahan"] = *req.KeluhanTambahan
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "anamnesis updated"})
}

// Hide menangani PATCH /anamnesis/:id/hide — set `visible` = 0 (soft delete)
func (h *Handler) Hide(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	updates := map[string]interface{}{"visible": 0}
	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "anamnesis hidden"})
}
