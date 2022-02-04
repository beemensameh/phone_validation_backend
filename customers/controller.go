package customers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CustomerControllerProvider interface {
	IndexController(c echo.Context) error
}

type CustomerController struct {
	customerService CustomerServiceProvider
}

func NewCustomerController(customerService CustomerServiceProvider) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

func (c CustomerController) IndexController(ctx echo.Context) error {
	country_code := ctx.QueryParam("country_code")
	state := ctx.QueryParam("state")

	customers, err := c.customerService.GetCustomers(country_code)
	if err != nil {
		logrus.Error(err)
		return ctx.JSON(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, IndexResponse(customers, state))
}
