package rujukulang

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/rujuk_ulang"
	usecase "rme/internal/usecase/rujuk_ulang"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Create menangani POST /rujuk_ulang
func (h *Handler) Create(c *gin.Context) {
	var req RujukUlangCreateRequest
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

	data := &entity.RujukUlang{
		IDPasien:           req.IDPasien,
		IDDokter:           req.IDDokter,
		PoliAsal:           req.PoliAsal,
		PoliTujuan:         req.PoliTujuan,
		DiagnosisSementara: req.DiagnosisSementara,
		Catatan:            req.Catatan,
		DateMake:           dateMake,
		DateUpdate:         dateUpdate,
		Visible:            1,
	}

	if err := h.uc.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "rujuk ulang created"})
}

// GetByID menangani GET /rujuk_ulang/:id
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
	resp := RujukUlangResponse{
		ID:                 data.ID,
		IDPasien:           data.IDPasien,
		IDDokter:           data.IDDokter,
		PoliAsal:           data.PoliAsal,
		PoliTujuan:         data.PoliTujuan,
		DiagnosisSementara: data.DiagnosisSementara,
		Catatan:            data.Catatan,
		DateMake:           data.DateMake,
		DateUpdate:         data.DateUpdate,
		Visible:            data.Visible,
		NamaPasien:         data.NamaPasien,
		NamaPoliAsal:       data.NamaPoliAsal,
		NamaPoliTujuan:     data.NamaPoliTujuan,
		NamaDokter:         data.NamaDokter,
	}
	c.JSON(http.StatusOK, resp)
}

// GetAll menangani GET /rujuk_ulang
func (h *Handler) GetAll(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	var idPasienPtr *int
	if idPasienStr != "" {
		v, err := strconv.Atoi(idPasienStr)
		if err == nil {
			idPasienPtr = &v
		}
	}
	list, err := h.uc.GetAll(idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]RujukUlangResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, RujukUlangResponse{
			ID:                 d.ID,
			IDPasien:           d.IDPasien,
			IDDokter:           d.IDDokter,
			PoliAsal:           d.PoliAsal,
			PoliTujuan:         d.PoliTujuan,
			DiagnosisSementara: d.DiagnosisSementara,
			Catatan:            d.Catatan,
			DateMake:           d.DateMake,
			DateUpdate:         d.DateUpdate,
			Visible:            d.Visible,
			NamaPasien:         d.NamaPasien,
			NamaPoliAsal:       d.NamaPoliAsal,
			NamaPoliTujuan:     d.NamaPoliTujuan,
			NamaDokter:         d.NamaDokter,
		})
	}
	c.JSON(http.StatusOK, resp)
}

// Update menangani PATCH /rujuk_ulang/:id
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
		"diagnosis_sementara": true,
		"catatan":             true,
		"poli_asal":           true,
		"poli_tujuan":         true,
		"id_dokter":           true,
		"date_update":         true,
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

// Hide menangani PATCH /rujuk_ulang/:id/hide
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

// Delete menangani DELETE /rujuk_ulang/:id
func (h *Handler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
