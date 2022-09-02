package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
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
	// non blocking op within a go routine
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()
	// create the channel
	c := make(chan os.Signal, 1)
	// if receving os.Interrupt block the rest of the func
	signal.Notify(c, os.Interrupt)
	<-c
	// for 15 sec while shutdown the server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 15+time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")
	return nil
}
