package repositiories

import "github.com/neutrin/hospital_booking_managment_system/internal/core/domain"

type PatientRepo interface {
	Save(pateint *domain.Patient)
	Find(name string) (*domain.Patient, error)
	FindAll() []*domain.Patient
}
