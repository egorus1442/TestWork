package app

import (
	"fmt"
	"log"
	"testWork/internal/app/endpoints"
	"testWork/internal/app/midleware"
	"testWork/internal/app/service"

	"github.com/labstack/echo/v4"
)

type App struct {
	e    *endpoints.Endpoint
	s    *service.Service
	echo *echo.Echo
}

func New() (*App, error) {
	a := &App{}

	a.s = service.New()

	a.e = endpoints.New(a.s)

	a.echo = echo.New()

	a.echo.Use(midleware.RoleCheck)

	a.echo.GET("/status", a.e.Status)

	return a, nil
}

func (a App) Run() error {
	fmt.Println("Server is running")

	err := a.echo.Start(":8080")
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
