package doctor

import (
	"net/http"
	"strconv"

	entity "rme/internal/entity/doctor"
	usecase "rme/internal/usecase/doctor"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req DoctorCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := &entity.Doctor{
		IDUser:       req.IDUser,
		Nama:         req.NamaDokter,
		Poli:         req.Poli,
		NomorTelepon: req.NomorTelepon,
		IDDataKlinik: req.IDDataKlinik,
	}

	if err := h.uc.Create(data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "dokter berhasil dibuat"})
}

func (h *Handler) GetAll(c *gin.Context) {
	idDataKlinikStr := c.Query("idDataKlinik")
	if idDataKlinikStr == "" {
		idDataKlinikStr = c.Query("id_data_klinik")
	}

	var (
		list []*entity.Doctor
		err  error
	)

	if idDataKlinikStr != "" {
		idDataKlinik, convErr := strconv.Atoi(idDataKlinikStr)
		if convErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDataKlinik"})
			return
		}
		list, err = h.uc.GetAllByIDDataKlinik(idDataKlinik)
	} else {
		list, err = h.uc.GetAll()
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]DoctorResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, DoctorResponse{
			ID:           d.ID,
			IDUser:       d.IDUser,
			NamaDokter:   d.Nama,
			Poli:         d.Poli,
			NomorTelepon: d.NomorTelepon,
			IDDataKlinik: d.IDDataKlinik,
		})
	}

	c.JSON(http.StatusOK, resp)
}
