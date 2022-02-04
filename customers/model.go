package customers

import (
	"bytes"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/spf13/viper"
)

type Customer struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

func (c *Customer) TableName() string {
	return "customer"
}

func (c *Customer) GetPhoneCode() string {
	codeBrackets := regexp.MustCompile(`(?m)[\(\)]`)
	phoneNumeber := codeBrackets.ReplaceAllString(c.Phone, "")
	return strings.Split(phoneNumeber, " ")[0]
}

type PhoneRegex struct {
	CountryName string `json:"country_name"`
	CountryCode string `json:"country_code"`
	Regex       string `json:"regex"`
}

func (c *PhoneRegex) IsValidate(phone string) bool {
	reg := regexp.MustCompile(c.Regex)
	return reg.MatchString(phone)
}

func transcode(in interface{}, out *PhoneRegex) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPhoneRegexes() []PhoneRegex {
	var allPhoneRegexes []PhoneRegex
	phoneRegexes := viper.GetStringMap("phone_regex")
	for _, value := range phoneRegexes {
		phoneRegex := PhoneRegex{}
		transcode(value, &phoneRegex)

		allPhoneRegexes = append(allPhoneRegexes, phoneRegex)
	}

	return allPhoneRegexes
}
