package repositiories

import (
	"fmt"
	"strings"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
)

type PatientRepoImpl struct {
	patients []domain.Patient
}

func NewPatientRepoImpl() repositiories.PatientRepo {
	return &PatientRepoImpl{
		patients: make([]domain.Patient, 0),
	}
}

/*
 */

func (pat *PatientRepoImpl) Save(pateint *domain.Patient) {
	foundIdx := -1
	for index := range pat.patients {
		if pat.patients[index].Id == pateint.Id {
			foundIdx = index
			break
		}
	}
	if foundIdx == -1 {
		pat.patients = append(pat.patients, *pateint)
		return
	}
	pat.patients[foundIdx] = *pateint
}

func (pat *PatientRepoImpl) Find(name string) (*domain.Patient, error) {
	foundIdx := -1
	for index := range pat.patients {
		if strings.Compare(pat.patients[index].Name, name) == 0 {
			foundIdx = index
			break
		}
	}
	if foundIdx == -1 {
		return nil, fmt.Errorf(" pateint with name = %s not found", name)
	}
	return &pat.patients[foundIdx], nil
}

func (pat *PatientRepoImpl) FindAll() []*domain.Patient {
	pateints := make([]*domain.Patient, 0)
	for index := range pat.patients {
		pateints = append(pateints, &pat.patients[index])
	}
	return pateints
}
