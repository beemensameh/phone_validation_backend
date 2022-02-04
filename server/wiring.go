package server

import (
	"github.com/beemensameh/phone_validation/countries"
	"github.com/beemensameh/phone_validation/customers"
)

var wire map[string]interface{} = make(map[string]interface{})

func GetWire(key string) interface{} {
	return wire[key]
}

func InitWire() {
	wire["customerService"] = customers.InitializeCustomerService()
	wire["customerController"] = customers.InitializeCustomerController(wire["customerService"].(customers.CustomerServiceProvider))
	wire["countryService"] = countries.InitializeCountryController()
}
