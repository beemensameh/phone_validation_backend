package countries

type allCountryCodesResponse struct {
	CountriesCode []string `json:"countries_code,omitempty"`
}

func GetCountryCodeResponse() allCountryCodesResponse {
	return allCountryCodesResponse{
		CountriesCode: GetCountriesCodes(),
	}
}
