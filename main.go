package main

import (
	"context"
	"fmt"

	"project-art-museum/config"
	"project-art-museum/util"

	//stdLog "log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {

	//load config if available or set to default
	config := config.GetConfig()

	//initialize database connection based on given config
	dbCon := util.NewDatabaseConnection(config)

	//initiate item repository
	//controllers := modules.RegisterController(dbCon)

	//create echo http
	e := echo.New()
	e.HideBanner = true
	// index route
	e.GET("/", func(c echo.Context) error {
		message := `ok`
		return c.String(http.StatusOK, message)
	})

	//register API path and handler
	//api.RegisterPath(e, controllers)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	//close db
	defer dbCon.CloseConnection()

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
