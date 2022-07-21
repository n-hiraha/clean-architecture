package main

import (
	"github.com/cosmtrek/air/handler"
	"github.com/cosmtrek/air/injector"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	todoHandler := injector.InjectTodoHandler()
	e := echo.New()
	handler.InitRouting(e, todoHandler)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.Fatal(e.Start(":8081"))
}
