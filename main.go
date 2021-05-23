package main

import (
	"fmt"
	"net/http"
	"wallett/controllers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type GoPlaygroundValidator struct {
	validator *validator.Validate
}

func (cv *GoPlaygroundValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func main() {
	// wsd := initWSD()
	e := echo.New()
	e.Validator = &GoPlaygroundValidator{validator: validator.New()}
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
	api.GET("/users", controllers.HandleListUsers)
	api.GET("/users/:id", controllers.HandleGetUserByID)
	api.POST("/users", controllers.CreateUser)
	// api.PUT("/users/:id")
	// api.DELETE("/users/:id")

	// api.GET("/wallets")
	// api.GET("/wallets/:id")
	// api.POST("/wallets")
	// api.PUT("/wallets/:id")
	// api.DELETE("/wallets/:id")

	// e.GET("/wallets/:wallet_id/transactions")
	// e.GET("/wallets/:wallet_id/transactions/:trx_id")
	// e.POST("/wallets/:wallet_id/transactions")
	// e.PUT("/wallets/:wallet_id/transactions/:trx_id")
	// e.DELETE("/wallets/:wallet_id/transactions/:trx_id")
	e.Logger.Fatal(e.Start(":3333"))

}
