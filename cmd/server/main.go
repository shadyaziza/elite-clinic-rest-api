package main

import (
	"fmt"
	"github.com/shadyaziza/elite-clinic-rest-api/internal/appointment"
	database "github.com/shadyaziza/elite-clinic-rest-api/internal/db"
	transportHttp "github.com/shadyaziza/elite-clinic-rest-api/internal/transport/http"
)

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - is going to be responsible for
// the instantiation and startup of our
// go application
func (app *App) Run() error {
	fmt.Println("starting our api")

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Failed to connect to the database")
		return err
	}

	fmt.Println("Successfully connected and pinged database")

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to perform migration")
		return err
	}

	appointmentService := appointment.NewService(db)

	httpHandler := transportHttp.NewHandler(appointmentService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}
	fmt.Println("Server started...")
	return nil
}
func main() {
	fmt.Println("rest")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
