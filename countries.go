// Package countries is a complete go-ready list of countries in all standardized formats
//
// If you have any suggestions or comments, please feel free to open an issue on
// this GitHub repository!
//
// By @MrZ1836
package countries

import (
	"strings"
)

// Generate the structs from JSON
//go:generate go run generate/generate.go

// CountryList is a list of country structs
type CountryList []*Country

// Country is the single country in the list of countries (ISO-3166)
type Country struct {
	Alpha2                 string `json:"alpha-2"`
	Alpha3                 string `json:"alpha-3"`
	CountryCode            string `json:"country-code"`
	IntermediateRegion     string `json:"intermediate-region"`
	IntermediateRegionCode string `json:"intermediate-region-code"`
	ISO31662               string `json:"iso_3166-2"`
	Name                   string `json:"name"`
	Region                 string `json:"region"`
	RegionCode             string `json:"region-code"`
	SubRegion              string `json:"sub-region"`
	SubRegionCode          string `json:"sub-region-code"`
}

// GetByName will return a country by a given name
// Forces the case to lowercase for comparison
func GetByName(name string) *Country {
	name = strings.ToLower(name)
	for _, country := range countries {
		if strings.ToLower(country.Name) == name {
			return country
		}
	}
	return nil
}

// GetByAlpha2 will return a country by a given alpha-2
// Forces the case to uppercase for comparison
func GetByAlpha2(alpha2 string) *Country {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range countries {
		if country.Alpha2 == alpha2 {
			return country
		}
	}
	return nil
}

// GetByAlpha3 will return a country by a given alpha-3
// Forces the case to uppercase for comparison
func GetByAlpha3(alpha3 string) *Country {
	alpha3 = strings.ToUpper(alpha3)
	for _, country := range countries {
		if country.Alpha3 == alpha3 {
			return country
		}
	}
	return nil
}

// GetByCountryCode will return a country by a given country-code
func GetByCountryCode(code string) *Country {
	for _, country := range countries {
		if country.CountryCode == code {
			return country
		}
	}
	return nil
}

// GetByISO31662 will return a country by a ISO31662 number
// Forces the case to uppercase for comparison
func GetByISO31662(iso string) *Country {
	iso = strings.ToUpper(iso)
	for _, country := range countries {
		if country.ISO31662 == iso {
			return country
		}
	}
	return nil
}
