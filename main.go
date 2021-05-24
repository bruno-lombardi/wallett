package main

import (
	"fmt"
	"wallett/controllers"
	"wallett/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// wsd := initWSD()
	e := echo.New()
	e.Validator = middlewares.NewCustomValidator()
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
	api.POST("/users", controllers.HandleCreateUser)
	api.PUT("/users/:id", controllers.HandleUpdateUser)
	api.DELETE("/users/:id", controllers.HandleDeleteUserByID)

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
