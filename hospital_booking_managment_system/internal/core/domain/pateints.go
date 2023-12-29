package domain

var patientId int

type Patient struct {
	Id   int
	Name string
}

func NewPatient(name string) *Patient {
	patientId = patientId + 1
	return &Patient{
		Id:   patientId,
		Name: name,
	}
}
