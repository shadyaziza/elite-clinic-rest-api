package appointment

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

var (
	ErrFetchingAppointment = errors.New("failed to fetch appointment by time id")
	ErrCreatingAppointment = errors.New("failed to create appointment")
	ErrUpdatingAppointment = errors.New("failed to update appointment")
	ErrDeletingAppointment = errors.New("failed to delete appointment")
	ErrNotImplemented      = errors.New("not implemented")
)

// Store - this interface defines all the methods
// that our service need in order to operate
type Store interface {
	GetAppointment(context.Context, int) (Appointment, error)
	CreateAppointment(context.Context, CreateNewAppointmentRequest) (Appointment, error)
	UpdateAppointment(context.Context, UpdateNewAppointmentRequest) (Appointment, error)
	DeleteAppointment(context.Context, int) error
}

//type AppointmentStore interface {
//	GetAppointment(context.Context, string) (Appointment, error)
//	CreateAppointment(context.Context, CreateNewAppointmentRequest) (Appointment, error)
//}

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

// Appointment - a representation of the appointment
// structure for our service
type Appointment struct {
	ID            int64
	AppointmentID uuid.UUID
	Date          time.Time
	Comment       string
	PatientID     int64
	DoctorID      int64
}

//// Service - is the struct on which all our
//// logic will be built on top of
//type Service struct {
//	Store internal.Store
//}

// NewService - returns a pointer to a new
// service
//func NewService(store internal.Store) *Service {
//	return &Service{
//		Store: store,
//	}
//}

func (s *Service) GetAppointment(ctx context.Context, id int) (Appointment, error) {
	fmt.Println("retrieving an appointment")
	appointment, err := s.Store.GetAppointment(ctx, id)
	if err != nil {
		return serviceErrorHandler(Appointment{}, err, ErrFetchingAppointment)
	}
	return appointment, nil

}

type CreateNewAppointmentRequest struct {
	Date    time.Time
	Comment string
}

func (s *Service) CreateAppointment(ctx context.Context, req CreateNewAppointmentRequest) (Appointment, error) {
	fmt.Println("creating an appointment")
	appointment, err := s.Store.CreateAppointment(ctx, req)
	if err != nil {
		return serviceErrorHandler(Appointment{}, err, ErrCreatingAppointment)
	}
	return appointment, nil
}

func serviceErrorHandler[K any](data K, implError error, serviceError error) (K, error) {

	// use this error to know implementation errors
	// from logs
	fmt.Println(implError)
	// return our own custom errors from service
	// layer to transport layer to guard
	// our implementation details from being
	// exposed to potentially failed client
	// calls
	return data, serviceError

}

type UpdateNewAppointmentRequest struct {
	ID      int64
	Date    time.Time
	Comment string
}

func (s *Service) UpdateAppointment(ctx context.Context, req UpdateNewAppointmentRequest) (Appointment, error) {
	fmt.Println("updating an appointment")
	appointment, err := s.Store.UpdateAppointment(ctx, req)
	if err != nil {
		return serviceErrorHandler(Appointment{}, err, ErrUpdatingAppointment)
	}
	return appointment, nil
}

func (s *Service) DeleteAppointment(ctx context.Context, id int) error {
	fmt.Println("deleting an appointment")
	err := s.Store.DeleteAppointment(ctx, id)
	if err != nil {

		return ErrDeletingAppointment
	}
	return nil
}
