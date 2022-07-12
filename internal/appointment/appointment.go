package appointment

import (
	"context"
	"fmt"
	"time"
)

// Appointment - a representation of the appointment
// structure for our service
type Appointment struct {
	id        time.Time
	comment   string
	patientID string
	doctorID  string
}

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
		fmt.Println(err)
		return Appointment{}, err
	}
	return appointment, nil

}
