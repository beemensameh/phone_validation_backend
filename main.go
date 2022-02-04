package main

import (
	"github.com/beemensameh/phone_validation/configuration"
	"github.com/beemensameh/phone_validation/countries"
	"github.com/beemensameh/phone_validation/customers"
	"github.com/beemensameh/phone_validation/db"
	"github.com/beemensameh/phone_validation/server"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configuration.Init("development", "yaml", "./configuration")
	db.Init()

	// Wiring
	server.InitWire()

	customerController := server.GetWire("customerController").(customers.CustomerControllerProvider)
	countryController := server.GetWire("countryService").(countries.CountryControllerProvider)

	// Routes
	e.GET("api/v1/countries-code", countryController.IndexController)
	e.GET("api/v1/customers", customerController.IndexController)

	e.Logger.Fatal(e.Start(":1234"))
}
