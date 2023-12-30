package services

import (
	"fmt"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
	"github.com/neutrin/hospital_booking_managment_system/internal/models"
)

type DoctorService struct {
	repo repositiories.DoctorRepo
}

func NewDoctorService(repo repositiories.DoctorRepo) *DoctorService {
	return &DoctorService{
		repo: repo,
	}
}

func (service *DoctorService) RedisterDoc(name string, designation domain.DoctorEnum) string {
	doctor := domain.NewDoctor(name, designation)
	service.repo.Save(doctor)
	resp := fmt.Sprintf("Welcome Dr. %s", doctor.Name)
	return resp
}

// todo validation missinfg that all slots hsould be between 9 am and 9pm
func (service *DoctorService) MarkDocAvailable(name string, reqSlots []models.SlotsReq) (string, error) {
	var resp string
	var slots = make([]*domain.Slot, 0)
	for _, slotReq := range reqSlots {
		curSlot, err := domain.NewSlot(slotReq.StartTime, slotReq.EndTime)
		if err != nil {
			fmt.Printf(" failed for slot = %s startTime and endTime = %s", slotReq.StartTime, slotReq.EndTime)
			return resp, err
		}
		slots = append(slots, curSlot)
	}
	doctor, err := service.repo.Find(name)
	if err != nil {
		return resp, err
	}
	for index := range slots {
		doctor.AddSlot(slots[index])
	}
	service.repo.Save(doctor)
	resp = fmt.Sprintf("Done Doc !")
	return resp, err

}

// find all doc by designation
// For each doc found find all its slots which are available
// Create a corresponding response and return accordingly
func (service *DoctorService) ShowAvailByspeciality(desg domain.DoctorEnum) (resp []models.DoctorDetBySpeciality) {
	responses := make([]models.DoctorDetBySpeciality, 0)
	doctors := service.repo.FindAll()
	for _, curDoctor := range doctors {
		if curDoctor.Designation == desg {
			for _, curSlot := range curDoctor.Slot {
				if curSlot.Status == domain.Available {
					responses = append(responses, models.DoctorDetBySpeciality{
						Name:      curDoctor.Name,
						StartTime: curSlot.StartTime.Format(domain.TimeFromat),
						EndTime:   curSlot.EndTime.Format(domain.TimeFromat),
					})
				}
			}
		}
	}

	return responses
}

func filterByDesg(doctors []*domain.Doctor, desg domain.DoctorEnum) []*domain.Doctor {
	filteredDoc := make([]*domain.Doctor, 0)
	for index := range doctors {
		if doctors[index].Designation == desg {
			filteredDoc = append(filteredDoc, doctors[index])
		}
	}
	return filteredDoc
}
