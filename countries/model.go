package countries

import "github.com/beemensameh/phone_validation/customers"

func GetCountriesCodes() []string {
	var countriesCodes []string
	for _, phoneRegex := range customers.GetPhoneRegexes() {
		countriesCodes = append(countriesCodes, phoneRegex.CountryCode)
	}

	return countriesCodes
}
