// generate/generate.go
//go:generate go run generate.go

// This program generates the structs from JSON
package main

import (
	"bytes"
	"encoding/json"
	"go/format"
	"log"
	"os"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/mrz1836/go-countries/data"
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

	// Unmarshall the countries
	var c CountryList
	if err := json.Unmarshal([]byte(data.ISO3166JSONData), &c); err != nil {
		log.Panic("failed to load countries", err.Error())
	}

	// Unmarshall the temporary data
	var ca countriesWithCurrencies
	if err := json.Unmarshal([]byte(data.CountryCurrencyJSONData), &ca); err != nil {
		log.Panic("failed to load alternate country data", err.Error())
	}

	// Loop and combine the data
	capitalSeen := make(map[string]struct{})
	var capitalEntries []mapEntry

	for index, country := range c {

		// Loop alternate data
		for _, altCountry := range ca {
			if country.Alpha2 == altCountry.CountryCode {
				c[index].Capital = altCountry.Capital
				c[index].ContinentName = altCountry.ContinentName
				c[index].CurrencyCode = altCountry.CurrencyCode
			}
		}

		if country.Capital != "" {
			key := strings.ToLower(country.Capital)
			if _, ok := capitalSeen[key]; !ok {
				capitalSeen[key] = struct{}{}
				capitalEntries = append(capitalEntries, mapEntry{Key: key, Index: index})
			}
		}
	}

	sort.Slice(capitalEntries, func(i, j int) bool {
		return capitalEntries[i].Key < capitalEntries[j].Key
	})

	// Repo URL
	const url = "https://github.com/mrz1836/go-countries"

	// Create the file or overwrite
	f, err := os.Create("../countries_data.go")
	if err != nil {
		log.Fatal(err)
	}

	// Close the file
	defer func() {
		_ = f.Close()
	}()

	// Execute the template into a buffer
	var buf bytes.Buffer
	packageTemplate := getPackageTemplate()
	if err = packageTemplate.Execute(&buf, struct {
		Timestamp time.Time
		URL       string
		Countries CountryList
		Capitals  []mapEntry
	}{
		Timestamp: time.Now(),
		URL:       url,
		Countries: c,
		Capitals:  capitalEntries,
	}); err != nil {
		log.Fatalf("package template execution failed: %v", err)
	}

	// Format the generated source
	var formatted []byte
	if formatted, err = format.Source(buf.Bytes()); err != nil {
		log.Fatalf("failed to format generated source: %v", err)
	}

	// Write the formatted source to the file
	if _, err = f.Write(formatted); err != nil {
		log.Fatalf("failed to write formatted source: %v", err)
	}
}

// getPackageTemplate returns a pre-defined template for generating the country data file.
func getPackageTemplate() *template.Template {
	return template.Must(template.New("").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).Parse(`// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots at
// {{ .URL }}
package countries

var (
	countries = []*Country{
	{{- range .Countries }}
		{
			Alpha2:                 {{ printf "%q" .Alpha2 }},
			Alpha3:                 {{ printf "%q" .Alpha3 }},
			Capital:             	{{ printf "%q" .Capital }},
			ContinentName:          {{ printf "%q" .ContinentName }},
			CountryCode:            {{ printf "%q" .CountryCode }},
			CurrencyCode:           {{ printf "%q" .CurrencyCode }},
			IntermediateRegion:     {{ printf "%q" .IntermediateRegion }},
			IntermediateRegionCode: {{ printf "%q" .IntermediateRegionCode }},
			ISO31662:               {{ printf "%q" .ISO31662 }},
			Name:                   {{ printf "%q" .Name }},
			Region:                 {{ printf "%q" .Region }},
			RegionCode:             {{ printf "%q" .RegionCode }},
			SubRegion:              {{ printf "%q" .SubRegion }},
			SubRegionCode:          {{ printf "%q" .SubRegionCode }},
		},
        {{- end }}
        }

        byName = map[string]*Country{
        {{- range $index, $c := .Countries }}
                {{ printf "%q" (lower $c.Name) }}: countries[{{ $index }}],
        {{- end }}
        }

        byAlpha2 = map[string]*Country{
        {{- range $index, $c := .Countries }}
                {{ printf "%q" $c.Alpha2 }}: countries[{{ $index }}],
        {{- end }}
        }

        byAlpha3 = map[string]*Country{
        {{- range $index, $c := .Countries }}
                {{ printf "%q" $c.Alpha3 }}: countries[{{ $index }}],
        {{- end }}
        }

        byCode = map[string]*Country{
        {{- range $index, $c := .Countries }}
                {{ printf "%q" $c.CountryCode }}: countries[{{ $index }}],
        {{- end }}
        }

        byCapital = map[string]*Country{
        {{- range $_, $pair := .Capitals }}
                {{ printf "%q" $pair.Key }}: countries[{{ $pair.Index }}],
        {{- end }}
        }

        byISO31662 = map[string]*Country{
        {{- range $index, $c := .Countries }}
                {{ printf "%q" $c.ISO31662 }}: countries[{{ $index }}],
        {{- end }}
        }
)
`))
}
