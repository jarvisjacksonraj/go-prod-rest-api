package main

import (
	"fmt"
	"net/http"

	"github.com/jarvisjacksonraj/go-prod-rest-api/internal/database"

	transportHTTP "github.com/jarvisjacksonraj/go-prod-rest-api/internal/transport/http"
)

// App - struct to contain things like pointer to database connections
type App struct{}

//Run - sets up our application
func (app *App) Run() error {
	fmt.Println("Setting Up our App")

	var err error
	_, err = database.NewDatabase()
	if err != nil {
		return err
	}

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to Setting Up our App")
		return err
	}

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
