package countries

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testCountry       = "united states of america"
	testCountryAlpha2 = "US"
	testCountryAlpha3 = "USA"
	testCountryCode   = "840"
	testCountryIso    = "ISO 3166-2:US"
)

// TestCountries will test our pre-loaded countries
func TestCountries(t *testing.T) {
	assert.Equal(t, 249, len(countries))
}

// TestGetByName will test the method GetByName()
func TestGetByName(t *testing.T) {
	t.Parallel()

	t.Run("Lower to capital", func(t *testing.T) {
		country := GetByName(testCountry)
		assert.NotNil(t, country)
		assert.Equal(t, "United States of America", country.Name)
	})

	t.Run("Format to lower, mixed caps", func(t *testing.T) {
		country := GetByName("AfghanistaN")
		assert.NotNil(t, country)
		assert.Equal(t, "Afghanistan", country.Name)
	})

	t.Run("Symbol detection", func(t *testing.T) {
		country := GetByName("Åland Islands")
		assert.NotNil(t, country)
		assert.Equal(t, "Åland Islands", country.Name)
	})

	t.Run("All caps", func(t *testing.T) {
		country := GetByName("ALBANIA")
		assert.NotNil(t, country)
		assert.Equal(t, "Albania", country.Name)
	})

	t.Run("no country found", func(t *testing.T) {
		country := GetByName("no-country")
		assert.Nil(t, country)
	})
}

// ExampleGetByName is an example of GetByName()
func ExampleGetByName() {
	country := GetByName(testCountry)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// ExampleGetByName_ShowAll
func ExampleGetByName_showAll() {
	country := GetByName(testCountry)
	fmt.Printf("%+v\n", country)
	// Output:&{Alpha2:US Alpha3:USA CountryCode:840 IntermediateRegion: IntermediateRegionCode: ISO31662:ISO 3166-2:US Name:United States of America Region:Americas RegionCode:019 SubRegion:Northern America SubRegionCode:021}
}

// BenchmarkGetByName benchmarks the method GetByName()
func BenchmarkGetByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByName(testCountry)
	}
}

// TestGetByAlpha2 will test the method GetByAlpha2()
func TestGetByAlpha2(t *testing.T) {
	t.Parallel()

	t.Run("All caps", func(t *testing.T) {
		country := GetByAlpha2("AF")
		assert.NotNil(t, country)
		assert.Equal(t, "AF", country.Alpha2)
	})

	t.Run("Lowercase", func(t *testing.T) {
		country := GetByAlpha2("ax")
		assert.NotNil(t, country)
		assert.Equal(t, "AX", country.Alpha2)
	})

	t.Run("Valid case", func(t *testing.T) {
		country := GetByAlpha2(testCountryAlpha2)
		assert.NotNil(t, country)
		assert.Equal(t, testCountryAlpha2, country.Alpha2)
	})

	t.Run("Invalid country", func(t *testing.T) {
		country := GetByAlpha2("NANA")
		assert.Nil(t, country)

		country = GetByAlpha2("N")
		assert.Nil(t, country)
	})
}

// ExampleGetByAlpha2 is an example of GetByAlpha2()
func ExampleGetByAlpha2() {
	country := GetByAlpha2(testCountryAlpha2)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// BenchmarkGetByAlpha2 benchmarks the method GetByAlpha2()
func BenchmarkGetByAlpha2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByAlpha2(testCountryAlpha2)
	}
}

// TestGetByAlpha3 will test the method GetByAlpha3()
func TestGetByAlpha3(t *testing.T) {
	t.Parallel()

	t.Run("All caps", func(t *testing.T) {
		country := GetByAlpha3("AFG")
		assert.NotNil(t, country)
		assert.Equal(t, "AFG", country.Alpha3)
	})

	t.Run("Lowercase", func(t *testing.T) {
		country := GetByAlpha3("ala")
		assert.NotNil(t, country)
		assert.Equal(t, "ALA", country.Alpha3)
	})

	t.Run("Valid case", func(t *testing.T) {
		country := GetByAlpha3(testCountryAlpha3)
		assert.NotNil(t, country)
		assert.Equal(t, testCountryAlpha3, country.Alpha3)
	})

	t.Run("Invalid country", func(t *testing.T) {
		country := GetByAlpha3("NANA")
		assert.Nil(t, country)

		country = GetByAlpha3("N")
		assert.Nil(t, country)
	})
}

// ExampleGetByAlpha3 is an example of GetByAlpha3()
func ExampleGetByAlpha3() {
	country := GetByAlpha3(testCountryAlpha3)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// BenchmarkGetByAlpha3 benchmarks the method GetByAlpha3()
func BenchmarkGetByAlpha3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByAlpha3(testCountryAlpha2)
	}
}

// TestGetByCountryCode will test the method GetByCountryCode()
func TestGetByCountryCode(t *testing.T) {
	t.Parallel()

	t.Run("Valid codes", func(t *testing.T) {
		country := GetByCountryCode(testCountryCode)
		assert.NotNil(t, country)
		assert.Equal(t, testCountryCode, country.CountryCode)

		country = GetByCountryCode("248")
		assert.NotNil(t, country)
		assert.Equal(t, "248", country.CountryCode)
	})

	t.Run("Invalid codes", func(t *testing.T) {
		country := GetByCountryCode("0")
		assert.Nil(t, country)

		country = GetByCountryCode("12345")
		assert.Nil(t, country)
	})
}

// ExampleGetByCountryCode is an example of GetByCountryCode()
func ExampleGetByCountryCode() {
	country := GetByCountryCode(testCountryCode)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// BenchmarkGetByCountryCode benchmarks the method GetByCountryCode()
func BenchmarkGetByCountryCode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByCountryCode(testCountryCode)
	}
}

// TestGetByISO31662 will test the method GetByISO31662()
func TestGetByISO31662(t *testing.T) {
	t.Parallel()

	t.Run("Valid codes", func(t *testing.T) {
		country := GetByISO31662(testCountryIso)
		assert.NotNil(t, country)
		assert.Equal(t, testCountryIso, country.ISO31662)
	})

	t.Run("Invalid codes", func(t *testing.T) {
		country := GetByISO31662("0")
		assert.Nil(t, country)

		country = GetByISO31662("12345")
		assert.Nil(t, country)
	})
}

// ExampleGetByISO31662 is an example of GetByISO31662()
func ExampleGetByISO31662() {
	country := GetByISO31662(testCountryIso)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// BenchmarkGetByISO31662 benchmarks the method GetByISO31662()
func BenchmarkGetByISO31662(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByISO31662(testCountryIso)
	}
}
