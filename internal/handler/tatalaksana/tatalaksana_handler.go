package tatalaksana

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/tatalaksana"
	usecase "rme/internal/usecase/tatalaksana"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Create menangani POST /tatalaksana
func (h *Handler) Create(c *gin.Context) {
	var req TatalaksanaCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const layout = "2006-01-02 15:04:05"
	var dateMake, dateUpdate time.Time
	var err error
	if req.DateMake != "" {
		dateMake, err = time.Parse(layout, req.DateMake)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateMake format"})
			return
		}
	} else {
		dateMake = time.Now()
	}
	if req.DateUpdate != "" {
		dateUpdate, err = time.Parse(layout, req.DateUpdate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid dateUpdate format"})
			return
		}
	} else {
		dateUpdate = dateMake
	}

	data := &entity.Tatalaksana{
		IDPasien:   req.IDPasien,
		IDDokter:   req.IDDokter,
		DateMake:   dateMake,
		DateUpdate: dateUpdate,
		NamaObat:   req.NamaObat,
		Dosis:      req.Dosis,
		Frekuensi:  req.Frekuensi,
		Durasi:     req.Durasi,
		CaraPakai:  req.CaraPakai,
		Catatan:    req.Catatan,
		Visible:    1,
	}

	if err := h.uc.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "tatalaksana created"})
}

// GetByID menangani GET /tatalaksana/:id
func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
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
	resp := TatalaksanaResponse{
		ID:         data.ID,
		IDPasien:   data.IDPasien,
		IDDokter:   data.IDDokter,
		DateMake:   data.DateMake,
		DateUpdate: data.DateUpdate,
		NamaObat:   data.NamaObat,
		Dosis:      data.Dosis,
		Frekuensi:  data.Frekuensi,
		Durasi:     data.Durasi,
		CaraPakai:  data.CaraPakai,
		Catatan:    data.Catatan,
		Visible:    data.Visible,
		NamaPasien: data.NamaPasien,
		NamaDokter: data.NamaDokter,
	}
	c.JSON(http.StatusOK, resp)
}

// GetAll menangani GET /tatalaksana
func (h *Handler) GetAll(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	idDokterStr := c.Query("idDokter")
	var idPasienPtr *int
	var idDokterPtr *int
	if idPasienStr != "" {
		v, err := strconv.Atoi(idPasienStr)
		if err == nil {
			idPasienPtr = &v
		}
	}
	if idDokterStr != "" {
		v, err := strconv.Atoi(idDokterStr)
		if err == nil {
			idDokterPtr = &v
		}
	}
	list, err := h.uc.GetAll(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]TatalaksanaResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, TatalaksanaResponse{
			ID:         d.ID,
			IDPasien:   d.IDPasien,
			IDDokter:   d.IDDokter,
			DateMake:   d.DateMake,
			DateUpdate: d.DateUpdate,
			NamaObat:   d.NamaObat,
			Dosis:      d.Dosis,
			Frekuensi:  d.Frekuensi,
			Durasi:     d.Durasi,
			CaraPakai:  d.CaraPakai,
			Catatan:    d.Catatan,
			Visible:    d.Visible,
			NamaPasien: d.NamaPasien,
			NamaDokter: d.NamaDokter,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// Update menangani PATCH /tatalaksana/:id
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

	// handle date_update if provided as string
	if v, ok := payload["date_update"].(string); ok && v != "" {
		const layout = "2006-01-02 15:04:05"
		if t, err := time.Parse(layout, v); err == nil {
			payload["date_update"] = t
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_update format"})
			return
		}
	} else {
		payload["date_update"] = time.Now()
	}

	// Only allow certain fields
	allowed := map[string]bool{
		"nama_obat":   true,
		"dosis":       true,
		"frekuensi":   true,
		"durasi":      true,
		"cara_pakai":  true,
		"catatan":     true,
		"date_update": true,
		"id_dokter":   true,
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

// Hide menangani PATCH /tatalaksana/:id/hide
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
