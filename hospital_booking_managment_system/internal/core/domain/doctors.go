package domain

var doctorID int

type Doctor struct {
	Id          int
	Name        string
	Designation DoctorEnum
	Slot        []*Slot
}

func NewDoctor(name string, des DoctorEnum) *Doctor {
	doctorID = doctorID + 1
	return &Doctor{
		Id:          doctorID,
		Name:        name,
		Designation: des,
		Slot:        make([]*Slot, 0),
	}
}

func (doctor *Doctor) AddSlot(slot *Slot) {
	doctor.Slot = append(doctor.Slot, slot)
}
