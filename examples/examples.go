// Package main provides an example of how to use the go-countries package.
package main

import (
	"log"

	"github.com/mrz1836/go-countries"
)

func main() {

	// Get by name (show ISO 3166-2)
	country := countries.GetByName("United States of America")
	log.Println("Found: " + country.Alpha2)

	// Get by ISO 3166-2  (show name)
	country = countries.GetByAlpha2("US")
	log.Println("Found: " + country.Name)
}
