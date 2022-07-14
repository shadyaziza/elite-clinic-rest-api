package main

import (
	"context"
	"fmt"
	"github.com/shadyaziza/elite-clinic-rest-api/internal"
	"github.com/shadyaziza/elite-clinic-rest-api/internal/appointment"
	database "github.com/shadyaziza/elite-clinic-rest-api/internal/db"
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
	if err := db.Ping(context.Background()); err != nil {
		return err
	}
	fmt.Println("Successfully connected and pinged database")

	if err := db.MigrateDB(); err != nil {
		fmt.Println("Failed to perform migration")
		return err
	}

	service, err := appointment.NewService(internal.NewStore(db.Client)).GetAppointment(context.Background(), "7e3577e8-0d70-4389-b937-e3abc7a3b0ee")

	if err != nil {
		fmt.Println(fmt.Errorf("error setting up service %w", err))
	}
	fmt.Println(service)
	//handler := transportHttp.NewHandler()
	//handler.SetupRoutes()
	//
	//if err := http.ListenAndServe(":8080", handler.Router); err != nil {
	//	println("Failed to set up server")
	//	return err
	//}
	return nil
}
func main() {
	fmt.Println("rest")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
