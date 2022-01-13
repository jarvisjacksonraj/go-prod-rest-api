package main

import "fmt"

// App - struct to contain things like pointer to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up our App")
	return nil
}

func main() {
	fmt.Println("Go Rest API")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up our REST API")
		fmt.Println(err)
	}
}
