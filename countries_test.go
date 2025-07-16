package countries

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testCountry        = "united states of america"
	testCountryAlpha2  = "US"
	testCountryAlpha3  = "USA"
	testCountryCode    = "840"
	testCountryISO     = "ISO 3166-2:US"
	testCountryCapital = "Washington"
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

// TestLookupMaps_Populated ensures all lookup maps contain every country
func TestLookupMaps_Populated(t *testing.T) {
	require.NotNil(t, byName)
	require.NotNil(t, byAlpha2)
	require.NotNil(t, byAlpha3)
	require.NotNil(t, byCode)
	require.NotNil(t, byISO31662)
	require.NotNil(t, byCapital)

	assert.Len(t, byName, 249)
	assert.Len(t, byAlpha2, 249)
	assert.Len(t, byAlpha3, 249)
	assert.Len(t, byCode, 249)
	assert.Len(t, byISO31662, 249)
	assert.GreaterOrEqual(t, len(byCapital), 240)

	require.Equal(t, GetByAlpha2(testCountryAlpha2), byAlpha2[testCountryAlpha2])
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
	// Output:&{Alpha2:US Alpha3:USA Capital:Washington ContinentName:North America CountryCode:840 CurrencyCode:USD ISO31662:ISO 3166-2:US IntermediateRegion: IntermediateRegionCode: Name:United States of America Region:Americas RegionCode:019 SubRegion:Northern America SubRegionCode:021}
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

// TestGetByCapital tests the GetByCapital function.
func TestGetByCapital_VariousFormats(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectNil bool
	}{
		{name: "Valid case", input: testCountryCapital, expected: testCountryCapital},
		{name: "Lowercase", input: strings.ToLower(testCountryCapital), expected: testCountryCapital},
		{name: "Invalid", input: "NoCity", expectNil: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByCapital(tt.input)
			if tt.expectNil {
				require.Nil(t, country)
				return
			}

			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.Capital)
		})
	}
}

// ExampleGetByCapital is an example of GetByCapital()
func ExampleGetByCapital() {
	country := GetByCapital(testCountryCapital)
	fmt.Printf(
		"country: %s capital: %s",
		country.Name, country.Capital,
	)
	// Output:country: United States of America capital: Washington
}

