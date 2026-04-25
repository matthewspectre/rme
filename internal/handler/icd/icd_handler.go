package icd

import (
	"net/http"
	usecase "rme/internal/usecase/icd"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

// GetAll handles GET /icd10
func (h *Handler) GetAll(c *gin.Context) {
	rows, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rows)
}
