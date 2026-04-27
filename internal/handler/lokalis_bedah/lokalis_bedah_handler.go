package lokalis_bedah

import (
	"net/http"
	"strconv"
	"time"

	entity "rme/internal/entity/lokalis_bedah"
	usecase "rme/internal/usecase/lokalis_bedah"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	uc usecase.Usecase
}

func NewHandler(uc usecase.Usecase) *Handler {
	return &Handler{uc: uc}
}

func (h *Handler) Create(c *gin.Context) {
	var req LokalisBedahCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tanggal time.Time
	if req.Tanggal != "" {
		t, err := time.Parse(time.RFC3339, req.Tanggal)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid tanggal format, use RFC3339"})
			return
		}
		tanggal = t
	} else {
		tanggal = time.Now()
	}

	d := &entity.LokalisBedah{
		IDPasien:       req.IDPasien,
		IDDokter:       req.IDDokter,
		Tanggal:        tanggal,
		LokasiKelainan: req.LokasiKelainan,
		JenisKelainan:  req.JenisKelainan,
		Ukuran:         req.Ukuran,
		Warna:          req.Warna,
		NyeriTekan:     req.NyeriTekan,
		Konsistensi:    req.Konsistensi,
		Mobilitas:      req.Mobilitas,
		TandaRadang:    req.TandaRadang,
		Fluktuasi:      req.Fluktuasi,
		Catatan:        req.Catatan,
		Visible:        1,
	}

	if err := h.uc.Create(d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": d.ID})
}

func (h *Handler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	d, err := h.uc.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if d == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "not found"})
		return
	}

	resp := LokalisBedahResponse{
		ID:             d.ID,
		IDPasien:       d.IDPasien,
		IDDokter:       d.IDDokter,
		Tanggal:        d.Tanggal.Format(time.RFC3339),
		LokasiKelainan: d.LokasiKelainan,
		JenisKelainan:  d.JenisKelainan,
		Ukuran:         d.Ukuran,
		Warna:          d.Warna,
		NyeriTekan:     d.NyeriTekan,
		Konsistensi:    d.Konsistensi,
		Mobilitas:      d.Mobilitas,
		TandaRadang:    d.TandaRadang,
		Fluktuasi:      d.Fluktuasi,
		Catatan:        d.Catatan,
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAll(c *gin.Context) {
	var idDokterPtr *int
	var idPasienPtr *int
	idDokterStr := c.Query("idDokter")
	idPasienStr := c.Query("idPasien")
	if idDokterStr != "" {
		idd, err := strconv.Atoi(idDokterStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idDokter"})
			return
		}
		idDokterPtr = &idd
	}
	if idPasienStr != "" {
		idp, err := strconv.Atoi(idPasienStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid idPasien"})
			return
		}
		idPasienPtr = &idp
	}

	list, err := h.uc.GetAll(idDokterPtr, idPasienPtr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]LokalisBedahResponse, 0, len(list))
	for _, d := range list {
		resp = append(resp, LokalisBedahResponse{
			ID:             d.ID,
			IDPasien:       d.IDPasien,
			IDDokter:       d.IDDokter,
			Tanggal:        d.Tanggal.Format(time.RFC3339),
			LokasiKelainan: d.LokasiKelainan,
			JenisKelainan:  d.JenisKelainan,
			Ukuran:         d.Ukuran,
			Warna:          d.Warna,
			NyeriTekan:     d.NyeriTekan,
			Konsistensi:    d.Konsistensi,
			Mobilitas:      d.Mobilitas,
			TandaRadang:    d.TandaRadang,
			Fluktuasi:      d.Fluktuasi,
			Catatan:        d.Catatan,
		})
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req LokalisBedahUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.LokasiKelainan != nil {
		updates["lokasi_kelainan"] = *req.LokasiKelainan
	}
	if req.JenisKelainan != nil {
		updates["jenis_kelainan"] = *req.JenisKelainan
	}
	if req.Ukuran != nil {
		updates["ukuran"] = *req.Ukuran
	}
	if req.Warna != nil {
		updates["warna"] = *req.Warna
	}
	if req.NyeriTekan != nil {
		updates["nyeri_tekan"] = *req.NyeriTekan
	}
	if req.Konsistensi != nil {
		updates["konsistensi"] = *req.Konsistensi
	}
	if req.Mobilitas != nil {
		updates["mobilitas"] = *req.Mobilitas
	}
	if req.TandaRadang != nil {
		updates["tanda_radang"] = *req.TandaRadang
	}
	if req.Fluktuasi != nil {
		updates["fluktuasi"] = *req.Fluktuasi
	}
	if req.Catatan != nil {
		updates["catatan"] = *req.Catatan
	}
	if req.Tanggal != nil {
		updates["tanggal"] = *req.Tanggal
	}

	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no fields to update"})
		return
	}

	if err := h.uc.Update(id, updates); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "lokalis bedah updated"})
}

func (h *Handler) Hide(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	if err := h.uc.Hide(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "lokalis bedah hidden"})
}
