package repositiories

import (
	"fmt"
	"strings"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
)

type DoctorRepoImpl struct {
	doctor []domain.Doctor
}

func NewDoctorRepoImpl() repositiories.DoctorRepo {
	return &DoctorRepoImpl{
		doctor: make([]domain.Doctor, 0),
	}
}

func (doc *DoctorRepoImpl) Save(doctor *domain.Doctor) {
	foundIdx := -1
	for index := range doc.doctor {
		if doc.doctor[index].Id == doctor.Id {
			foundIdx = index
			break
		}
	}

	if foundIdx == -1 {

		doc.doctor = append(doc.doctor, *doctor)
		return
	}
	doc.doctor[foundIdx] = *doctor
}
func (doc *DoctorRepoImpl) Find(name string) (*domain.Doctor, error) {
	foundIdx := -1

	for index := range doc.doctor {

		if strings.Compare(doc.doctor[index].Name, name) == 0 {
			foundIdx = index
			break
		}
	}
	if foundIdx == -1 {
		return nil, fmt.Errorf(" doctor with name = %s not found", name)
	}
	doctor := doc.doctor[foundIdx]
	return &doctor, nil
}
func (doc *DoctorRepoImpl) FindAll() []*domain.Doctor {
	doctors := make([]*domain.Doctor, 0)
	for index := range doc.doctor {
		doctors = append(doctors, &doc.doctor[index])
	}
	return doctors
}
