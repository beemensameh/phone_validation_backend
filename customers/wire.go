//go:build wireinject
// +build wireinject

package customers

import "github.com/google/wire"

func InitializeCustomerController(customerService CustomerServiceProvider) *CustomerController {
	wire.Build(NewCustomerController)
	return &CustomerController{}
}

func InitializeCustomerService() *CustomerService {
	wire.Build(NewCustomerService)
	return &CustomerService{}
}
