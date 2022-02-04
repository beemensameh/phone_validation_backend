package customers

import "github.com/beemensameh/phone_validation/network"

var GetAllCustomersCode = 1000

var GetAllCustomersError = network.InternalServerError(GetAllCustomersCode, "Error while getting all customers")
