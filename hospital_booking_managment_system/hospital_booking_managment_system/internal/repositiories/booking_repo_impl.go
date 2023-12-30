package repositiories

import (
	"fmt"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
)

type BookingRepoImpl struct {
	bookings []domain.Booking
}

/*
 */
func NewBookingRepoImpl() repositiories.BookingRepo {
	return &BookingRepoImpl{
		bookings: make([]domain.Booking, 0),
	}
}

func (book *BookingRepoImpl) Save(booking *domain.Booking) {
	foundIdx := -1
	for index := range book.bookings {
		if book.bookings[index].Id == booking.Id {
			foundIdx = index
			break
		}
	}
	if foundIdx == -1 {
		book.bookings = append(book.bookings, *booking)
		return
	}
	book.bookings[foundIdx] = *booking
}

func (book *BookingRepoImpl) Find(id int) (*domain.Booking, error) {
	foundIdx := -1
	for index := range book.bookings {
		if book.bookings[index].Id == id {
			foundIdx = index
			break
		}
	}
	if foundIdx == -1 {
		return nil, fmt.Errorf(" booking with id = %d not found", id)
	}
	return &book.bookings[foundIdx], nil
}

func (book *BookingRepoImpl) FindAll() []*domain.Booking {
	bookings := make([]*domain.Booking, 0)
	for index := range book.bookings {
		bookings = append(bookings, &book.bookings[index])
	}
	return bookings
}
