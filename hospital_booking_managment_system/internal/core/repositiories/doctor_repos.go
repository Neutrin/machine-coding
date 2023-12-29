package repositiories

import "github.com/neutrin/hospital_booking_managment_system/internal/core/domain"

type DoctorRepo interface {
	Save(doctor *domain.Doctor)
	Find(name string) (*domain.Doctor, error)
	FindAll() []*domain.Doctor
}
