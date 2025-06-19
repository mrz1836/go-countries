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
