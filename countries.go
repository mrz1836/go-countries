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
	Alpha2                 string `json:"alpha-2"`                  // ISO 3166-1 alpha-2 code
	Alpha3                 string `json:"alpha-3"`                  // ISO 3166-1 alpha-3 code
	Capital                string `json:"capital"`                  // Capital city of the country
	ContinentName          string `json:"continent_name"`           // The Name of the continent the country is located in
	CountryCode            string `json:"country-code"`             // Numeric ISO 3166-1 code
	CurrencyCode           string `json:"currency_code"`            // ISO 4217 currency code
	ISO31662               string `json:"iso_3166-2"`               // ISO 3166-2 code for subdivisions
	IntermediateRegion     string `json:"intermediate-region"`      // Name of the intermediate region (if applicable)
	IntermediateRegionCode string `json:"intermediate-region-code"` // Code for the intermediate region (if applicable)
	Name                   string `json:"name"`                     // Name of the country
	Region                 string `json:"region"`                   // Name of the region the country is located in
	RegionCode             string `json:"region-code"`              // Code for the region (e.g., continent code)
	SubRegion              string `json:"sub-region"`               // The Name of the subregion the country is located in
	SubRegionCode          string `json:"sub-region-code"`          // Code for the sub-region (e.g., continent sub-region code)
}

// GetByName retrieves a Country by its name in a case-insensitive search.
//
// This function performs the following steps:
// - Converts the input name to lowercase for normalization
// - Performs a constant-time map lookup using the normalized name
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
// - Lookup uses a prebuilt map for constant-time access
// - The result references the internal Country struct without copying
func GetByName(name string) *Country {
	return byName[strings.ToLower(name)]
}

// GetByAlpha2 retrieves a Country by its alpha-2 code in a case-insensitive search.
//
// This function performs the following steps:
// - Normalizes the provided code to uppercase
// - Performs a constant-time map lookup using the normalized code
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
// - Lookup uses a map for constant-time retrieval
// - Returned pointer references package-level data without copying
func GetByAlpha2(alpha2 string) *Country {
	return byAlpha2[strings.ToUpper(alpha2)]
}

// GetByAlpha3 retrieves a Country using its alpha-3 code in a case-insensitive search.
//
// This function performs the following steps:
// - Converts the incoming code to uppercase
// - Performs a constant-time map lookup using the uppercase code
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
// - Lookup uses a map for constant-time retrieval
// - Returned pointer references global data and should not be mutated
func GetByAlpha3(alpha3 string) *Country {
	return byAlpha3[strings.ToUpper(alpha3)]
}

// GetByCountryCode looks up a Country by its numeric code using a case-sensitive comparison.
//
// This function performs the following steps:
// - Performs a constant-time map lookup using the numeric code
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
// - Lookup uses a map for constant-time retrieval
// - The returned Country pointer references package data directly
func GetByCountryCode(code string) *Country {
	return byCode[code]
}

// GetByCapital retrieves a Country by its capital city in a case-insensitive search.
//
// This function performs the following steps:
// - Converts the provided capital name to lowercase
// - Performs a constant-time map lookup using the normalized name
//
// - Returns the Country pointer when a match exists
// - Returns nil if the capital is not found
//
// Parameters:
// - capital: the capital city used for the lookup
//
// Returns:
// - Pointer to the Country struct, or nil when no match is found
//
// Side Effects:
// - None
//
// Notes:
// - Lookup uses a prebuilt map for constant-time retrieval
// - Returned pointer references package data directly
func GetByCapital(capital string) *Country {
	return byCapital[strings.ToLower(capital)]
}

// GetByISO31662 locates a Country by its ISO 3166-2 code using a case-insensitive match.
//
// This function performs the following steps:
// - Converts the provided code to uppercase for uniform comparison
// - Performs a constant-time map lookup using the normalized code
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
// - Lookup uses a map for constant-time retrieval
// - Returned pointer references global data and should be treated as read-only
func GetByISO31662(iso string) *Country {
	return byISO31662[strings.ToUpper(iso)]
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
	return append(CountryList(nil), countries...)
}
