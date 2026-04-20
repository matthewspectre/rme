package anamnesis

// HTTP handlers for anamnesis endpoints.

import (
	"net/http"
	"strconv"
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
		IDPasien:                req.IDPasien,
		IDDokter:                req.IDDokter,
		Text:                    req.Text,
		DateMake:                dateMake,
		DateUpdate:              dateUpdate,
		IDDataKlinik:            req.IDDataKlinik,
		RiwayatPengobatan:       req.RiwayatPengobatan,
		RiwayatKeluarga:         req.RiwayatKeluarga,
		RiwayatPekerjaan:        req.RiwayatPekerjaan,
		RiwayatAutoanamnesis:    req.RiwayatAutoanamnesis,
		RiwayatPenyakitDahulu:   req.RiwayatPenyakitDahulu,
		RiwayatPenyakitSekarang: req.RiwayatPenyakitSekarang,
		RiwayatPenyakitLain:     req.RiwayatPenyakitLain,
		HubunganPasien:          req.HubunganPasien,
		RiwayatAnestesiBedah:    req.RiwayatAnestesiBedah,
		RiwayatKeluhanUtama:     req.RiwayatKeluhanUtama,
		StatusKehamilan:         req.StatusKehamilan,
		KeluhanTambahan:         req.KeluhanTambahan,
		Catatan:                 req.Catatan,
		Visible:                 1,
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
	idPoliStr := c.Query("idPoli")
	if idStr == "" || idPoliStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id_pasien and idPoli are required"})
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id_pasien"})
		return
	}
	idPoli, err := strconv.Atoi(idPoliStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPoli"})
		return
	}

	data, err := h.uc.GetByID(id, idPoli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if data == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	resp := AnamnesisResponse{
		IDPasien:                data.IDPasien,
		IDDokter:                data.IDDokter,
		Text:                    data.Text,
		DateMake:                data.DateMake,
		DateUpdate:              data.DateUpdate,
		IDDataKlinik:            data.IDDataKlinik,
		RiwayatPengobatan:       data.RiwayatPengobatan,
		RiwayatKeluarga:         data.RiwayatKeluarga,
		RiwayatPekerjaan:        data.RiwayatPekerjaan,
		RiwayatAutoanamnesis:    data.RiwayatAutoanamnesis,
		RiwayatPenyakitDahulu:   data.RiwayatPenyakitDahulu,
		RiwayatPenyakitSekarang: data.RiwayatPenyakitSekarang,
		RiwayatPenyakitLain:     data.RiwayatPenyakitLain,
		HubunganPasien:          data.HubunganPasien,
		RiwayatAnestesiBedah:    data.RiwayatAnestesiBedah,
		RiwayatKeluhanUtama:     data.RiwayatKeluhanUtama,
		StatusKehamilan:         data.StatusKehamilan,
		KeluhanTambahan:         data.KeluhanTambahan,
		Catatan:                 data.Catatan,
		Visible:                 data.Visible,
	}

	c.JSON(http.StatusOK, resp)
}

// GetAll menangani GET /anamnesis (Gin handler)
func (h *Handler) GetAll(c *gin.Context) {
	idPoliStr := c.Query("idPoli")
	if idPoliStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "idPoli is required"})
		return
	}
	idPoli, err := strconv.Atoi(idPoliStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPoli"})
		return
	}

	list, err := h.uc.GetAll(idPoli)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]AnamnesisResponse, 0, len(list))
	for _, data := range list {
		resp = append(resp, AnamnesisResponse{
			IDPasien:                data.IDPasien,
			IDDokter:                data.IDDokter,
			Text:                    data.Text,
			DateMake:                data.DateMake,
			DateUpdate:              data.DateUpdate,
			IDDataKlinik:            data.IDDataKlinik,
			RiwayatPengobatan:       data.RiwayatPengobatan,
			RiwayatKeluarga:         data.RiwayatKeluarga,
			RiwayatPekerjaan:        data.RiwayatPekerjaan,
			RiwayatAutoanamnesis:    data.RiwayatAutoanamnesis,
			RiwayatPenyakitDahulu:   data.RiwayatPenyakitDahulu,
			RiwayatPenyakitSekarang: data.RiwayatPenyakitSekarang,
			RiwayatPenyakitLain:     data.RiwayatPenyakitLain,
			HubunganPasien:          data.HubunganPasien,
			RiwayatAnestesiBedah:    data.RiwayatAnestesiBedah,
			RiwayatKeluhanUtama:     data.RiwayatKeluhanUtama,
			StatusKehamilan:         data.StatusKehamilan,
			KeluhanTambahan:         data.KeluhanTambahan,
			Catatan:                 data.Catatan,
			Visible:                 data.Visible,
		})
	}

	c.JSON(http.StatusOK, resp)
}
