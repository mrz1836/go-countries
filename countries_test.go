package countries

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testCountry       = "united states of america"
	testCountryAlpha2 = "US"
	testCountryAlpha3 = "USA"
	testCountryCode   = "840"
	testCountryISO    = "ISO 3166-2:US"
)

// TestCountries_Loaded tests that the country data is preloaded
func TestCountries_Loaded(t *testing.T) {

	// Make sure all countries are there
	require.NotNil(t, countries)
	assert.Len(t, countries, 249)

	// Spot check a country
	usa := GetByAlpha2(testCountryAlpha2)
	require.NotNil(t, usa)

	// All fields (USA only)
	assert.Equal(t, testCountryAlpha2, usa.Alpha2)
	assert.Equal(t, testCountryAlpha3, usa.Alpha3)
	assert.Equal(t, "Washington", usa.Capital)
	assert.Equal(t, "North America", usa.ContinentName)
	assert.Equal(t, testCountryCode, usa.CountryCode)
	assert.Equal(t, "USD", usa.CurrencyCode)
	assert.Empty(t, usa.IntermediateRegion)
	assert.Empty(t, usa.IntermediateRegionCode)
	assert.Equal(t, testCountryISO, usa.ISO31662)
	assert.Equal(t, "United States of America", usa.Name)
	assert.Equal(t, "Americas", usa.Region)
	assert.Equal(t, "019", usa.RegionCode)
	assert.Equal(t, "Northern America", usa.SubRegion)
	assert.Equal(t, "021", usa.SubRegionCode)
}

// TestGetByName_VariousFormats tests GetByName with different input cases
func TestGetByName_VariousFormats(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		expectedName string
		expectNil    bool
	}{
		{name: "Lower to capital", input: testCountry, expectedName: "United States of America"},
		{name: "Format to lower, mixed caps", input: "AfghanistaN", expectedName: "Afghanistan"},
		{name: "Symbol detection", input: "Åland Islands", expectedName: "Åland Islands"},
		{name: "All caps", input: "ALBANIA", expectedName: "Albania"},
		{name: "no country found", input: "no-country", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByName(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expectedName, country.Name)
		})
	}
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
	// Output:&{Alpha2:US Alpha3:USA Capital:Washington ContinentName:North America CountryCode:840 CurrencyCode:USD IntermediateRegion: IntermediateRegionCode: ISO31662:ISO 3166-2:US Name:United States of America Region:Americas RegionCode:019 SubRegion:Northern America SubRegionCode:021}
}

// BenchmarkGetByName benchmarks the method GetByName()
func BenchmarkGetByName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByName(testCountry)
	}
}

// TestGetByAlpha2 tests the GetByAlpha2 function.
func TestGetByAlpha2_VariousFormats(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectNil bool
	}{
		{name: "All caps", input: "AF", expected: "AF"},
		{name: "Lowercase", input: "ax", expected: "AX"},
		{name: "Valid case", input: testCountryAlpha2, expected: testCountryAlpha2},
		{name: "Invalid country", input: "NANA", expectNil: true},
		{name: "Invalid short", input: "N", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByAlpha2(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.Alpha2)
		})
	}
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

// TestGetByAlpha3 tests the GetByAlpha3 function.
func TestGetByAlpha3_VariousFormats(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectNil bool
	}{
		{name: "All caps", input: "AFG", expected: "AFG"},
		{name: "Lowercase", input: "ala", expected: "ALA"},
		{name: "Valid case", input: testCountryAlpha3, expected: testCountryAlpha3},
		{name: "Invalid country", input: "NANA", expectNil: true},
		{name: "Invalid short", input: "N", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByAlpha3(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.Alpha3)
		})
	}
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
		_ = GetByAlpha3(testCountryAlpha3)
	}
}

// TestGetByCountryCode tests the GetByCountryCode function.
func TestGetByCountryCode_ValidInvalid(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectNil bool
	}{
		{name: "Valid code 1", input: testCountryCode, expected: testCountryCode},
		{name: "Valid code 2", input: "248", expected: "248"},
		{name: "Invalid zero", input: "0", expectNil: true},
		{name: "Invalid long", input: "12345", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByCountryCode(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.CountryCode)
		})
	}
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

// TestGetByISO31662 tests the GetByISO31662 function.
func TestGetByISO31662_ValidInvalid(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectNil bool
	}{
		{name: "Valid code", input: testCountryISO, expected: testCountryISO},
		{name: "Invalid zero", input: "0", expectNil: true},
		{name: "Invalid long", input: "12345", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByISO31662(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.ISO31662)
		})
	}
}

// ExampleGetByISO31662 is an example of GetByISO31662()
func ExampleGetByISO31662() {
	country := GetByISO31662(testCountryISO)
	fmt.Printf(
		"country: %s alpha2: %s alpha3: %s code: %s",
		country.Name, country.Alpha2, country.Alpha3, country.CountryCode,
	)
	// Output:country: United States of America alpha2: US alpha3: USA code: 840
}

// BenchmarkGetByISO31662 benchmarks the method GetByISO31662()
func BenchmarkGetByISO31662(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByISO31662(testCountryISO)
	}
}

// TestGetAll tests the GetAll function.
func TestGetAll_Basic(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "valid countries",
			test: func(t *testing.T) {
				c := GetAll()
				require.NotNil(t, c)
				assert.Len(t, c, 249)
			},
		},
		{
			name: "returns a copy",
			test: func(t *testing.T) {
				c1 := GetAll()
				require.NotNil(t, c1)

				c2 := GetAll()
				require.NotNil(t, c2)

				// Modify the first slice
				c1[0] = &Country{Alpha2: "XX"}
				c1 = append(c1, &Country{Alpha2: "YY"})

				// Verify the second slice remains unchanged
				assert.NotSame(t, c1[0], c2[0])
				assert.Len(t, c1, 250)
				assert.Len(t, c2, 249)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

// ExampleGetAll is an example of GetAll()
func ExampleGetAll() {
	all := GetAll()
	fmt.Printf("countries found: %d", len(all))
	// Output:countries found: 249
}

// BenchmarkGetAll benchmarks the method GetAll()
func BenchmarkGetAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetAll()
	}
}
