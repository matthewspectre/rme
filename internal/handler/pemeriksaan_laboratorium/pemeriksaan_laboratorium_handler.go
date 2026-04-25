package pemeriksaan_laboratorium

import (
	"fmt"
	"net/http"
	entity "rme/internal/entity/pemeriksaan_laboratorium"
	usecase "rme/internal/usecase/pemeriksaan_laboratorium"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ent := &entity.PemeriksaanLaboratorium{
		IDPasien:     req.IDPasien,
		IDDokter:     req.IDDokter,
		Hb:           req.Hb,
		Ht:           req.Ht,
		Leukosit:     req.Leukosit,
		Trombosit:    req.Trombosit,
		GulaPuasa:    req.GulaPuasa,
		GulaSewaktu:  req.GulaSewaktu,
		HbA1c:        req.HbA1c,
		Kolesterol:   req.Kolesterol,
		HDL:          req.HDL,
		LDL:          req.LDL,
		Trigliserida: req.Trigliserida,
		SGOT:         req.SGOT,
		SGPT:         req.SGPT,
		Ureum:        req.Ureum,
		Kreatinin:    req.Kreatinin,
		AsamUrat:     req.AsamUrat,
		Natrium:      req.Natrium,
		Kalium:       req.Kalium,
		Klorida:      req.Klorida,
		Visible:      1,
	}
	id, err := h.uc.Create(ent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "created", "id": id})
}

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
			ID:           r.ID,
			IDPasien:     r.IDPasien,
			Visible:      r.Visible,
			NamaPasien:   r.NamaPasien,
			IDDokter:     r.IDDokter,
			NamaDokter:   r.NamaDokter,
			Hb:           r.Hb,
			Ht:           r.Ht,
			Leukosit:     r.Leukosit,
			Trombosit:    r.Trombosit,
			GulaPuasa:    r.GulaPuasa,
			GulaSewaktu:  r.GulaSewaktu,
			HbA1c:        r.HbA1c,
			Kolesterol:   r.Kolesterol,
			HDL:          r.HDL,
			LDL:          r.LDL,
			Trigliserida: r.Trigliserida,
			SGOT:         r.SGOT,
			SGPT:         r.SGPT,
			Ureum:        r.Ureum,
			Kreatinin:    r.Kreatinin,
			AsamUrat:     r.AsamUrat,
			Natrium:      r.Natrium,
			Kalium:       r.Kalium,
			Klorida:      r.Klorida,
			Waktu:        r.Waktu,
		})
	}
	c.JSON(http.StatusOK, res)
}

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
	res := Response{
		ID:           row.ID,
		IDPasien:     row.IDPasien,
		Visible:      row.Visible,
		NamaPasien:   row.NamaPasien,
		IDDokter:     row.IDDokter,
		NamaDokter:   row.NamaDokter,
		Hb:           row.Hb,
		Ht:           row.Ht,
		Leukosit:     row.Leukosit,
		Trombosit:    row.Trombosit,
		GulaPuasa:    row.GulaPuasa,
		GulaSewaktu:  row.GulaSewaktu,
		HbA1c:        row.HbA1c,
		Kolesterol:   row.Kolesterol,
		HDL:          row.HDL,
		LDL:          row.LDL,
		Trigliserida: row.Trigliserida,
		SGOT:         row.SGOT,
		SGPT:         row.SGPT,
		Ureum:        row.Ureum,
		Kreatinin:    row.Kreatinin,
		AsamUrat:     row.AsamUrat,
		Natrium:      row.Natrium,
		Kalium:       row.Kalium,
		Klorida:      row.Klorida,
		Waktu:        row.Waktu,
	}
	c.JSON(http.StatusOK, res)
}

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
	if req.Hb != nil {
		updates["hb"] = req.Hb
	}
	if req.Ht != nil {
		updates["ht"] = req.Ht
	}
	if req.Leukosit != nil {
		updates["leukosit"] = req.Leukosit
	}
	if req.Trombosit != nil {
		updates["trombosit"] = req.Trombosit
	}
	if req.GulaPuasa != nil {
		updates["gula_puasa"] = req.GulaPuasa
	}
	if req.GulaSewaktu != nil {
		updates["gula_sewaktu"] = req.GulaSewaktu
	}
	if req.HbA1c != nil {
		updates["hba1c"] = req.HbA1c
	}
	if req.Kolesterol != nil {
		updates["kolesterol_total"] = req.Kolesterol
	}
	if req.HDL != nil {
		updates["hdl"] = req.HDL
	}
	if req.LDL != nil {
		updates["ldl"] = req.LDL
	}
	if req.Trigliserida != nil {
		updates["trigliserida"] = req.Trigliserida
	}
	if req.SGOT != nil {
		updates["sgot"] = req.SGOT
	}
	if req.SGPT != nil {
		updates["sgpt"] = req.SGPT
	}
	if req.Ureum != nil {
		updates["ureum"] = req.Ureum
	}
	if req.Kreatinin != nil {
		updates["kreatinin"] = req.Kreatinin
	}
	if req.AsamUrat != nil {
		updates["asam_urat"] = req.AsamUrat
	}
	if req.Natrium != nil {
		updates["natrium"] = req.Natrium
	}
	if req.Kalium != nil {
		updates["kalium"] = req.Kalium
	}
	if req.Klorida != nil {
		updates["klorida"] = req.Klorida
	}
	if req.Visible != nil {
		updates["visible"] = *req.Visible
	}

	updated, err := h.uc.Update(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

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
	updates := map[string]interface{}{"visible": 0}
	if _, err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "hidden"})
}
