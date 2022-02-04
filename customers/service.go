package customers

import (
	"fmt"

	"github.com/beemensameh/phone_validation/db"
)

type CustomerServiceProvider interface {
	GetCustomers(country_code string) ([]Customer, error)
}

type CustomerService struct{}

func NewCustomerService() *CustomerService {
	return &CustomerService{}
}

func (s CustomerService) GetCustomers(country_code string) ([]Customer, error) {
	var customers []Customer
	query := db.GetConnection()

	if country_code != "" {
		query = query.Where("phone like ?", fmt.Sprintf("%%%s%%", country_code))
	}

	if err := query.Find(&customers).Error; err != nil {
		return nil, GetAllCustomersError
	}

	return customers, nil
}
