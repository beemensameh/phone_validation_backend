//go:build wireinject
// +build wireinject

package countries

import "github.com/google/wire"

func InitializeCountryController() *CountryController {
	wire.Build(NewCountryController)
	return &CountryController{}
}
