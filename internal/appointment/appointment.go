package appointment

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/shadyaziza/elite-clinic-rest-api/internal"
	db "github.com/shadyaziza/elite-clinic-rest-api/internal/db/sqlc"
	"time"
)

var (
	ErrFetchingAppointment = errors.New("failed to fetch appointment by time id")
	ErrNotImplemented      = errors.New("not implemented")
)

func creatAppointment(appointment db.Appointment) Appointment {
	return Appointment{
		id:        appointment.ID,
		comment:   appointment.Comment.String,
		patientID: appointment.PatientID,
		doctorID:  appointment.DoctorID,
	}
}

// Appointment - a representation of the appointment
// structure for our service
type Appointment struct {
	id        time.Time
	comment   string
	patientID string
	doctorID  uuid.UUID
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store internal.Store
}

// NewService - returns a pointer to a new
// service
func NewService(store internal.Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetAppointment(ctx context.Context, id string) (Appointment, error) {
	fmt.Println("retrieving an appointment")
	appointment, err := s.Store.GetAppointment(ctx, id)
	if err != nil {
		return appoitnmentServiceError(Appointment{}, err, ErrFetchingAppointment)
	}
	return creatAppointment(appointment), nil

}

func appoitnmentServiceError[K any](data K, implError error, serviceError error) (K, error) {

	// use this err to know implementation errors
	// from logs
	fmt.Println(implError)
	// return our own custom errors from service
	// layer to transport layer to guard
	// our implementation details from being
	// exposed to potentially failed client
	// calls
	return data, serviceError

}

func (s *Service) UpdateAppointment(ctx context.Context, updatedAppointment Appointment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteAppointment(ctx context.Context, id time.Time) error {
	return ErrNotImplemented
}

func (s *Service) CreateAppointment(ctx context.Context, appointment Appointment) (Appointment, error) {
	return Appointment{}, ErrNotImplemented
}
