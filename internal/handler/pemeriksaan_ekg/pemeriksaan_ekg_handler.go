package pemeriksaan_ekg

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/pemeriksaan_ekg"
	usecase "rme/internal/usecase/pemeriksaan_ekg"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// Create POST /pemeriksaan_ekg
func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const layout = "2006-01-02 15:04:05"
	var dateMake time.Time
	var dateUpdate time.Time
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

	ent := &entity.PemeriksaanEkg{
		IDPasien:           req.IDPasien,
		IDDokter:           req.IDDokter,
		DetakJantung:       req.DetakJantung,
		Irama:              req.Irama,
		PRInterval:         req.PRInterval,
		QRSDuration:        req.QRSDuration,
		QTTcInterval:       req.QTTcInterval,
		AxisJantung:        req.AxisJantung,
		STElevationDepress: req.STElevationDepress,
		TWaveAbnormality:   req.TWaveAbnormality,
		InterpretasiDokter: req.InterpretasiDokter,
		DateMake:           dateMake,
		DateUpdate:         dateUpdate,
		Visible:            1,
	}

	id, err := h.uc.Create(ent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "created", "id": id})
}

// GetAll GET /pemeriksaan_ekg
func (h *Handler) GetAll(c *gin.Context) {
	idPasienStr := c.Query("idPasien")
	idDokterStr := c.Query("idDokter")
	var idPasienPtr *int
	var idDokterPtr *int
	if idPasienStr != "" {
		if v, err := strconv.Atoi(idPasienStr); err == nil {
			idPasienPtr = &v
		}
	}
	if idDokterStr != "" {
		if v, err := strconv.Atoi(idDokterStr); err == nil {
			idDokterPtr = &v
		}
	}

	rows, err := h.uc.GetAll(idPasienPtr, idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	res := make([]Response, 0, len(rows))
	for _, r := range rows {
		res = append(res, Response{
			ID:                 r.ID,
			IDPasien:           r.IDPasien,
			NamaPasien:         r.NamaPasien,
			IDDokter:           r.IDDokter,
			NamaDokter:         r.NamaDokter,
			Visible:            r.Visible,
			DetakJantung:       r.DetakJantung,
			Irama:              r.Irama,
			PRInterval:         r.PRInterval,
			QRSDuration:        r.QRSDuration,
			QTTcInterval:       r.QTTcInterval,
			AxisJantung:        r.AxisJantung,
			STElevationDepress: r.STElevationDepress,
			TWaveAbnormality:   r.TWaveAbnormality,
			InterpretasiDokter: r.InterpretasiDokter,
			DateMake:           r.DateMake,
		})
	}
	c.JSON(http.StatusOK, res)
}

// GetByID GET /pemeriksaan_ekg/:id
func (h *Handler) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var id int
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	row, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if row == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}
	res := Response{
		ID:                 row.ID,
		IDPasien:           row.IDPasien,
		NamaPasien:         row.NamaPasien,
		IDDokter:           row.IDDokter,
		NamaDokter:         row.NamaDokter,
		Visible:            row.Visible,
		DetakJantung:       row.DetakJantung,
		Irama:              row.Irama,
		PRInterval:         row.PRInterval,
		QRSDuration:        row.QRSDuration,
		QTTcInterval:       row.QTTcInterval,
		AxisJantung:        row.AxisJantung,
		STElevationDepress: row.STElevationDepress,
		TWaveAbnormality:   row.TWaveAbnormality,
		InterpretasiDokter: row.InterpretasiDokter,
		DateMake:           row.DateMake,
	}
	c.JSON(http.StatusOK, res)
}

// Update PATCH /pemeriksaan_ekg/:id
func (h *Handler) Update(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var id int
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var req UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	updates := make(map[string]interface{})
	if req.DetakJantung != nil {
		updates["detak_jantung"] = req.DetakJantung
	}
	if req.Irama != nil {
		updates["irama"] = *req.Irama
	}
	if req.PRInterval != nil {
		updates["pr_interval"] = req.PRInterval
	}
	if req.QRSDuration != nil {
		updates["qrs_duration"] = req.QRSDuration
	}
	if req.QTTcInterval != nil {
		updates["qt_qtc_interval"] = req.QTTcInterval
	}
	if req.AxisJantung != nil {
		updates["axis_jantung"] = *req.AxisJantung
	}
	if req.STElevationDepress != nil {
		updates["st_elevation_depression"] = *req.STElevationDepress
	}
	if req.TWaveAbnormality != nil {
		updates["t_wave_abnormality"] = *req.TWaveAbnormality
	}
	if req.InterpretasiDokter != nil {
		updates["interpretasi_dokter"] = *req.InterpretasiDokter
	}
	if req.Visible != nil {
		updates["visible"] = *req.Visible
	}
	// set date_update
	updates["date_update"] = time.Now()

	updated, err := h.uc.Update(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

// Hide PATCH /pemeriksaan_ekg/:id/hide
func (h *Handler) Hide(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var id int
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	updates := map[string]interface{}{"visible": 0, "date_update": time.Now()}
	if _, err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}

// Delete DELETE /pemeriksaan_ekg/:id
func (h *Handler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		return
	}
	var id int
	if _, err := fmt.Sscan(idParam, &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.uc.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
