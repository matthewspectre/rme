package poli

import (
	"net/http"
	"rme/internal/usecase/poli"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc poli.Usecase
}

func NewHandler(uc poli.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) GetAll(c *gin.Context) {
	idDataKlinikStr := c.Query("id_data_klinik")
	var idDataKlinik *int
	if idDataKlinikStr != "" {
		if val, err := strconv.Atoi(idDataKlinikStr); err == nil {
			idDataKlinik = &val
		}
	}
	polis, err := h.uc.GetAll(idDataKlinik)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var responses []PoliResponse
	for _, p := range polis {
		responses = append(responses, PoliResponse{
			ID:             p.ID,
			NamaPoli:       p.NamaPoli,
			Deskripsi:      p.Deskripsi,
			IDDataKlinik:   p.IDDataKlinik,
			Aktif:          p.Aktif,
			DibuatPada:     p.DibuatPada,
			DiperbaruiPada: p.DiperbaruiPada,
		})
	}
	c.JSON(http.StatusOK, responses)
}
