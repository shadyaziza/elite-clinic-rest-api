package main

import "fmt"

// App - the struct which contains things like pointers
// to database connections
type App struct {
}

// Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting up api")
	return nil
}
func main() {
	fmt.Println("rest")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up api")
	}
}
