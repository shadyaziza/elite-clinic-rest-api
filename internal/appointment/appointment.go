package appointment

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shadyaziza/elite-clinic-rest-api/internal"
	"time"
)

var (
	ErrFetchingAppointment = errors.New("failed to fetch appointment by time id")
	ErrNotImplemented      = errors.New("not implemented")
)

// Appointment - a representation of the appointment
// structure for our service
type Appointment struct {
	id        time.Time
	comment   string
	patientID string
	doctorID  uuid.UUID
}

// AppointmentsService - is the struct on which all our
// logic will be built on top of
type AppointmentsService struct {
	Store internal.Store
}

// NewService - returns a pointer to a new
// service
func NewService(store internal.Store) *AppointmentsService {
	return &AppointmentsService{
		Store: store,
	}
}

func (s *AppointmentsService) GetAppointment(ctx context.Context, id string) (Appointment, error) {
	fmt.Println("retrieving an appointment")
	appointment, err := s.Store.GetAppointment(ctx, id)
	if err != nil {
		// use this err to know implementation errors
		// from logs
		fmt.Println(err)
		// return our own custom errors from service
		// layer to transport layer to guard
		// our implementation details from being
		// exposed to potentially failed client
		// calls
		return Appointment{}, ErrFetchingAppointment
	}
	return Appointment{
		id:        appointment.ID,
		comment:   appointment.Comment.String,
		patientID: appointment.PatientID,
		doctorID:  appointment.DoctorID,
	}, nil

}

func (s *AppointmentsService) UpdateAppointment(ctx context.Context, updatedAppointment Appointment) error {
	return ErrNotImplemented
}

func (s *AppointmentsService) DeleteAppointment(ctx context.Context, id time.Time) error {
	return ErrNotImplemented
}

func (s *AppointmentsService) CreateAppointment(ctx context.Context, appointment Appointment) (Appointment, error) {
	return Appointment{}, ErrNotImplemented
}
