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

// GetByName retrieves a Country by its name in a case-insensitive search.
//
// This function performs the following steps:
// - Converts the input name to lowercase for normalization
// - Iterates over the in-memory list of countries
//   - Compares each country's lowercase name to the normalized input
//
// - Returns the matching Country pointer when found
// - Returns nil if no matching country exists
//
// Parameters:
// - name: country name used for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - The search performs a linear scan over the loaded country list
// - The result references the internal Country struct without copying
func GetByName(name string) *Country {
	name = strings.ToLower(name)
	for _, country := range countries {
		if strings.ToLower(country.Name) == name {
			return country
		}
	}
	return nil
}

// GetByAlpha2 retrieves a Country by its alpha-2 code in a case-insensitive search.
//
// This function performs the following steps:
// - Normalizes the provided code to uppercase
// - Scans the in-memory list of countries
//   - Matches the country's alpha-2 field against the normalized code
//
// - Returns the Country pointer on success
// - Returns nil if no match is located
//
// Parameters:
// - alpha2: two-letter ISO 3166 code used for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - Search is linear over the preloaded slice of countries
// - Returned pointer references package-level data without copying
func GetByAlpha2(alpha2 string) *Country {
	alpha2 = strings.ToUpper(alpha2)
	for _, country := range countries {
		if country.Alpha2 == alpha2 {
			return country
		}
	}
	return nil
}

// GetByAlpha3 retrieves a Country using its alpha-3 code in a case-insensitive search.
//
// This function performs the following steps:
// - Converts the incoming code to uppercase
// - Iterates through the loaded list of countries
//   - Compares each country's alpha-3 code with the uppercase input
//
// - Returns the Country pointer if found
// - Returns nil when the code is not present
//
// Parameters:
// - alpha3: three-letter ISO 3166 code used for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - Performs a linear scan over the in-memory country slice
// - Returned pointer references global data and should not be mutated
func GetByAlpha3(alpha3 string) *Country {
	alpha3 = strings.ToUpper(alpha3)
	for _, country := range countries {
		if country.Alpha3 == alpha3 {
			return country
		}
	}
	return nil
}

// GetByCountryCode looks up a Country by its numeric code using a case-sensitive comparison.
//
// This function performs the following steps:
// - Iterates over the loaded list of countries
//   - Compares each country's numeric code with the input value
//
// - Returns the Country pointer when a match is found
// - Returns nil if the code does not exist in the list
//
// Parameters:
// - code: numeric ISO 3166 code provided for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - The search uses a linear scan over the package's country slice
// - The returned Country pointer references package data directly
func GetByCountryCode(code string) *Country {
	for _, country := range countries {
		if country.CountryCode == code {
			return country
		}
	}
	return nil
}

// GetByISO31662 locates a Country by its ISO 3166-2 code using a case-insensitive match.
//
// This function performs the following steps:
// - Converts the provided code to uppercase for uniform comparison
// - Scans the internal list of countries
//   - Compares each country's ISO31662 field with the normalized code
//
// - Returns the Country pointer when a match exists
// - Returns nil if no matching code is found
//
// Parameters:
// - iso: the ISO 3166-2 code used for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - Search iterates sequentially over the slice of countries
// - Returned pointer references global data and should be treated as read-only
func GetByISO31662(iso string) *Country {
	iso = strings.ToUpper(iso)
	for _, country := range countries {
		if country.ISO31662 == iso {
			return country
		}
	}
	return nil
}

// GetAll provides a copy of every Country currently loaded.
//
// This function performs the following steps:
// - Creates a new slice with the same length as the internal country slice
// - Appends all existing Country pointers into that slice
// - Returns the new slice to the caller
//
// Parameters:
// - None
//
// Returns:
// - CountryList containing pointers to all Country structs
//
// Side Effects:
// - None
//
// Notes:
// - The Country pointers reference global data but the returned slice is a copy
// - Modifying the slice does not alter the package-level slice
func GetAll() CountryList {
	clone := append(CountryList(nil), countries...)
	return clone
}
