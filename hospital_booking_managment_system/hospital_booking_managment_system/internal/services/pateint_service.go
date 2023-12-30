package services

import (
	"fmt"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/core/repositiories"
)

type PateintService struct {
	repo repositiories.PatientRepo
}

func NewPatientService(repo repositiories.PatientRepo) *PateintService {
	return &PateintService{
		repo: repo,
	}
}

func (service *PateintService) Register(name string) string {
	pateint := domain.NewPatient(name)
	service.repo.Save(pateint)
	resp := fmt.Sprintf("%s registered succesfully!!!", name)
	return resp
}
