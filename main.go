package main

import (
	"fmt"
	"wallett/data"
	"wallett/main/handlers"
	"wallett/presentation/helpers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	data := data.NewWSD("wsd.dat")
	e := echo.New()
	e.Validator = helpers.NewCustomValidator()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		DisablePrintStack: true,
	}))

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		fmt.Println(c.Path(), c.QueryParams(), err)
		e.DefaultHTTPErrorHandler(err, c)
	}
	api := e.Group("/api/v1")
	userHandlers := handlers.NewUserHandlers(data)
	userHandlers.SetupRoutes(api)

	walletHandlers := handlers.NewWalletHandlers(data)
	walletHandlers.SetupHandlers(api)

	e.Logger.Fatal(e.Start(":3333"))
}
