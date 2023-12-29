package main

import (
	"fmt"

	"github.com/neutrin/hospital_booking_managment_system/internal/core/domain"
	"github.com/neutrin/hospital_booking_managment_system/internal/models"
	booking_repo "github.com/neutrin/hospital_booking_managment_system/internal/repositiories"
	doctor_repo "github.com/neutrin/hospital_booking_managment_system/internal/repositiories"
	pateint_repo "github.com/neutrin/hospital_booking_managment_system/internal/repositiories"
	"github.com/neutrin/hospital_booking_managment_system/internal/services"
)

func main() {
	repo := doctor_repo.NewDoctorRepoImpl()
	pateintRepo := pateint_repo.NewPatientRepoImpl()
	doctorService := services.NewDoctorService(repo)
	pateintService := services.NewPatientService(pateintRepo)
	bookingService := services.NewBookingService(booking_repo.NewBookingRepoImpl(), repo, pateintRepo)
	//registerDoc -> Curious-> Cardiologist
	resp := doctorService.RedisterDoc("Curious", domain.Cardiologist)
	fmt.Println(resp)
	//markDocAvail: Curious 9:30-10:30
	resp, err := doctorService.MarkDocAvailable("Curious", []models.SlotsReq{
		models.SlotsReq{
			StartTime: "09:30",
			EndTime:   "10:30",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	//markDocAvail: Curious 9:30-10:00, 12:30-13:00, 16:00-16:30
	resp, err = doctorService.MarkDocAvailable("Curious", []models.SlotsReq{
		models.SlotsReq{
			StartTime: "09:30",
			EndTime:   "10:00",
		},
		models.SlotsReq{
			StartTime: "12:30",
			EndTime:   "13:00",
		},
		models.SlotsReq{
			StartTime: "16:00",
			EndTime:   "16:30",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//registerDoc -> Dreadful-> Dermatologist
	resp = doctorService.RedisterDoc("Dreadful", domain.Dermatologist)
	fmt.Println(resp)
	//markDocAvail: Dreadful 9:30-10:00, 12:30-13:00, 16:00-16:30
	resp, err = doctorService.MarkDocAvailable("Dreadful", []models.SlotsReq{
		models.SlotsReq{
			StartTime: "09:30",
			EndTime:   "10:00",
		},
		models.SlotsReq{
			StartTime: "12:30",
			EndTime:   "13:00",
		},
		models.SlotsReq{
			StartTime: "16:00",
			EndTime:   "16:30",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)
	//showAvailByspeciality: Cardiologist
	responses := doctorService.ShowAvailByspeciality(domain.Cardiologist)
	for _, curResp := range responses {
		fmt.Println(curResp.String())
	}
	//registerPatient ->PatientA
	resp = pateintService.Register("PatientA")
	fmt.Println(resp)
	//bookAppointment: (PatientA, Dr.Curious, 12:30)
	resp, err = bookingService.BookAppointMent("PatientA", "Curious", "12:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
	//showAvailByspeciality: Cardiologist
	responses = doctorService.ShowAvailByspeciality(domain.Cardiologist)
	for _, curResp := range responses {
		fmt.Println(curResp.String())
	}
	// cancelBookingId: 1234
	cancelBookingResp, err := bookingService.CancelBooking(1)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, curCancelResp := range cancelBookingResp {
			fmt.Println(curCancelResp)
		}
	}
	//showAvailByspeciality: Cardiologist
	responses = doctorService.ShowAvailByspeciality(domain.Cardiologist)
	for _, curResp := range responses {
		fmt.Println(curResp.String())
	}
	resp = pateintService.Register("PatientB")
	//bookAppointment: (PatientB, Dr.Curious, 12:30)
	resp, err = bookingService.BookAppointMent("PatientB", "Curious", "12:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	//registerDoc -> Daring-> Dermatologist
	resp = doctorService.RedisterDoc("Daring", domain.Dermatologist)
	fmt.Println(resp)

	resp, err = doctorService.MarkDocAvailable("Daring", []models.SlotsReq{
		models.SlotsReq{
			StartTime: "11:30",
			EndTime:   "12:00",
		},
		models.SlotsReq{
			StartTime: "14:00",
			EndTime:   "14:30",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(resp)

	//showAvailByspeciality: Dermatologist
	responses = doctorService.ShowAvailByspeciality(domain.Dermatologist)
	for _, curResp := range responses {
		fmt.Println(curResp.String())
	}
	pateintService.Register("PatientC")
	pateintService.Register("PatientD")
	pateintService.Register("PatientE")
	pateintService.Register("PatientF")
	pateintService.Register("PatientG")

	resp, err = bookingService.BookAppointMent("PatientC", "Dreadful", "09:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	resp, err = bookingService.BookAppointMent("PatientD", "Dreadful", "12:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	resp, err = bookingService.BookAppointMent("PatientE", "Dreadful", "16:00")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	resp, err = bookingService.BookAppointMent("PatientF", "Daring", "11:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	resp, err = bookingService.BookAppointMent("PatientG", "Daring", "14:00")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	resp, err = bookingService.BookAppointMent("PatientG", "Daring", "11:30")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
	cancelBookingResp, err = bookingService.CancelBooking(6)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, curCancelResp := range cancelBookingResp {
			fmt.Println(curCancelResp)
		}
	}

	responses = doctorService.ShowAvailByspeciality(domain.Dermatologist)
	for _, curResp := range responses {
		fmt.Println(curResp.String())
	}

}
