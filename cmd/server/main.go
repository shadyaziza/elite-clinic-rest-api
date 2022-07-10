package main

import (
	"fmt"
	transportHttp "github.com/shadyaziza/elite-clinic-rest-api/internal/transport/http"
	"net/http"
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
	handler := transportHttp.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		println("Failed to set up server")
		return err
	}
	return nil
}
func main() {
	fmt.Println("rest")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println(err)
	}
}
