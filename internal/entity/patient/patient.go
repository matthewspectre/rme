package patient

import "time"

// Domain entity for patients.

type Patient struct {
	ID             int
	Name           string
	AdmissionDate  time.Time
	NIK            string
	Gender         string
	BloodType      string
	BirthPlaceDate string
	Phone          string
	Address        string
	Category       string
	Job            string
	IDDataKlinik   int
}
