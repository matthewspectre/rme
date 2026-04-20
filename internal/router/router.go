package router

import (
	hanamnesis "rme/internal/handler/anamnesis"
	hantrian "rme/internal/handler/antrian"
	hauth "rme/internal/handler/auth"
	hdoctor "rme/internal/handler/doctor"
	hpatient "rme/internal/handler/patient"
	hpoli "rme/internal/handler/poli"

	"github.com/gin-gonic/gin"
)

// RegisterAntrianRoutes mendaftarkan endpoint /antrian ke Gin router.
func RegisterAntrianRoutes(r *gin.Engine, handler *hantrian.Handler) {
	group := r.Group("/antrian")
	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)
	group.PATCH("/:id", handler.Update)
}

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

	// Soft-hide endpoint
	group.PATCH("/:id/hide", handler.Hide)
	// Generic update (partial)
	group.PATCH("/:id", handler.Update)
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

// RegisterDoctorRoutes mendaftarkan endpoint /dokter ke Gin router.
func RegisterDoctorRoutes(r *gin.Engine, handler *hdoctor.Handler) {
	group := r.Group("/dokter")

	group.POST("/", handler.Create)
	group.GET("/", handler.GetAll)

	// Kompatibilitas endpoint lama
	groupLegacy := r.Group("/doctors")
	groupLegacy.POST("/", handler.Create)
	groupLegacy.GET("/", handler.GetAll)
}

// RegisterPoliRoutes mendaftarkan endpoint /poli ke Gin router.
func RegisterPoliRoutes(r *gin.Engine, handler *hpoli.Handler) {
	group := r.Group("/poli")
	group.GET("/", handler.GetAll)
}
