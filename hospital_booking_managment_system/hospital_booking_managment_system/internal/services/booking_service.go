package services

import (
	"fmt"
	"strings"
	"time"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
	"github.com/neutrin/hospital_booking_managment_system/internal/models"
)

type BookingService struct {
	repo repositiories.BookingRepo
	//need to check this should be service or repo
	doctorRepo  repositiories.DoctorRepo
	pateintRepo repositiories.PatientRepo
}

func NewBookingService(repo repositiories.BookingRepo, doctorRepo repositiories.DoctorRepo, pateintRepo repositiories.PatientRepo) *BookingService {
	return &BookingService{
		repo:        repo,
		doctorRepo:  doctorRepo,
		pateintRepo: pateintRepo,
	}
}

// find all active bookings of patient
// Check no booking with same time
// Find doctor
// Check slots of doctor
// if availabel make booking and slot unavaiolabel
// If not add booking to pending status
func (service *BookingService) BookAppointMent(patientName string, docName string, inputTime string) (string, error) {
	var (
		err     error
		resp    string
		pateint *domain.Patient
		doctor  *domain.Doctor
	)
	curTime, err := time.Parse(domain.TimeFromat, inputTime)
	if err != nil {
		return resp, err
	}
	allBookings := service.repo.FindAll()
	for _, curBooking := range allBookings {
		if strings.Compare(curBooking.Patient.Name, patientName) == 0 {
			if curBooking.Slot.StartTime.Equal(curTime) {
				return resp, fmt.Errorf(" booking exists with doctor = %s", curBooking.Doctor.Name)
			}
		}
	}
	doctor, err = service.doctorRepo.Find(docName)
	if err != nil {
		return resp, err
	}
	pateint, err = service.pateintRepo.Find(patientName)
	if err != nil {
		return resp, err
	}
	slot := &domain.Slot{
		StartTime: curTime,
		EndTime:   curTime.Add(30 * time.Minute),
	}
	booking := domain.NewBooking(doctor, pateint, slot)
	for _, curSlot := range doctor.Slot {
		if curSlot.StartTime.Equal(curTime) {
			if curSlot.Status == domain.Booked {
				booking.Status = domain.Waiting
			} else {
				curSlot.Status = domain.Booked
			}
		}
	}
	service.repo.Save(booking)
	service.doctorRepo.Save(doctor)

	return fmt.Sprintf(" Booked Booked id = %d and status = %s", booking.Id, booking.Status), nil
}

func (service *BookingService) CancelBooking(bookingId int) ([]string, error) {
	resp := make([]string, 0)
	curBooking, err := service.repo.Find(bookingId)
	if err != nil {
		return resp, err
	}
	curBooking.Status = domain.Cancelled
	doctor, err := service.doctorRepo.Find(curBooking.Doctor.Name)
	if err != nil {
		return resp, err
	}

	service.repo.Save(curBooking)

	resp = append(resp, fmt.Sprintf("booking cancelled = %d", curBooking.Id))
	for _, bookings := range service.repo.FindAll() {
		if strings.Compare(bookings.Doctor.Name, doctor.Name) == 0 && bookings.Status == domain.Waiting && bookings.Slot.StartTime.Equal(curBooking.Slot.StartTime) {
			bookings.Status = domain.Available
			resp = append(resp, fmt.Sprintf("Booking condfiremd for booking = %d", bookings.Id))
			service.repo.Save(bookings)
			return resp, nil
		}
	}
	for _, curSlot := range doctor.Slot {
		if curSlot.StartTime.Equal(curBooking.Slot.StartTime) {
			curSlot.Status = domain.Available
			break
		}
	}
	service.doctorRepo.Save(doctor)
	return resp, nil
}

func (service *BookingService) ShowbookingByPateint(patientName string) ([]models.BookingPateintResp, error) {
	resp := make([]models.BookingPateintResp, 0)

	for _, curBooking := range service.repo.FindAll() {
		if strings.Compare(curBooking.Patient.Name, patientName) == 0 && curBooking.Status != domain.Cancelled {
			resp = append(resp, models.BookingPateintResp{
				BookingId: curBooking.Id,
				DocName:   curBooking.Doctor.Name,
				Time:      curBooking.Slot.StartTime.Format(domain.TimeFromat),
			})
		}
	}
	return resp, nil
}
