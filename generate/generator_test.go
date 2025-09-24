package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Static errors for testing
var (
	errFailedToLoadData     = errors.New("failed to load data")
	errFailedToLoadCurrency = errors.New("failed to load currency data")
	errTemplateError        = errors.New("template error")
	errFailedToCreateFile   = errors.New("failed to create file")
	errLoadError            = errors.New("load error")
	errCurrencyError        = errors.New("currency error")
)

func TestNewGenerator(t *testing.T) {
	config := GeneratorConfig{
		DataLoader:       &MockDataLoader{},
		FileWriter:       &MockFileWriter{},
		TemplateProvider: &MockTemplateProvider{},
		OutputPath:       "test.go",
		RepoURL:          "https://github.com/test/repo",
	}

	generator := NewGenerator(config)

	assert.NotNil(t, generator)
	assert.Equal(t, "test.go", generator.outputPath)
	assert.Equal(t, "https://github.com/test/repo", generator.repoURL)
}

func TestGenerator_LoadCountries_Success(t *testing.T) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	countries, err := generator.LoadCountries()

	require.NoError(t, err)
	assert.Len(t, countries, 2)
	assert.Equal(t, "Test Country", countries[0].Name)
	assert.Equal(t, "TC", countries[0].Alpha2)
	assert.Equal(t, "TST", countries[0].Alpha3)
}

func TestGenerator_LoadCountries_DataLoaderError(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.ISO3166Error = errFailedToLoadData

	countries, err := generator.LoadCountries()

	require.Error(t, err)
	assert.Nil(t, countries)
	assert.Contains(t, err.Error(), "failed to load ISO3166 data")
}

func TestGenerator_LoadCountries_InvalidJSON(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.ISO3166Data = []byte("invalid json")

	countries, err := generator.LoadCountries()

	require.Error(t, err)
	assert.Nil(t, countries)
	assert.Contains(t, err.Error(), "failed to unmarshal countries data")
}

func TestGenerator_LoadCurrencies_Success(t *testing.T) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	currencies, err := generator.LoadCurrencies()

	require.NoError(t, err)
	assert.Len(t, currencies, 2)
	assert.Equal(t, "TC", currencies[0].CountryCode)
	assert.Equal(t, "Test Country", currencies[0].CountryName)
	assert.Equal(t, "TST", currencies[0].CurrencyCode)
}

func TestGenerator_LoadCurrencies_DataLoaderError(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.CurrencyError = errFailedToLoadCurrency

	currencies, err := generator.LoadCurrencies()

	require.Error(t, err)
	assert.Nil(t, currencies)
	assert.Contains(t, err.Error(), "failed to load currency data")
}

func TestGenerator_LoadCurrencies_InvalidJSON(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.CurrencyData = []byte("invalid json")

	currencies, err := generator.LoadCurrencies()

	require.Error(t, err)
	assert.Nil(t, currencies)
	assert.Contains(t, err.Error(), "failed to unmarshal currency data")
}

func TestGenerator_MergeData(t *testing.T) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	countries, err := generator.LoadCountries()
	require.NoError(t, err)

	currencies, err := generator.LoadCurrencies()
	require.NoError(t, err)

	generator.MergeData(countries, currencies)
	assert.Equal(t, "Test Capital", countries[0].Capital)
	assert.Equal(t, "Test Continent", countries[0].ContinentName)
	assert.Equal(t, "TST", countries[0].CurrencyCode)
	assert.Equal(t, "Another Capital", countries[1].Capital)
	assert.Equal(t, "Another Continent", countries[1].ContinentName)
	assert.Equal(t, "ANO", countries[1].CurrencyCode)
}

func TestGenerator_GenerateCapitalMap(t *testing.T) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	countries := CountryList{
		{Capital: "Washington", Name: "USA"},
		{Capital: "Ottawa", Name: "Canada"},
		{Capital: "Berlin", Name: "Germany"},
		{Capital: "", Name: "Country without capital"},
		{Capital: "Washington", Name: "Duplicate capital"},
	}

	capitals := generator.GenerateCapitalMap(countries)

	assert.Len(t, capitals, 3)
	assert.Equal(t, "berlin", capitals[0].Key)
	assert.Equal(t, 2, capitals[0].Index)
	assert.Equal(t, "ottawa", capitals[1].Key)
	assert.Equal(t, 1, capitals[1].Index)
	assert.Equal(t, "washington", capitals[2].Key)
	assert.Equal(t, 0, capitals[2].Index)
}

func TestGenerator_GenerateCode_Success(t *testing.T) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	countries := CountryList{
		{Name: "Test Country", Alpha2: "TC"},
	}
	capitals := []mapEntry{{Key: "test capital", Index: 0}}

	code, err := generator.GenerateCode(countries, capitals)

	require.NoError(t, err)
	assert.Contains(t, string(code), "package countries")
	assert.Contains(t, string(code), "Test Country")
	assert.Contains(t, string(code), "TC")
}

