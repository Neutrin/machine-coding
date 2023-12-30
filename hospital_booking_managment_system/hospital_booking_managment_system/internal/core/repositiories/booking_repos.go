package repositiories

import "github.com/neutrin/hospital_booking_managment_system/internal/core/domain"

type BookingRepo interface {
	Save(booking *domain.Booking)
	Find(id int) (*domain.Booking, error)
	FindAll() []*domain.Booking
}
