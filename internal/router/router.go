package router

// HTTP router setup for anamnesis and other modules.

import (
	hanamnesis "rme/internal/handler/anamnesis"
	hauth "rme/internal/handler/auth"
	hpatient "rme/internal/handler/patient"

	"github.com/gin-gonic/gin"
)

// RegisterAnamnesisRoutes mendaftarkan endpoint /anamnesis ke Gin router.
// Contoh pemakaian di main:
//
//	r := gin.Default()
//	handler := anamnesis.NewHandler(uc)
//	router.RegisterAnamnesisRoutes(r, handler)
//	r.Run(":8080")

func RegisterAnamnesisRoutes(r *gin.Engine, handler *hanamnesis.Handler) {
	group := r.Group("/anamnesis")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.GET("/:id_pasien", handler.GetByID)
}

// RegisterAuthRoutes mendaftarkan endpoint /auth ke Gin router.
func RegisterAuthRoutes(r *gin.Engine, handler *hauth.Handler) {
	group := r.Group("/auth")

	group.POST("/login", handler.Login)
}

// RegisterPatientRoutes mendaftarkan endpoint /patients ke Gin router.
func RegisterPatientRoutes(r *gin.Engine, handler *hpatient.Handler) {
	group := r.Group("/patients")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
}
