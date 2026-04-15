package auth

// HTTP handlers for authentication endpoints.

import (
	"net/http"

	usecase "rme/internal/usecase/auth"

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

// Login menangani POST /auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.uc.Login(req.Username, req.Password)
	if err != nil {
		// Jangan bocorkan detail kesalahan ke client
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	// Tentukan nama role berdasarkan kode angka
	var roleName string
	switch user.Role {
	case 1:
		roleName = "Admin"
	case 2:
		roleName = "Dokter"
	case 3:
		roleName = "Front Office"
	default:
		roleName = "Unknown"
	}

	resp := LoginResponse{
		ID:       user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Role:     user.Role,
		RoleName: roleName,
	}

	c.JSON(http.StatusOK, resp)
}
