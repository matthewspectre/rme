package patient

// HTTP handlers for patient endpoints.

import (
	"net/http"
	"time"

	entity "rme/internal/entity/patient"
	usecase "rme/internal/usecase/patient"

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

// Create menangani POST /patients
func (h *Handler) Create(c *gin.Context) {
	var req PatientCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	const layout = "2006-01-02 15:04:05"
	var admission time.Time
	var err error
	if req.AdmissionDate != "" {
		admission, err = time.Parse(layout, req.AdmissionDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid admission_date format, use YYYY-MM-DD HH:MM:SS"})
			return
		}
	} else {
		admission = time.Now()
	}

	data := &entity.Patient{
		Name:           req.Name,
		AdmissionDate:  admission,
		NIK:            req.NIK,
		Gender:         req.Gender,
		BloodType:      req.BloodType,
		BirthPlaceDate: req.BirthPlaceDate,
		Phone:          req.Phone,
		Address:        req.Address,
		Category:       req.Category,
		Job:            req.Job,
		IDDataKlinik:   req.IDDataKlinik,
	}

	updated, err := h.uc.CreateOrUpdateByNIK(data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	statusCode := http.StatusCreated
	message := "patient created"
	if updated {
		statusCode = http.StatusOK
		message = "pasien sudah terdaftar, dan diperbaharui"
	}

	c.JSON(statusCode, gin.H{"message": message})
}

// GetAll menangani GET /patients
func (h *Handler) GetAll(c *gin.Context) {
	list, err := h.uc.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]PatientResponse, 0, len(list))
	for _, p := range list {
		resp = append(resp, PatientResponse{
			ID:             p.ID,
			Name:           p.Name,
			AdmissionDate:  p.AdmissionDate,
			NIK:            p.NIK,
			Gender:         p.Gender,
			BloodType:      p.BloodType,
			BirthPlaceDate: p.BirthPlaceDate,
			Phone:          p.Phone,
			Address:        p.Address,
			Category:       p.Category,
			Job:            p.Job,
			IDDataKlinik:   p.IDDataKlinik,
		})
	}

	c.JSON(http.StatusOK, resp)
}
