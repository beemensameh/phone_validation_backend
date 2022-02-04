package countries

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CountryControllerProvider interface {
	IndexController(c echo.Context) error
}

type CountryController struct{}

func NewCountryController() *CountryController {
	return &CountryController{}
}

func (c CountryController) IndexController(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, GetCountryCodeResponse())
}
