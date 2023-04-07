package main

import (
	"drivers-location-h3/src/config"
	"drivers-location-h3/src/info"
	"fmt"
	"log"
	"net/http"
	"time"

	driver_server "drivers-location-h3/src/driver/server"
	order_server "drivers-location-h3/src/order/server"

	"github.com/labstack/echo"
)

func main() {
	info.PrintCountTable()

	server := setupServer()

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST). See
	// documentation below for more options.

	fmt.Printf("Listening on %v\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func setupServer() *http.Server {
	e := echo.New()

	e.GET("/ping", pingHandler)
	e.GET("/env", envHandler)

	driver_server.NewDriverHandler(e)
	order_server.NewOrderHandler(e)

	e.Static("/", "../static")

	return &http.Server{
		Addr:           ":8000",
		Handler:        e,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func pingHandler(c echo.Context) error {
	c.String(http.StatusOK, "pong")

	return nil
}

func envHandler(c echo.Context) error {
	c.JSON(http.StatusOK, config.GetConfigs())

	return nil
}
