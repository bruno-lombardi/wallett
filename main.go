package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
	"wallett/data"
	"wallett/infra/persistence/db/sqlite"
	"wallett/main/handlers"
	"wallett/presentation/helpers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	data := data.NewWSD("wsd.dat")
	sqlite.Connect("gorm.db")

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
	userHandlers := handlers.NewUserHandlers(sqlite.GetDB())
	userHandlers.SetupRoutes(api)

	walletHandlers := handlers.NewWalletHandlers(data)
	walletHandlers.SetupHandlers(api)

	go func() {
		if err := e.Start(":3333"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
