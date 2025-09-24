// generate/generate.go
//go:generate go run .

// This program generates the structs from JSON
package main

import (
	"log"
)

// CountriesWithCurrencies is a shim for parsing
type countriesWithCurrencies []*alternateCountryData

// currencyObject is a country object that contains the currency
type alternateCountryData struct {
	Capital       string `json:"capital"`
	ContinentName string `json:"continentName"`
	CountryCode   string `json:"countryCode"`
	CountryName   string `json:"countryName"`
	CurrencyCode  string `json:"currencyCode"`
	Population    string `json:"population"`
}

// Country mirrors the main package struct for code generation
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

// mapEntry is a helper struct to hold the key and index of a capital in the sorted list
type mapEntry struct {
	Key   string
	Index int
}

// CountryList is a slice of Country pointers
type CountryList []*Country

// main is the entry point for the code generation tool.
// It loads country and currency data from embedded JSON sources,
// merges the datasets to enrich country information with currency and capital details,
// and then generates a Go source file (`countries_data.go`) containing the combined data as Go structs.
// The generated file is formatted and ready for use in the main package.
// This process ensures that the country data remains up to date and consistent with the source JSON.
func main() {
	// Create production implementations
	dataLoader := &EmbeddedDataLoader{}
	fileWriter := &OSFileWriter{}
	templateProvider := &DefaultTemplateProvider{}

	// Configure the generator
	config := GeneratorConfig{
		DataLoader:       dataLoader,
		FileWriter:       fileWriter,
		TemplateProvider: templateProvider,
		OutputPath:       "../countries_data.go",
		RepoURL:          "https://github.com/mrz1836/go-countries",
	}

	// Create and run the generator
	generator := NewGenerator(config)
	if err := generator.Generate(); err != nil {
		log.Fatalf("Failed to generate countries data: %v", err)
	}
}
