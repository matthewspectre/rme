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
	handlerAntrian "rme/internal/handler/antrian"
	handlerAuth "rme/internal/handler/auth"
	handlerDiagnosis "rme/internal/handler/diagnosis"
	handlerDoctor "rme/internal/handler/doctor"
	handlerICD "rme/internal/handler/icd"
	handlerPatient "rme/internal/handler/patient"
	handlerPemeriksaanEkg "rme/internal/handler/pemeriksaan_ekg"
	handlerPemeriksaanLaboratorium "rme/internal/handler/pemeriksaan_laboratorium"
	handlerPemeriksaanVital "rme/internal/handler/pemeriksaan_vital"
	handlerPoli "rme/internal/handler/poli"
	handlerRujukUlang "rme/internal/handler/rujuk_ulang"
	handlerTatalaksana "rme/internal/handler/tatalaksana"

	mysqlAnamnesis "rme/internal/mysql/anamnesis"
	mysqlAntrian "rme/internal/mysql/antrian"
	mysqlDiagnosis "rme/internal/mysql/diagnosis"
	mysqlDoctor "rme/internal/mysql/doctor"
	mysqlICD "rme/internal/mysql/icd"
	mysqlPatient "rme/internal/mysql/patient"
	mysqlPemeriksaanEkg "rme/internal/mysql/pemeriksaan_ekg"
	mysqlPemeriksaanLaboratorium "rme/internal/mysql/pemeriksaan_laboratorium"
	mysqlPemeriksaanVital "rme/internal/mysql/pemeriksaan_vital"
	mysqlPoli "rme/internal/mysql/poli"
	mysqlRujukUlang "rme/internal/mysql/rujuk_ulang"
	mysqlTatalaksana "rme/internal/mysql/tatalaksana"
	mysqlUser "rme/internal/mysql/user"

	repoAntrianInterface "rme/internal/repository/antrian"
	repoPoliInterface "rme/internal/repository/poli"
	router "rme/internal/router"
	usecaseAnamnesis "rme/internal/usecase/anamnesis"
	usecaseAntrian "rme/internal/usecase/antrian"
	usecaseAuth "rme/internal/usecase/auth"
	usecaseDiagnosis "rme/internal/usecase/diagnosis"
	usecaseDoctor "rme/internal/usecase/doctor"
	usecaseICD "rme/internal/usecase/icd"
	usecasePatient "rme/internal/usecase/patient"
	usecasePemeriksaanEkg "rme/internal/usecase/pemeriksaan_ekg"
	usecasePemeriksaanLaboratorium "rme/internal/usecase/pemeriksaan_laboratorium"
	usecasePemeriksaanVital "rme/internal/usecase/pemeriksaan_vital"
	usecasePoli "rme/internal/usecase/poli"
	usecaseRujukUlang "rme/internal/usecase/rujuk_ulang"
	usecaseTatalaksana "rme/internal/usecase/tatalaksana"
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

	// Wiring clean architecture untuk Doctor
	repoDoctor := mysqlDoctor.NewRepository(db)
	ucDoctor := usecaseDoctor.NewUsecase(repoDoctor)
	handlerDoctorHTTP := handlerDoctor.NewHandler(ucDoctor)

	// Wiring clean architecture untuk Poli
	repoPoli := mysqlPoli.NewRepositoryMySQL(db)
	var repoPoliIface repoPoliInterface.Repository = repoPoli
	ucPoli := usecasePoli.NewUsecase(repoPoliIface)
	handlerPoliHTTP := handlerPoli.NewHandler(ucPoli)

	// Wiring clean architecture untuk Antrian
	repoAntrian := mysqlAntrian.NewRepositoryMySQL(db)
	var repoAntrianIface repoAntrianInterface.Repository = repoAntrian
	ucAntrian := usecaseAntrian.NewUsecase(repoAntrianIface)
	handlerAntrianHTTP := handlerAntrian.NewHandler(ucAntrian)

	// Wiring pemeriksaan_vital
	repoPemeriksaanVital := mysqlPemeriksaanVital.NewRepository(db)
	ucPemeriksaanVital := usecasePemeriksaanVital.NewUsecase(repoPemeriksaanVital)
	handlerPemeriksaanVitalHTTP := handlerPemeriksaanVital.NewHandler(ucPemeriksaanVital)

	// Wiring pemeriksaan_laboratorium
	repoPemeriksaanLaboratorium := mysqlPemeriksaanLaboratorium.NewRepository(db)
	ucPemeriksaanLaboratorium := usecasePemeriksaanLaboratorium.NewUsecase(repoPemeriksaanLaboratorium)
	handlerPemeriksaanLaboratoriumHTTP := handlerPemeriksaanLaboratorium.NewHandler(ucPemeriksaanLaboratorium)

	// Wiring pemeriksaan_ekg
	repoPemeriksaanEkg := mysqlPemeriksaanEkg.NewRepository(db)
	ucPemeriksaanEkg := usecasePemeriksaanEkg.NewUsecase(repoPemeriksaanEkg)
	handlerPemeriksaanEkgHTTP := handlerPemeriksaanEkg.NewHandler(ucPemeriksaanEkg)

	// Wiring ICD
	repoICD := mysqlICD.NewRepository(db)
	ucICD := usecaseICD.NewUsecase(repoICD)
	handlerICDHTTP := handlerICD.NewHandler(ucICD)

	// Wiring tatalaksana
	repoTatalaksana := mysqlTatalaksana.NewRepository(db)
	ucTatalaksana := usecaseTatalaksana.NewUsecase(repoTatalaksana)
	handlerTatalaksanaHTTP := handlerTatalaksana.NewHandler(ucTatalaksana)

	// Wiring diagnosis
	repoDiagnosis := mysqlDiagnosis.NewMySQLRepo(db)
	ucDiagnosis := usecaseDiagnosis.NewUsecase(repoDiagnosis)
	handlerDiagnosisHTTP := handlerDiagnosis.NewHandler(ucDiagnosis)
	repoRujukUlang := mysqlRujukUlang.NewRepository(db)
	ucRujukUlang := usecaseRujukUlang.NewUsecase(repoRujukUlang)
	handlerRujukUlangHTTP := handlerRujukUlang.NewHandler(ucRujukUlang)
	r := gin.Default()
	r.Use(cors.Default())
	router.RegisterAnamnesisRoutes(r, handlerAnam)
	router.RegisterAuthRoutes(r, handlerAuthHTTP)
	router.RegisterPatientRoutes(r, handlerPatientHTTP)
	router.RegisterDoctorRoutes(r, handlerDoctorHTTP)
	router.RegisterPoliRoutes(r, handlerPoliHTTP)
	router.RegisterAntrianRoutes(r, handlerAntrianHTTP)
	router.RegisterPemeriksaanVitalRoutes(r, handlerPemeriksaanVitalHTTP)
	router.RegisterPemeriksaanLaboratoriumRoutes(r, handlerPemeriksaanLaboratoriumHTTP)
	router.RegisterPemeriksaanEkgRoutes(r, handlerPemeriksaanEkgHTTP)
	router.RegisterICDRoutes(r, handlerICDHTTP)
	router.RegisterTatalaksanaRoutes(r, handlerTatalaksanaHTTP)
	router.RegisterRujukUlangRoutes(r, handlerRujukUlangHTTP)
	router.RegisterDiagnosisRoutes(r, handlerDiagnosisHTTP)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
