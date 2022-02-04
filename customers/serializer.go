package customers

type allCustomersResponse struct {
	Customers []getCustomersResponse `json:"customers,omitempty"`
}

type getCustomersResponse struct {
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	State       bool   `json:"state"`
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
}

func IndexResponse(custs []Customer, state string) allCustomersResponse {
	var customersResponse []getCustomersResponse

	for _, customer := range custs {
		countryCode := customer.GetPhoneCode()
		for _, country := range GetPhoneRegexes() {
			isValidCountry := country.IsValidate(customer.Phone)
			if countryCode == country.CountryCode && (state == "" || (state == "valid" && isValidCountry) || (state == "invalid" && !isValidCountry)) {
				customersResponse = append(customersResponse, getCustomersResponse{
					Name:        customer.Name,
					Phone:       customer.Phone,
					State:       isValidCountry,
					CountryName: country.CountryName,
					CountryCode: country.CountryCode,
				})
				break
			}
		}
	}

	return allCustomersResponse{
		Customers: customersResponse,
	}
}