func TestGenerator_GenerateCode_TemplateError(t *testing.T) {
	generator, _, _, mockTemplate := NewTestGenerator()
	mockTemplate.Error = errTemplateError

	countries := CountryList{}
	capitals := []mapEntry{}

	code, err := generator.GenerateCode(countries, capitals)

	require.Error(t, err)
	assert.Nil(t, code)
	assert.Contains(t, err.Error(), "failed to get template")
}

func TestGenerator_WriteOutput_Success(t *testing.T) {
	generator, _, mockWriter, _ := NewTestGenerator()
	testCode := []byte("package test\nvar x = 1")

	err := generator.WriteOutput(testCode)

	require.NoError(t, err)
	assert.Contains(t, mockWriter.CreatedFiles, "test_output.go")
	assert.Equal(t, testCode, mockWriter.CreatedFiles["test_output.go"].Bytes())
}

func TestGenerator_WriteOutput_CreateError(t *testing.T) {
	generator, _, mockWriter, _ := NewTestGenerator()
	mockWriter.CreateError = errFailedToCreateFile
	testCode := []byte("package test")

	err := generator.WriteOutput(testCode)

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to create output file")
}

func TestGenerator_Generate_EndToEnd(t *testing.T) {
	generator, _, mockWriter, _ := NewTestGenerator()

	err := generator.Generate()

	require.NoError(t, err)
	assert.Contains(t, mockWriter.CreatedFiles, "test_output.go")

	output := mockWriter.CreatedFiles["test_output.go"].String()
	assert.Contains(t, output, "package countries")
	assert.Contains(t, output, "Test Country")
	assert.Contains(t, output, "Another Country")
}

func TestGenerator_Generate_LoadCountriesError(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.ISO3166Error = errLoadError

	err := generator.Generate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load countries")
}

func TestGenerator_Generate_LoadCurrenciesError(t *testing.T) {
	generator, mockLoader, _, _ := NewTestGenerator()
	mockLoader.CurrencyError = errCurrencyError

	err := generator.Generate()

	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to load currencies")
}

// Integration tests with real implementations
func TestEmbeddedDataLoader_Integration(t *testing.T) {
	loader := &EmbeddedDataLoader{}

	iso3166Data, err := loader.LoadISO3166Data()
	require.NoError(t, err)
	assert.NotEmpty(t, iso3166Data)
	assert.Contains(t, string(iso3166Data), "United States of America")

	currencyData, err := loader.LoadCurrencyData()
	require.NoError(t, err)
	assert.NotEmpty(t, currencyData)
	assert.Contains(t, string(currencyData), "USD")
}

func TestOSFileWriter_Integration(t *testing.T) {
	writer := &OSFileWriter{}
	testPath := "/tmp/test_generator_file.txt"

	defer func() {
		_ = os.Remove(testPath) // ignore remove error in test cleanup
	}()

	file, err := writer.Create(testPath)
	require.NoError(t, err)
	defer func() {
		_ = file.Close() // ignore close error in defer
	}()

	testData := []byte("test content")
	n, err := file.Write(testData)
	require.NoError(t, err)
	assert.Equal(t, len(testData), n)

	err = file.Close()
	require.NoError(t, err)

	content, err := os.ReadFile(testPath)
	require.NoError(t, err)
	assert.Equal(t, testData, content)
}

func TestDefaultTemplateProvider_Integration(t *testing.T) {
	provider := &DefaultTemplateProvider{}

	template, err := provider.GetPackageTemplate()
	require.NoError(t, err)
	assert.NotEmpty(t, template)
	assert.Contains(t, template, "package countries")
	assert.Contains(t, template, "var (")
	assert.Contains(t, template, "countries = []*Country{")
}

// Benchmark tests
func BenchmarkGenerator_Generate(b *testing.B) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := generator.Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGenerator_MergeData(b *testing.B) {
	generator, mockLoader, mockWriter, mockTemplate := NewTestGenerator()
	_ = mockLoader
	_ = mockWriter
	_ = mockTemplate

	countries, _ := generator.LoadCountries()
	currencies, _ := generator.LoadCurrencies()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		generator.MergeData(countries, currencies)
	}
}

// Test helper for file-based data loading
func TestGeneratorWithFileData(t *testing.T) {
	countryData, err := os.ReadFile("testdata/test_countries.json")
	require.NoError(t, err)

	currencyData, err := os.ReadFile("testdata/test_currencies.json")
	require.NoError(t, err)

	mockLoader := &MockDataLoader{
		ISO3166Data:  countryData,
		CurrencyData: currencyData,
	}

	mockWriter := NewMockFileWriter()
	mockTemplate := &MockTemplateProvider{Template: (&TestDataProvider{}).GetSimpleTemplate()}

	config := GeneratorConfig{
		DataLoader:       mockLoader,
		FileWriter:       mockWriter,
		TemplateProvider: mockTemplate,
		OutputPath:       "file_test_output.go",
		RepoURL:          "https://github.com/test/file-repo",
	}

	generator := NewGenerator(config)

	err = generator.Generate()
	require.NoError(t, err)

	output := mockWriter.CreatedFiles["file_test_output.go"].String()
	assert.Contains(t, output, "United States of America")
	assert.Contains(t, output, "Canada")
	assert.Contains(t, output, "Germany")
}
