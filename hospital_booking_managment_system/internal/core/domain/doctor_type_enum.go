package domain

type DoctorEnum int

const (
	Cardiologist = iota + 1
	Dermatologist
)

var doctorTypeString = map[DoctorEnum]string{
	Cardiologist:  "Cardiologist",
	Dermatologist: "Dermatologist",
}

func (doctor DoctorEnum) String() string {
	return doctorTypeString[doctor]
}
