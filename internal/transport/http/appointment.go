package http

import (
	"context"
	"encoding/json"
	"github.com/shadyaziza/elite-clinic-rest-api/internal/appointment"
	"log"
	"net/http"
)

type AppointmentService interface {
	PostAppointment(ctx context.Context, appointment appointment.CreateNewAppointmentRequest) (appointment.Appointment, error)
	GetAppointment(ctx context.Context, ID int) (appointment.Appointment, error)
	UpdateAppointment(ctx context.Context, newAppointment appointment.UpdateNewAppointmentRequest) (appointment.Appointment, error)
	DeleteAppointment(ctx context.Context, ID int) error
}

func (h *Handler) PostAppointment(w http.ResponseWriter, r *http.Request) {
	var result appointment.Appointment
	
	var appointment appointment.CreateNewAppointmentRequest

	if err := json.NewDecoder(r.Body).Decode(&appointment); err != nil {
		log.Print(err)
		return
	}
	result, err := h.AppointmentService.PostAppointment(r.Context(), appointment)
	if err != nil {
		log.Print(err)
		return
	}
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAppointment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) UpdateAppointment(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) DeleteAppointment(w http.ResponseWriter, r *http.Request) {

}
