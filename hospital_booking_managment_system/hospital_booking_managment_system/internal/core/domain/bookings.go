package domain

var bookingId int

type Booking struct {
	Id      int
	Doctor  *Doctor
	Patient *Patient
	Slot    *Slot
	Status  StatusEnum
}

func NewBooking(doctor *Doctor, patient *Patient, slot *Slot) *Booking {
	bookingId = bookingId + 1
	return &Booking{
		Id:      bookingId,
		Doctor:  doctor,
		Patient: patient,
		Slot:    slot,
		Status:  Booked,
	}
}
