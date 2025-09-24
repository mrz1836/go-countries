package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"sort"
	"strings"
	"text/template"
	"time"
)

// Generator handles the country data generation process
type Generator struct {
	dataLoader       DataLoader
	fileWriter       FileWriter
	templateProvider TemplateProvider
	outputPath       string
	repoURL          string
}

// GeneratorConfig holds configuration for the Generator
type GeneratorConfig struct {
	DataLoader       DataLoader
	FileWriter       FileWriter
	TemplateProvider TemplateProvider
	OutputPath       string
	RepoURL          string
}

// NewGenerator creates a new Generator instance with the provided dependencies
func NewGenerator(config GeneratorConfig) *Generator {
	return &Generator{
		dataLoader:       config.DataLoader,
		fileWriter:       config.FileWriter,
		templateProvider: config.TemplateProvider,
		outputPath:       config.OutputPath,
		repoURL:          config.RepoURL,
	}
}

// Generate performs the complete country data generation process
func (g *Generator) Generate() error {
	countries, err := g.LoadCountries()
	if err != nil {
		return fmt.Errorf("failed to load countries: %w", err)
	}

	currencies, err := g.LoadCurrencies()
	if err != nil {
		return fmt.Errorf("failed to load currencies: %w", err)
	}

	g.MergeData(countries, currencies)

	capitals := g.GenerateCapitalMap(countries)

	code, err := g.GenerateCode(countries, capitals)
	if err != nil {
		return fmt.Errorf("failed to generate code: %w", err)
	}

	if err := g.WriteOutput(code); err != nil {
		return fmt.Errorf("failed to write output: %w", err)
	}

	return nil
}

// LoadCountries loads and parses the ISO 3166 country data
func (g *Generator) LoadCountries() (CountryList, error) {
	data, err := g.dataLoader.LoadISO3166Data()
	if err != nil {
		return nil, fmt.Errorf("failed to load ISO3166 data: %w", err)
	}

	var countries CountryList
	if err := json.Unmarshal(data, &countries); err != nil {
		return nil, fmt.Errorf("failed to unmarshal countries data: %w", err)
	}

	return countries, nil
}

// LoadCurrencies loads and parses the currency data
func (g *Generator) LoadCurrencies() (countriesWithCurrencies, error) {
	data, err := g.dataLoader.LoadCurrencyData()
	if err != nil {
		return nil, fmt.Errorf("failed to load currency data: %w", err)
	}

	var currencies countriesWithCurrencies
	if err := json.Unmarshal(data, &currencies); err != nil {
		return nil, fmt.Errorf("failed to unmarshal currency data: %w", err)
	}

	return currencies, nil
}

// MergeData combines country and currency data
func (g *Generator) MergeData(countries CountryList, currencies countriesWithCurrencies) {
	for index, country := range countries {
		for _, altCountry := range currencies {
			if country.Alpha2 == altCountry.CountryCode {
				countries[index].Capital = altCountry.Capital
				countries[index].ContinentName = altCountry.ContinentName
				countries[index].CurrencyCode = altCountry.CurrencyCode
				break
			}
		}
	}
}

// GenerateCapitalMap creates a sorted map of capitals to country indices
func (g *Generator) GenerateCapitalMap(countries CountryList) []mapEntry {
	capitalSeen := make(map[string]struct{})
	var capitalEntries []mapEntry

	for index, country := range countries {
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

	return capitalEntries
}

// GenerateCode generates the formatted Go source code
func (g *Generator) GenerateCode(countries CountryList, capitals []mapEntry) ([]byte, error) {
	templateStr, err := g.templateProvider.GetPackageTemplate()
	if err != nil {
		return nil, fmt.Errorf("failed to get template: %w", err)
	}

	tmpl := template.Must(template.New("").Funcs(template.FuncMap{
		"lower": strings.ToLower,
	}).Parse(templateStr))

	var buf bytes.Buffer
	if execErr := tmpl.Execute(&buf, struct {
		Timestamp time.Time
		URL       string
		Countries CountryList
		Capitals  []mapEntry
	}{
		Timestamp: time.Now(),
		URL:       g.repoURL,
		Countries: countries,
		Capitals:  capitals,
	}); execErr != nil {
		return nil, fmt.Errorf("template execution failed: %w", execErr)
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("failed to format generated source: %w", err)
	}

	return formatted, nil
}

// WriteOutput writes the generated code to the output file
func (g *Generator) WriteOutput(code []byte) error {
	file, err := g.fileWriter.Create(g.outputPath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer func() {
		_ = file.Close() // ignore close error in defer
	}()

	if _, err := file.Write(code); err != nil {
		return fmt.Errorf("failed to write code to file: %w", err)
	}

	return nil
}
