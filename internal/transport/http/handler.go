package http

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Handler - stores pointer to our service
type Handler struct {
	Router             *mux.Router
	AppointmentService AppointmentService
	Server             *http.Server
}

type AppointmentService interface{}

// NewHandler - returns a pointer to a Handler
func NewHandler(appService AppointmentService) *Handler {
	h := &Handler{
		AppointmentService: appService,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()
	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}
	return h
}

// mapRoutes - sets up all the routes for our application
func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hellow Shady")
	})

}

func (h *Handler) Serve() error {
	if err := h.Server.ListenAndServe(); err != nil {
		return err
	}
	return nil
}
