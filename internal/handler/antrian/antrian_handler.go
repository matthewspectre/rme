package antrian

import (
	"fmt"
	"net/http"
	entity "rme/internal/entity/antrian"
	usecase "rme/internal/usecase/antrian"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req AntrianCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	antrianBaru := &entity.Antrian{
		IDPasien: req.IDPasien,
		IDDokter: req.IDDokter,
		IDPoli:   req.IDPoli,
	}
	assigned, err := h.uc.Create(antrianBaru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Antrian berhasil ditambahkan", "nomorAntrian": assigned})
}

func (h *Handler) GetAll(c *gin.Context) {
	idDokterStr := c.Query("idDokter")
	var idDokterPtr *int
	if idDokterStr != "" {
		var v int
		if _, err := fmt.Sscan(idDokterStr, &v); err == nil {
			idDokterPtr = &v
		}
	}

	antrians, err := h.uc.GetAll(idDokterPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responses []AntrianResponse
	for _, a := range antrians {
		responses = append(responses, AntrianResponse{
			ID:           a.ID,
			IDPasien:     a.IDPasien,
			NamaPasien:   a.NamaPasien,
			IDDokter:     a.IDDokter,
			NamaDokter:   a.NamaDokter,
			Status:       a.Status,
			NomorAntrian: a.NomorAntrian,
			IDPoli:       a.IDPoli,
			Waktu:        a.Waktu,
		})
	}
	c.JSON(http.StatusOK, responses)
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

	var req AntrianUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.IDPasien != nil {
		updates["idPasien"] = *req.IDPasien
	}
	if req.IDDokter != nil {
		updates["idDokter"] = *req.IDDokter
	}
	if req.IDPoli != nil {
		updates["idPoli"] = *req.IDPoli
	}
	if req.NomorAntrian != nil {
		updates["nomorAntrian"] = *req.NomorAntrian
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	updated, err := h.uc.Update(id, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := AntrianResponse{
		ID:           updated.ID,
		IDPasien:     updated.IDPasien,
		NamaPasien:   updated.NamaPasien,
		IDDokter:     updated.IDDokter,
		NamaDokter:   updated.NamaDokter,
		Status:       updated.Status,
		NomorAntrian: updated.NomorAntrian,
		IDPoli:       updated.IDPoli,
		Waktu:        updated.Waktu,
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateByPatient menangani PATCH /antrian/patient/:id_pasien
func (h *Handler) UpdateByPatient(c *gin.Context) {
	idParam := c.Param("id_pasien")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_pasien is required"})
		return
	}
	var idPasien int
	if _, err := fmt.Sscan(idParam, &idPasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_pasien"})
		return
	}

	var req AntrianUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.IDPasien != nil {
		updates["idPasien"] = *req.IDPasien
	}
	if req.IDDokter != nil {
		updates["idDokter"] = *req.IDDokter
	}
	if req.IDPoli != nil {
		updates["idPoli"] = *req.IDPoli
	}
	if req.NomorAntrian != nil {
		updates["nomorAntrian"] = *req.NomorAntrian
	}
	if req.Status != nil {
		updates["status"] = *req.Status
	}

	updatedRows, err := h.uc.UpdateByPatient(idPasien, updates)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// build response array
	res := make([]AntrianResponse, 0, len(updatedRows))
	for _, a := range updatedRows {
		res = append(res, AntrianResponse{
			ID:           a.ID,
			IDPasien:     a.IDPasien,
			NamaPasien:   a.NamaPasien,
			IDDokter:     a.IDDokter,
			NamaDokter:   a.NamaDokter,
			Status:       a.Status,
			NomorAntrian: a.NomorAntrian,
			IDPoli:       a.IDPoli,
			Waktu:        a.Waktu,
		})
	}
	c.JSON(http.StatusOK, res)
}

// DeleteByPatient menangani DELETE /antrian/patient/:id_pasien
func (h *Handler) DeleteByPatient(c *gin.Context) {
	idParam := c.Param("id_pasien")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_pasien is required"})
		return
	}
	var idPasien int
	if _, err := fmt.Sscan(idParam, &idPasien); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_pasien"})
		return
	}

	if err := h.uc.DeleteByPatient(idPasien); err != nil {
		// Log detailed error to the server console for troubleshooting
		fmt.Printf("DeleteByPatient error id_pasien=%d: %v\n", idPasien, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "delete failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
