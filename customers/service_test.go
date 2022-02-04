package customers_test

import (
	"testing"

	"github.com/beemensameh/phone_validation/configuration"
	"github.com/beemensameh/phone_validation/customers"
	"github.com/beemensameh/phone_validation/db"
	"github.com/beemensameh/phone_validation/server"
	"github.com/stretchr/testify/require"
)

func TestGetCustomers(t *testing.T) {
	// Wiring
	server.InitWire()

	customerController := server.GetWire("customerService").(customers.CustomerServiceProvider)

	configuration.Init("testing", "yaml", "./configuration")
	db.Init()

	testCases := map[string]struct {
		countryCode string
	}{
		"Return list of customers when don't send country code": {
			countryCode: "",
		},
		"Return list of customers depend on the country code": {
			countryCode: "212",
		},
	}
	for desc, tc := range testCases {
		tc := tc
		t.Run(desc, func(t *testing.T) {
			t.Parallel()

			actualResponse, err := customerController.GetCustomers(tc.countryCode)
			require.NoError(t, err)
			require.NotEmpty(t, actualResponse)

			if tc.countryCode != "" {
				for _, customer := range actualResponse {
					actualCountryCode := customer.GetPhoneCode()
					require.Equal(t, tc.countryCode, actualCountryCode)
				}
			}
		})
	}
}
