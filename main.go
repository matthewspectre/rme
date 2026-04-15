package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	handlerAnamnesis "rme/internal/handler/anamnesis"
	handlerAuth "rme/internal/handler/auth"
	handlerPatient "rme/internal/handler/patient"
	mysqlAnamnesis "rme/internal/mysql/anamnesis"
	mysqlPatient "rme/internal/mysql/patient"
	mysqlUser "rme/internal/mysql/user"
	router "rme/internal/router"
	usecaseAnamnesis "rme/internal/usecase/anamnesis"
	usecaseAuth "rme/internal/usecase/auth"
	usecasePatient "rme/internal/usecase/patient"
)

func main() {
	// Ambil konfigurasi DB dari environment (sesuai launch.json)
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	if user == "" || host == "" || port == "" || name == "" {
		log.Fatal("database environment variables (DB_USER, DB_HOST, DB_PORT, DB_NAME) must not be empty")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	// Inisialisasi GORM MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Wiring clean architecture untuk Anamnesis
	repoAnamnesis := mysqlAnamnesis.NewRepository(db)
	ucAnamnesis := usecaseAnamnesis.NewUsecase(repoAnamnesis)
	handlerAnam := handlerAnamnesis.NewHandler(ucAnamnesis)

	// Wiring clean architecture untuk Auth/Login
	repoUser := mysqlUser.NewRepository(db)
	ucAuth := usecaseAuth.NewUsecase(repoUser)
	handlerAuthHTTP := handlerAuth.NewHandler(ucAuth)

	// Wiring clean architecture untuk Patient
	repoPatient := mysqlPatient.NewRepository(db)
	ucPatient := usecasePatient.NewUsecase(repoPatient)
	handlerPatientHTTP := handlerPatient.NewHandler(ucPatient)

	// Setup Gin router
	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterAnamnesisRoutes(r, handlerAnam)
	router.RegisterAuthRoutes(r, handlerAuthHTTP)
	router.RegisterPatientRoutes(r, handlerPatientHTTP)

	// Jalankan HTTP server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
