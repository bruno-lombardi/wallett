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

	e.GET("/users", controllers.HandleListUsers)
	e.GET("/users/:id", controllers.HandleGetUserByID)
	// e.POST("/users")
	// e.PUT("/users/:id")
	// e.DELETE("/users/:id")

	// e.GET("/wallets")
	// e.GET("/wallet/:id")
	// e.POST("/wallet")
	// e.PUT("/wallet/:id")
	// e.DELETE("/wallet/:id")

	// e.GET("/wallet/:wallet_id/transactions")
	// e.GET("/wallet/:wallet_id/transactions/:trx_id")
	// e.POST("/wallet/:wallet_id/transaction")
	// e.PUT("/wallet/:wallet_id/transaction/:trx_id")
	// e.DELETE("/wallet/:wallet_id/transaction/:trx_id")
	e.Logger.Fatal(e.Start(":3333"))

}
