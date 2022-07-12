package appointment

import (
	"context"
	"errors"
	"fmt"
	"time"
)

var (
	ErrFetchingAppointment = errors.New("failed to fetch appointment by time id")
	ErrNotImplemented      = errors.New("Not implemented")
)

// Appointment - a representation of the appointment
// structure for our service
type Appointment struct {
	id        time.Time
	comment   string
	patientID string
	doctorID  string
}

// Store - this interface defines all of the methods
// that  our service needs to operate
type Store interface {
	GetComment(ctx context.Context, id time.Time) (Appointment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
// service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id time.Time) (Appointment, error) {
	fmt.Println("retrieving a comment")
	appointment, err := s.Store.GetComment(ctx, id)
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
	return appointment, nil

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
