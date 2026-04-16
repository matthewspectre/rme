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
	antrians, err := h.uc.GetAll()
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
		NomorAntrian: updated.NomorAntrian,
		IDPoli:       updated.IDPoli,
		Waktu:        updated.Waktu,
	}
	c.JSON(http.StatusOK, resp)
}