// BenchmarkGetByCapital benchmarks the method GetByCapital()
func BenchmarkGetByCapital(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetByCapital(testCountryCapital)
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

// TestAlphaCodes_Constants checks all Alpha2 and Alpha3 constants for correctness
func TestAlphaCodes_Constants(t *testing.T) {
	// Alpha-2 codes
	assert.Equal(t, Alpha2AF, "AF")
	assert.Equal(t, Alpha2AX, "AX")
	assert.Equal(t, Alpha2AL, "AL")
	assert.Equal(t, Alpha2DZ, "DZ")
	assert.Equal(t, Alpha2AS, "AS")
	assert.Equal(t, Alpha2AD, "AD")
	assert.Equal(t, Alpha2AO, "AO")
	assert.Equal(t, Alpha2AI, "AI")
	assert.Equal(t, Alpha2AQ, "AQ")
	assert.Equal(t, Alpha2AG, "AG")
	assert.Equal(t, Alpha2AR, "AR")
	assert.Equal(t, Alpha2AM, "AM")
	assert.Equal(t, Alpha2AW, "AW")
	assert.Equal(t, Alpha2AU, "AU")
	assert.Equal(t, Alpha2AT, "AT")
	assert.Equal(t, Alpha2AZ, "AZ")
	assert.Equal(t, Alpha2BS, "BS")
	assert.Equal(t, Alpha2BH, "BH")
	assert.Equal(t, Alpha2BD, "BD")
	assert.Equal(t, Alpha2BB, "BB")
	assert.Equal(t, Alpha2BY, "BY")
	assert.Equal(t, Alpha2BE, "BE")
	assert.Equal(t, Alpha2BZ, "BZ")
	assert.Equal(t, Alpha2BJ, "BJ")
	assert.Equal(t, Alpha2BM, "BM")
	assert.Equal(t, Alpha2BT, "BT")
	assert.Equal(t, Alpha2BO, "BO")
	assert.Equal(t, Alpha2BQ, "BQ")
	assert.Equal(t, Alpha2BA, "BA")
	assert.Equal(t, Alpha2BW, "BW")
	assert.Equal(t, Alpha2BV, "BV")
	assert.Equal(t, Alpha2BR, "BR")
	assert.Equal(t, Alpha2IO, "IO")
	assert.Equal(t, Alpha2BN, "BN")
	assert.Equal(t, Alpha2BG, "BG")
	assert.Equal(t, Alpha2BF, "BF")
	assert.Equal(t, Alpha2BI, "BI")
	assert.Equal(t, Alpha2CV, "CV")
	assert.Equal(t, Alpha2KH, "KH")
	assert.Equal(t, Alpha2CM, "CM")
	assert.Equal(t, Alpha2CA, "CA")
	assert.Equal(t, Alpha2KY, "KY")
	assert.Equal(t, Alpha2CF, "CF")
	assert.Equal(t, Alpha2TD, "TD")
	assert.Equal(t, Alpha2CL, "CL")
	assert.Equal(t, Alpha2CN, "CN")
	assert.Equal(t, Alpha2CX, "CX")
	assert.Equal(t, Alpha2CC, "CC")
	assert.Equal(t, Alpha2CO, "CO")
	assert.Equal(t, Alpha2KM, "KM")
	assert.Equal(t, Alpha2CG, "CG")
	assert.Equal(t, Alpha2CD, "CD")
	assert.Equal(t, Alpha2CK, "CK")
	assert.Equal(t, Alpha2CR, "CR")
	assert.Equal(t, Alpha2CI, "CI")
	assert.Equal(t, Alpha2HR, "HR")
	assert.Equal(t, Alpha2CU, "CU")
	assert.Equal(t, Alpha2CW, "CW")
	assert.Equal(t, Alpha2CY, "CY")
	assert.Equal(t, Alpha2CZ, "CZ")
	assert.Equal(t, Alpha2DK, "DK")
	assert.Equal(t, Alpha2DJ, "DJ")
	assert.Equal(t, Alpha2DM, "DM")
	assert.Equal(t, Alpha2DO, "DO")
	assert.Equal(t, Alpha2EC, "EC")
	assert.Equal(t, Alpha2EG, "EG")
	assert.Equal(t, Alpha2SV, "SV")
	assert.Equal(t, Alpha2GQ, "GQ")
	assert.Equal(t, Alpha2ER, "ER")
	assert.Equal(t, Alpha2EE, "EE")
	assert.Equal(t, Alpha2SZ, "SZ")
	assert.Equal(t, Alpha2ET, "ET")
	assert.Equal(t, Alpha2FK, "FK")
	assert.Equal(t, Alpha2FO, "FO")
	assert.Equal(t, Alpha2FJ, "FJ")
	assert.Equal(t, Alpha2FI, "FI")
	assert.Equal(t, Alpha2FR, "FR")
	assert.Equal(t, Alpha2GF, "GF")
	assert.Equal(t, Alpha2PF, "PF")
	assert.Equal(t, Alpha2TF, "TF")
	assert.Equal(t, Alpha2GA, "GA")
	assert.Equal(t, Alpha2GM, "GM")
	assert.Equal(t, Alpha2GE, "GE")
	assert.Equal(t, Alpha2DE, "DE")
	assert.Equal(t, Alpha2GH, "GH")
	assert.Equal(t, Alpha2GI, "GI")
	assert.Equal(t, Alpha2GR, "GR")
	assert.Equal(t, Alpha2GL, "GL")
	assert.Equal(t, Alpha2GD, "GD")
	assert.Equal(t, Alpha2GP, "GP")
	assert.Equal(t, Alpha2GU, "GU")
	assert.Equal(t, Alpha2GT, "GT")
	assert.Equal(t, Alpha2GG, "GG")
	assert.Equal(t, Alpha2GN, "GN")
	assert.Equal(t, Alpha2GW, "GW")
	assert.Equal(t, Alpha2GY, "GY")
	assert.Equal(t, Alpha2HT, "HT")
	assert.Equal(t, Alpha2HM, "HM")
	assert.Equal(t, Alpha2VA, "VA")
	assert.Equal(t, Alpha2HN, "HN")
	assert.Equal(t, Alpha2HK, "HK")
	assert.Equal(t, Alpha2HU, "HU")
	assert.Equal(t, Alpha2IS, "IS")
	assert.Equal(t, Alpha2IN, "IN")
	assert.Equal(t, Alpha2ID, "ID")
	assert.Equal(t, Alpha2IR, "IR")
	assert.Equal(t, Alpha2IQ, "IQ")
	assert.Equal(t, Alpha2IE, "IE")
	assert.Equal(t, Alpha2IM, "IM")
	assert.Equal(t, Alpha2IL, "IL")
	assert.Equal(t, Alpha2IT, "IT")
	assert.Equal(t, Alpha2JM, "JM")
	assert.Equal(t, Alpha2JP, "JP") // Japan, should exist if all codes are present

	assert.Equal(t, Alpha2JE, "JE")
	assert.Equal(t, Alpha2JO, "JO")
	assert.Equal(t, Alpha2KZ, "KZ")
	assert.Equal(t, Alpha2KE, "KE")
	assert.Equal(t, Alpha2KI, "KI")
	assert.Equal(t, Alpha2KP, "KP")
	assert.Equal(t, Alpha2KR, "KR")
	assert.Equal(t, Alpha2KW, "KW")
	assert.Equal(t, Alpha2KG, "KG")
	assert.Equal(t, Alpha2LA, "LA")
	assert.Equal(t, Alpha2LV, "LV")
	assert.Equal(t, Alpha2LB, "LB")
	assert.Equal(t, Alpha2LS, "LS")
	assert.Equal(t, Alpha2LR, "LR")
	assert.Equal(t, Alpha2LY, "LY")
	assert.Equal(t, Alpha2LI, "LI")
	assert.Equal(t, Alpha2LT, "LT")
	assert.Equal(t, Alpha2LU, "LU")
	assert.Equal(t, Alpha2MO, "MO")
	assert.Equal(t, Alpha2MG, "MG")
	assert.Equal(t, Alpha2MW, "MW")
	assert.Equal(t, Alpha2MY, "MY")
	assert.Equal(t, Alpha2MV, "MV")
	assert.Equal(t, Alpha2ML, "ML")
	assert.Equal(t, Alpha2MT, "MT")
	assert.Equal(t, Alpha2MH, "MH")
	assert.Equal(t, Alpha2MQ, "MQ")
	assert.Equal(t, Alpha2MR, "MR")
	assert.Equal(t, Alpha2MU, "MU")
	assert.Equal(t, Alpha2YT, "YT")
	assert.Equal(t, Alpha2MX, "MX")
	assert.Equal(t, Alpha2FM, "FM")
	assert.Equal(t, Alpha2MD, "MD")
	assert.Equal(t, Alpha2MC, "MC")
	assert.Equal(t, Alpha2MN, "MN")
	assert.Equal(t, Alpha2ME, "ME")
	assert.Equal(t, Alpha2MS, "MS")
	assert.Equal(t, Alpha2MA, "MA")
	assert.Equal(t, Alpha2MZ, "MZ")
	assert.Equal(t, Alpha2MM, "MM")
	assert.Equal(t, Alpha2NA, "NA")
	assert.Equal(t, Alpha2NR, "NR")
	assert.Equal(t, Alpha2NP, "NP")
	assert.Equal(t, Alpha2NL, "NL")
	assert.Equal(t, Alpha2NC, "NC")
	assert.Equal(t, Alpha2NZ, "NZ")
	assert.Equal(t, Alpha2NI, "NI")
	assert.Equal(t, Alpha2NE, "NE")
	assert.Equal(t, Alpha2NG, "NG")
	assert.Equal(t, Alpha2NU, "NU")
	assert.Equal(t, Alpha2NF, "NF")
	assert.Equal(t, Alpha2MK, "MK")
	assert.Equal(t, Alpha2MP, "MP")
	assert.Equal(t, Alpha2NO, "NO")
	assert.Equal(t, Alpha2OM, "OM")
	assert.Equal(t, Alpha2PK, "PK")
	assert.Equal(t, Alpha2PW, "PW")
	assert.Equal(t, Alpha2PS, "PS")
	assert.Equal(t, Alpha2PA, "PA")
	assert.Equal(t, Alpha2PG, "PG")
	assert.Equal(t, Alpha2PY, "PY")
	assert.Equal(t, Alpha2PE, "PE")
	assert.Equal(t, Alpha2PH, "PH")
	assert.Equal(t, Alpha2PN, "PN")
	assert.Equal(t, Alpha2PL, "PL")
	assert.Equal(t, Alpha2PT, "PT")
	assert.Equal(t, Alpha2PR, "PR")
	assert.Equal(t, Alpha2QA, "QA")
	assert.Equal(t, Alpha2RE, "RE")
	assert.Equal(t, Alpha2RO, "RO")
	assert.Equal(t, Alpha2RU, "RU")
	assert.Equal(t, Alpha2RW, "RW")
	assert.Equal(t, Alpha2BL, "BL")
	assert.Equal(t, Alpha2SH, "SH")
	assert.Equal(t, Alpha2KN, "KN")
	assert.Equal(t, Alpha2LC, "LC")
	assert.Equal(t, Alpha2MF, "MF")
	assert.Equal(t, Alpha2PM, "PM")
	assert.Equal(t, Alpha2VC, "VC")
	assert.Equal(t, Alpha2WS, "WS")
	assert.Equal(t, Alpha2SM, "SM")
	assert.Equal(t, Alpha2ST, "ST")
	assert.Equal(t, Alpha2SA, "SA")
	assert.Equal(t, Alpha2SN, "SN")
	assert.Equal(t, Alpha2RS, "RS")
	assert.Equal(t, Alpha2SC, "SC")
	assert.Equal(t, Alpha2SL, "SL")
	assert.Equal(t, Alpha2SG, "SG")
	assert.Equal(t, Alpha2SX, "SX")
	assert.Equal(t, Alpha2SK, "SK")
	assert.Equal(t, Alpha2SI, "SI")
	assert.Equal(t, Alpha2SB, "SB")
	assert.Equal(t, Alpha2SO, "SO")
	assert.Equal(t, Alpha2ZA, "ZA")
	assert.Equal(t, Alpha2GS, "GS")
	assert.Equal(t, Alpha2SS, "SS")
	assert.Equal(t, Alpha2ES, "ES")
	assert.Equal(t, Alpha2LK, "LK")
	assert.Equal(t, Alpha2SD, "SD")
	assert.Equal(t, Alpha2SR, "SR")
	assert.Equal(t, Alpha2SJ, "SJ")
	assert.Equal(t, Alpha2SE, "SE")
	assert.Equal(t, Alpha2CH, "CH")
	assert.Equal(t, Alpha2SY, "SY")
	assert.Equal(t, Alpha2TW, "TW")
	assert.Equal(t, Alpha2TJ, "TJ")
	assert.Equal(t, Alpha2TZ, "TZ")
	assert.Equal(t, Alpha2TH, "TH")
	assert.Equal(t, Alpha2TL, "TL")
	assert.Equal(t, Alpha2TG, "TG")
	assert.Equal(t, Alpha2TK, "TK")
	assert.Equal(t, Alpha2TO, "TO")
	assert.Equal(t, Alpha2TT, "TT")
	assert.Equal(t, Alpha2TN, "TN")
	assert.Equal(t, Alpha2TR, "TR")
	assert.Equal(t, Alpha2TM, "TM")
	assert.Equal(t, Alpha2TC, "TC")
	assert.Equal(t, Alpha2TV, "TV")
	assert.Equal(t, Alpha2UG, "UG")
	assert.Equal(t, Alpha2UA, "UA")
	assert.Equal(t, Alpha2AE, "AE")
	assert.Equal(t, Alpha2GB, "GB")
	assert.Equal(t, Alpha2US, "US")
	assert.Equal(t, Alpha2UM, "UM")
	assert.Equal(t, Alpha2UY, "UY")
	assert.Equal(t, Alpha2UZ, "UZ")
	assert.Equal(t, Alpha2VU, "VU")
	assert.Equal(t, Alpha2VE, "VE")
	assert.Equal(t, Alpha2VN, "VN")
	assert.Equal(t, Alpha2VG, "VG")
	assert.Equal(t, Alpha2VI, "VI")
	assert.Equal(t, Alpha2WF, "WF")
	assert.Equal(t, Alpha2EH, "EH")
	assert.Equal(t, Alpha2YE, "YE")
	assert.Equal(t, Alpha2ZM, "ZM")
	assert.Equal(t, Alpha2ZW, "ZW")

	// Alpha-3 codes
	assert.Equal(t, Alpha3AFG, "AFG")
	assert.Equal(t, Alpha3ALA, "ALA")
	assert.Equal(t, Alpha3ALB, "ALB")
	assert.Equal(t, Alpha3DZA, "DZA")
	assert.Equal(t, Alpha3ASM, "ASM")
	assert.Equal(t, Alpha3AND, "AND")
	assert.Equal(t, Alpha3AGO, "AGO")
	assert.Equal(t, Alpha3AIA, "AIA")
	assert.Equal(t, Alpha3ATA, "ATA")
	assert.Equal(t, Alpha3ATG, "ATG")
	assert.Equal(t, Alpha3ARG, "ARG")
	assert.Equal(t, Alpha3ARM, "ARM")
	assert.Equal(t, Alpha3ABW, "ABW")
	assert.Equal(t, Alpha3AUS, "AUS")
	assert.Equal(t, Alpha3AUT, "AUT")
	assert.Equal(t, Alpha3AZE, "AZE")
	assert.Equal(t, Alpha3BHS, "BHS")
	assert.Equal(t, Alpha3BHR, "BHR")
	assert.Equal(t, Alpha3BGD, "BGD")
	assert.Equal(t, Alpha3BRB, "BRB")
	assert.Equal(t, Alpha3BLR, "BLR")
	assert.Equal(t, Alpha3BEL, "BEL")
	assert.Equal(t, Alpha3BLZ, "BLZ")
	assert.Equal(t, Alpha3BEN, "BEN")
	assert.Equal(t, Alpha3BMU, "BMU")
	assert.Equal(t, Alpha3BTN, "BTN")
	assert.Equal(t, Alpha3BOL, "BOL")
	assert.Equal(t, Alpha3BES, "BES")
	assert.Equal(t, Alpha3BIH, "BIH")
	assert.Equal(t, Alpha3BWA, "BWA")
	assert.Equal(t, Alpha3BVT, "BVT")
	assert.Equal(t, Alpha3BRA, "BRA")
	assert.Equal(t, Alpha3IOT, "IOT")
	assert.Equal(t, Alpha3BRN, "BRN")
	assert.Equal(t, Alpha3BGR, "BGR")
	assert.Equal(t, Alpha3BFA, "BFA")
	assert.Equal(t, Alpha3BDI, "BDI")
	assert.Equal(t, Alpha3CPV, "CPV")
	assert.Equal(t, Alpha3KHM, "KHM")
	assert.Equal(t, Alpha3CMR, "CMR")
	assert.Equal(t, Alpha3CAN, "CAN")
	assert.Equal(t, Alpha3CYM, "CYM")
	assert.Equal(t, Alpha3CAF, "CAF")
	assert.Equal(t, Alpha3TCD, "TCD")
	assert.Equal(t, Alpha3CHL, "CHL")
	assert.Equal(t, Alpha3CHN, "CHN")
	assert.Equal(t, Alpha3CXR, "CXR")
	assert.Equal(t, Alpha3CCK, "CCK")
	assert.Equal(t, Alpha3COL, "COL")
	assert.Equal(t, Alpha3COM, "COM")
	assert.Equal(t, Alpha3COG, "COG")
	assert.Equal(t, Alpha3COD, "COD")
	assert.Equal(t, Alpha3COK, "COK")
	assert.Equal(t, Alpha3CRI, "CRI")
	assert.Equal(t, Alpha3CIV, "CIV")
	assert.Equal(t, Alpha3HRV, "HRV")
	assert.Equal(t, Alpha3CUB, "CUB")
	assert.Equal(t, Alpha3CUW, "CUW")
	assert.Equal(t, Alpha3CYP, "CYP")
	assert.Equal(t, Alpha3CZE, "CZE")
	assert.Equal(t, Alpha3DNK, "DNK")
	assert.Equal(t, Alpha3DJI, "DJI")
	assert.Equal(t, Alpha3DMA, "DMA")
	assert.Equal(t, Alpha3DOM, "DOM")
	assert.Equal(t, Alpha3ECU, "ECU")
	assert.Equal(t, Alpha3EGY, "EGY")
	assert.Equal(t, Alpha3SLV, "SLV")
	assert.Equal(t, Alpha3GNQ, "GNQ")
	assert.Equal(t, Alpha3ERI, "ERI")
	assert.Equal(t, Alpha3EST, "EST")
	assert.Equal(t, Alpha3SWZ, "SWZ")
	assert.Equal(t, Alpha3ETH, "ETH")
	assert.Equal(t, Alpha3FLK, "FLK")
	assert.Equal(t, Alpha3FRO, "FRO")
	assert.Equal(t, Alpha3FJI, "FJI")
	assert.Equal(t, Alpha3FIN, "FIN")
	assert.Equal(t, Alpha3FRA, "FRA")
	assert.Equal(t, Alpha3GUF, "GUF")
	assert.Equal(t, Alpha3PYF, "PYF")
	assert.Equal(t, Alpha3ATF, "ATF")
	assert.Equal(t, Alpha3GAB, "GAB")
	assert.Equal(t, Alpha3GMB, "GMB")
	assert.Equal(t, Alpha3GEO, "GEO")
	assert.Equal(t, Alpha3DEU, "DEU")
	assert.Equal(t, Alpha3GHA, "GHA")
	assert.Equal(t, Alpha3GIB, "GIB")
	assert.Equal(t, Alpha3GRC, "GRC")
	assert.Equal(t, Alpha3GRL, "GRL")
	assert.Equal(t, Alpha3GRD, "GRD")
	assert.Equal(t, Alpha3GLP, "GLP")
	assert.Equal(t, Alpha3GUM, "GUM")
	assert.Equal(t, Alpha3GTM, "GTM")
	assert.Equal(t, Alpha3GGY, "GGY")
	assert.Equal(t, Alpha3GIN, "GIN")
	assert.Equal(t, Alpha3GNB, "GNB")
	assert.Equal(t, Alpha3GUY, "GUY")
	assert.Equal(t, Alpha3HTI, "HTI")
	assert.Equal(t, Alpha3HMD, "HMD")
	assert.Equal(t, Alpha3VAT, "VAT")
	assert.Equal(t, Alpha3HND, "HND")
	assert.Equal(t, Alpha3HKG, "HKG")
	assert.Equal(t, Alpha3HUN, "HUN")
	assert.Equal(t, Alpha3ISL, "ISL")
	assert.Equal(t, Alpha3IND, "IND")
	assert.Equal(t, Alpha3IDN, "IDN")
	assert.Equal(t, Alpha3IRN, "IRN")
	assert.Equal(t, Alpha3IRQ, "IRQ")
	assert.Equal(t, Alpha3IRL, "IRL")
	assert.Equal(t, Alpha3IMN, "IMN")
	assert.Equal(t, Alpha3ISR, "ISR")
	assert.Equal(t, Alpha3ITA, "ITA")
	assert.Equal(t, Alpha3JAM, "JAM")
	assert.Equal(t, Alpha3JPN, "JPN")
	assert.Equal(t, Alpha3JEY, "JEY")
	assert.Equal(t, Alpha3JOR, "JOR")
	assert.Equal(t, Alpha3KAZ, "KAZ")
	assert.Equal(t, Alpha3KEN, "KEN")
	assert.Equal(t, Alpha3KIR, "KIR")
	assert.Equal(t, Alpha3PRK, "PRK")
	assert.Equal(t, Alpha3KOR, "KOR")
	assert.Equal(t, Alpha3KWT, "KWT")
	assert.Equal(t, Alpha3KGZ, "KGZ")
	assert.Equal(t, Alpha3LAO, "LAO")
	assert.Equal(t, Alpha3LVA, "LVA")
	assert.Equal(t, Alpha3LBN, "LBN")
	assert.Equal(t, Alpha3LSO, "LSO")
	assert.Equal(t, Alpha3LBR, "LBR")
	assert.Equal(t, Alpha3LBY, "LBY")
	assert.Equal(t, Alpha3LIE, "LIE")
	assert.Equal(t, Alpha3LTU, "LTU")
	assert.Equal(t, Alpha3LUX, "LUX")
	assert.Equal(t, Alpha3MAC, "MAC")
	assert.Equal(t, Alpha3MDG, "MDG")
	assert.Equal(t, Alpha3MWI, "MWI")
	assert.Equal(t, Alpha3MYS, "MYS")
	assert.Equal(t, Alpha3MDV, "MDV")
	assert.Equal(t, Alpha3MLI, "MLI")
	assert.Equal(t, Alpha3MLT, "MLT")
	assert.Equal(t, Alpha3MHL, "MHL")
	assert.Equal(t, Alpha3MTQ, "MTQ")
	assert.Equal(t, Alpha3MRT, "MRT")
	assert.Equal(t, Alpha3MUS, "MUS")
	assert.Equal(t, Alpha3MYT, "MYT")
	assert.Equal(t, Alpha3MEX, "MEX")
	assert.Equal(t, Alpha3FSM, "FSM")
	assert.Equal(t, Alpha3MDA, "MDA")
	assert.Equal(t, Alpha3MCO, "MCO")
	assert.Equal(t, Alpha3MNG, "MNG")
	assert.Equal(t, Alpha3MNE, "MNE")
	assert.Equal(t, Alpha3MSR, "MSR")
	assert.Equal(t, Alpha3MAR, "MAR")
	assert.Equal(t, Alpha3MOZ, "MOZ")
	assert.Equal(t, Alpha3MMR, "MMR")
	assert.Equal(t, Alpha3NAM, "NAM")
	assert.Equal(t, Alpha3NRU, "NRU")
	assert.Equal(t, Alpha3NPL, "NPL")
	assert.Equal(t, Alpha3NLD, "NLD")
	assert.Equal(t, Alpha3NCL, "NCL")
	assert.Equal(t, Alpha3NZL, "NZL")
	assert.Equal(t, Alpha3NIC, "NIC")
	assert.Equal(t, Alpha3NER, "NER")
	assert.Equal(t, Alpha3NGA, "NGA")
	assert.Equal(t, Alpha3NIU, "NIU")
	assert.Equal(t, Alpha3NFK, "NFK")
	assert.Equal(t, Alpha3MKD, "MKD")
	assert.Equal(t, Alpha3MNP, "MNP")
	assert.Equal(t, Alpha3NOR, "NOR")
	assert.Equal(t, Alpha3OMN, "OMN")
	assert.Equal(t, Alpha3PAK, "PAK")
	assert.Equal(t, Alpha3PLW, "PLW")
	assert.Equal(t, Alpha3PSE, "PSE")
	assert.Equal(t, Alpha3PAN, "PAN")
	assert.Equal(t, Alpha3PNG, "PNG")
	assert.Equal(t, Alpha3PRY, "PRY")
	assert.Equal(t, Alpha3PER, "PER")
	assert.Equal(t, Alpha3PHL, "PHL")
	assert.Equal(t, Alpha3PCN, "PCN")
	assert.Equal(t, Alpha3POL, "POL")
	assert.Equal(t, Alpha3PRT, "PRT")
	assert.Equal(t, Alpha3PRI, "PRI")
	assert.Equal(t, Alpha3QAT, "QAT")
	assert.Equal(t, Alpha3REU, "REU")
	assert.Equal(t, Alpha3ROU, "ROU")
	assert.Equal(t, Alpha3RUS, "RUS")
	assert.Equal(t, Alpha3RWA, "RWA")
	assert.Equal(t, Alpha3BLM, "BLM")
	assert.Equal(t, Alpha3SHN, "SHN")
	assert.Equal(t, Alpha3KNA, "KNA")
	assert.Equal(t, Alpha3LCA, "LCA")
	assert.Equal(t, Alpha3MAF, "MAF")
	assert.Equal(t, Alpha3SPM, "SPM")
	assert.Equal(t, Alpha3VCT, "VCT")
	assert.Equal(t, Alpha3WSM, "WSM")
	assert.Equal(t, Alpha3SMR, "SMR")
	assert.Equal(t, Alpha3STP, "STP")
	assert.Equal(t, Alpha3SAU, "SAU")
	assert.Equal(t, Alpha3SEN, "SEN")
	assert.Equal(t, Alpha3SRB, "SRB")
	assert.Equal(t, Alpha3SYC, "SYC")
	assert.Equal(t, Alpha3SLE, "SLE")
	assert.Equal(t, Alpha3SGP, "SGP")
	assert.Equal(t, Alpha3SXM, "SXM")
	assert.Equal(t, Alpha3SVK, "SVK")
	assert.Equal(t, Alpha3SVN, "SVN")
	assert.Equal(t, Alpha3SLB, "SLB")
	assert.Equal(t, Alpha3SOM, "SOM")
	assert.Equal(t, Alpha3ZAF, "ZAF")
	assert.Equal(t, Alpha3SGS, "SGS")
	assert.Equal(t, Alpha3SSD, "SSD")
	assert.Equal(t, Alpha3ESP, "ESP")
	assert.Equal(t, Alpha3LKA, "LKA")
	assert.Equal(t, Alpha3SDN, "SDN")
	assert.Equal(t, Alpha3SUR, "SUR")
	assert.Equal(t, Alpha3SJM, "SJM")
	assert.Equal(t, Alpha3SWE, "SWE")
	assert.Equal(t, Alpha3CHE, "CHE")
	assert.Equal(t, Alpha3SYR, "SYR")
	assert.Equal(t, Alpha3TWN, "TWN")
	assert.Equal(t, Alpha3TJK, "TJK")
	assert.Equal(t, Alpha3TZA, "TZA")
	assert.Equal(t, Alpha3THA, "THA")
	assert.Equal(t, Alpha3TLS, "TLS")
	assert.Equal(t, Alpha3TGO, "TGO")
	assert.Equal(t, Alpha3TKL, "TKL")
	assert.Equal(t, Alpha3TON, "TON")
	assert.Equal(t, Alpha3TTO, "TTO")
	assert.Equal(t, Alpha3TUN, "TUN")
	assert.Equal(t, Alpha3TUR, "TUR")
	assert.Equal(t, Alpha3TKM, "TKM")
	assert.Equal(t, Alpha3TCA, "TCA")
	assert.Equal(t, Alpha3TUV, "TUV")
	assert.Equal(t, Alpha3UGA, "UGA")
	assert.Equal(t, Alpha3UKR, "UKR")
	assert.Equal(t, Alpha3ARE, "ARE")
	assert.Equal(t, Alpha3GBR, "GBR")
	assert.Equal(t, Alpha3USA, "USA")
	assert.Equal(t, Alpha3UMI, "UMI")
	assert.Equal(t, Alpha3URY, "URY")
	assert.Equal(t, Alpha3UZB, "UZB")
	assert.Equal(t, Alpha3VUT, "VUT")
	assert.Equal(t, Alpha3VEN, "VEN")
	assert.Equal(t, Alpha3VNM, "VNM")
	assert.Equal(t, Alpha3VGB, "VGB")
	assert.Equal(t, Alpha3VIR, "VIR")
	assert.Equal(t, Alpha3WLF, "WLF")
	assert.Equal(t, Alpha3ESH, "ESH")
	assert.Equal(t, Alpha3YEM, "YEM")
	assert.Equal(t, Alpha3ZMB, "ZMB")
	assert.Equal(t, Alpha3ZWE, "ZWE")
}

// TestCountryCapitals_Correctness checks the correctness of country capitals.
func TestCountryCapitals_Correctness(t *testing.T) {
	tests := []struct {
		name        string
		countryName string
		expected    string
	}{
		{name: "Afghanistan", countryName: "Afghanistan", expected: "Kabul"},
		{name: "Australia", countryName: "Australia", expected: "Canberra"},
		{name: "Brazil", countryName: "Brazil", expected: "Brasília"},
		{name: "Canada", countryName: "Canada", expected: "Ottawa"},
		{name: "France", countryName: "France", expected: "Paris"},
		{name: "Japan", countryName: "Japan", expected: "Tokyo"},
		{name: "United Kingdom", countryName: "United Kingdom of Great Britain and Northern Ireland", expected: "London"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			country := GetByName(tt.countryName)
			require.NotNil(t, country)
			assert.Equal(t, tt.expected, country.Capital)
		})
	}
}
