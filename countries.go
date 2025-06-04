// Package countries is a complete go-ready list of countries in all standardized formats.
//
// This package provides functionalities to retrieve country information based on various identifiers
// such as name, alpha-2 code, alpha-3 code, country code, and ISO 3166-2 code. It includes methods
// to get a country by these identifiers and to retrieve the entire list of countries.
//
// The package is designed to be straightforward to use and integrate into Go projects, making it simple to
// work with country data in a standardized way.
//
// If you have any suggestions or comments, please feel free to open an issue in
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
	Capital                string `json:"capital"`
	ContinentName          string `json:"continent_name"`
	CountryCode            string `json:"country-code"`
	CurrencyCode           string `json:"currency_code"`
	IntermediateRegion     string `json:"intermediate-region"`
	IntermediateRegionCode string `json:"intermediate-region-code"`
	ISO31662               string `json:"iso_3166-2"`
	Name                   string `json:"name"`
	Region                 string `json:"region"`
	RegionCode             string `json:"region-code"`
	SubRegion              string `json:"sub-region"`
	SubRegionCode          string `json:"sub-region-code"`
}

// GetByName will return a country by a given name.
// The comparison is case-insensitive, as the input name is converted to lowercase.
// If a country with the specified name is found, it returns a pointer to the Country struct.
// If no country is found, it returns nil.
func GetByName(name string) *Country {
	name = strings.ToLower(name)
	for _, country := range countries {
		if strings.ToLower(country.Name) == name {
			return country
		}
	}
	return nil
}

// GetByAlpha2 will return a country by a given alpha-2 code.
// The comparison is case-insensitive, as the input alpha-2 code is converted to uppercase.
// If a country with the specified alpha-2 code is found, it returns a pointer to the Country struct.
// If no country is found, it returns nil.
func GetByAlpha2(alpha2 string) *Country {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range countries {
		if country.Alpha2 == alpha2 {
			return country
		}
	}
	return nil
}

// GetByAlpha3 will return a country by a given alpha-3 code.
// The comparison is case-insensitive, as the input alpha-3 code is converted to uppercase.
// If a country with the specified alpha-3 code is found, it returns a pointer to the Country struct.
// If no country is found, it returns nil.
func GetByAlpha3(alpha3 string) *Country {
	alpha3 = strings.ToUpper(alpha3)
	for _, country := range countries {
		if country.Alpha3 == alpha3 {
			return country
		}
	}
	return nil
}

// GetByCountryCode will return a country by a given country code.
// The comparison is case-sensitive, as the input country code is used as-is.
// If a country with the specified country code is found, it returns a pointer to the Country struct.
// If no country is found, it returns nil.
func GetByCountryCode(code string) *Country {
	for _, country := range countries {
		if country.CountryCode == code {
			return country
		}
	}
	return nil
}

// GetByISO31662 will return a country by a given ISO 3166-2 code.
// The comparison is case-insensitive, as the input ISO 3166-2 code is converted to uppercase.
// If a country with the specified ISO 3166-2 code is found, it returns a pointer to the Country struct.
// If no country is found, it returns nil.
func GetByISO31662(iso string) *Country {
	iso = strings.ToUpper(iso)
	for _, country := range countries {
		if country.ISO31662 == iso {
			return country
		}
	}
	return nil
}

// GetAll returns a copy of all countries in the list.
// The returned slice has its own backing array, so callers can modify the slice
// without affecting the package data. The country structs themselves are not
// copied.
func GetAll() CountryList {
	clone := append(CountryList(nil), countries...)
	return clone
}
