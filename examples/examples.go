// Package main provides examples of how to use the go-countries package.
// Each public function is demonstrated using mostly the USA, Canada and Mexico.
package main

import (
	"log"

	"github.com/mrz1836/go-countries"
)

func main() {
	// Lookup by country name (United States)
	usa := countries.GetByName("United States of America")
	log.Printf("USA alpha-2: %s", usa.Alpha2)

	// Lookup by alpha-2 code (Canada)
	canada := countries.GetByAlpha2(countries.Alpha2CA)
	log.Printf("Canada name: %s", canada.Name)

	// Lookup by alpha-3 code (Mexico)
	mexico := countries.GetByAlpha3(countries.Alpha3MEX)
	log.Printf("Mexico capital: %s", mexico.Capital)

	// Lookup by numeric country code (US)
	byCode := countries.GetByCountryCode("840")
	log.Printf("Country for code 840: %s", byCode.Name)

	// Lookup by ISO 3166-2 code (Canada)
	byISO := countries.GetByISO31662("ISO 3166-2:CA")
	log.Printf("Country for ISO 3166-2 'ISO 3166-2:CA': %s", byISO.Name)

	// Get the full list of countries
	all := countries.GetAll()
	log.Printf("Total countries available: %d", len(all))
}
