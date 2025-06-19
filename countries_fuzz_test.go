package countries

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// FuzzGetByAlpha2 checks that GetByAlpha2 correctly handles
// alpha-2 codes with varying cases and returns a valid Country
// when the input represents a known code.
func FuzzGetByAlpha2(f *testing.F) {
	seed := []string{"us", "GB", "Ca", "zz", ""}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, code string) {
		c := GetByAlpha2(code)
		if c == nil {
			return
		}
		require.Len(t, c.Alpha2, 2)
		require.Equal(t, strings.ToUpper(code), c.Alpha2)
	})
}

// FuzzGetByName verifies GetByName returns consistent results for varying input cases.
func FuzzGetByName(f *testing.F) {
	seed := []string{"United States of America", "canada", "France", "", "zzzz"}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, name string) {
		c := GetByName(name)
		if c == nil {
			return
		}
		require.NotEmpty(t, c.Name)
		require.Equal(t, c, GetByAlpha2(c.Alpha2))
	})
}

// FuzzGetByAlpha3 ensures GetByAlpha3 handles different case inputs correctly.
func FuzzGetByAlpha3(f *testing.F) {
	seed := []string{"usa", "CAN", "gBr", "", "xx"}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, code string) {
		c := GetByAlpha3(code)
		if c == nil {
			return
		}
		require.Len(t, c.Alpha3, 3)
		require.Equal(t, strings.ToUpper(code), c.Alpha3)
	})
}

// FuzzGetByCountryCode checks GetByCountryCode with numeric codes of varying lengths.
func FuzzGetByCountryCode(f *testing.F) {
	seed := []string{"840", "036", "124", "1", "12345"}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, code string) {
		c := GetByCountryCode(code)
		if c == nil {
			return
		}
		require.Equal(t, code, c.CountryCode)
	})
}

// FuzzGetByCapital verifies GetByCapital returns a valid country for known capitals.
func FuzzGetByCapital(f *testing.F) {
	seed := []string{"Washington", "london", "Tokyo", "", "Unknown"}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, capital string) {
		c := GetByCapital(capital)
		if c == nil {
			return
		}
		require.True(t, strings.EqualFold(capital, c.Capital))
	})
}

// FuzzGetByISO31662 checks GetByISO31662 with uppercase and lowercase inputs.
func FuzzGetByISO31662(f *testing.F) {
	seed := []string{"ISO 3166-2:US", "iso 3166-2:ca", "ISO 3166-2:GB", "", "ISO"}
	for _, s := range seed {
		f.Add(s)
	}
	f.Fuzz(func(t *testing.T, iso string) {
		c := GetByISO31662(iso)
		if c == nil {
			return
		}
		require.Equal(t, strings.ToUpper(iso), c.ISO31662)
	})
}
