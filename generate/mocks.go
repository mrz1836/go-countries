package main

import (
	"bytes"
	"io"
)

// MockDataLoader is a mock implementation of DataLoader for testing
type MockDataLoader struct {
	ISO3166Data   []byte
	CurrencyData  []byte
	ISO3166Error  error
	CurrencyError error
}

func (m *MockDataLoader) LoadISO3166Data() ([]byte, error) {
	if m.ISO3166Error != nil {
		return nil, m.ISO3166Error
	}
	return m.ISO3166Data, nil
}

func (m *MockDataLoader) LoadCurrencyData() ([]byte, error) {
	if m.CurrencyError != nil {
		return nil, m.CurrencyError
	}
	return m.CurrencyData, nil
}

// MockFileWriter is a mock implementation of FileWriter for testing
type MockFileWriter struct {
	CreatedFiles map[string]*bytes.Buffer
	CreateError  error
}

func NewMockFileWriter() *MockFileWriter {
	return &MockFileWriter{
		CreatedFiles: make(map[string]*bytes.Buffer),
	}
}

func (m *MockFileWriter) Create(filename string) (io.WriteCloser, error) {
	if m.CreateError != nil {
		return nil, m.CreateError
	}

	buf := &bytes.Buffer{}
	m.CreatedFiles[filename] = buf
	return &mockWriteCloser{buf: buf}, nil
}

type mockWriteCloser struct {
	buf *bytes.Buffer
}

func (m *mockWriteCloser) Write(p []byte) (n int, err error) {
	return m.buf.Write(p)
}

func (m *mockWriteCloser) Close() error {
	return nil
}

// MockTemplateProvider is a mock implementation of TemplateProvider for testing
type MockTemplateProvider struct {
	Template string
	Error    error
}

func (m *MockTemplateProvider) GetPackageTemplate() (string, error) {
	if m.Error != nil {
		return "", m.Error
	}
	return m.Template, nil
}

// TestDataProvider provides sample data for testing
type TestDataProvider struct{}

func (t *TestDataProvider) GetSampleISO3166Data() []byte {
	return []byte(`[
		{
			"name": "Test Country",
			"alpha-2": "TC",
			"alpha-3": "TST",
			"country-code": "999",
			"iso_3166-2": "ISO 3166-2:TC",
			"region": "Test Region",
			"sub-region": "Test Sub-Region",
			"intermediate-region": "",
			"region-code": "001",
			"sub-region-code": "002",
			"intermediate-region-code": ""
		},
		{
			"name": "Another Country",
			"alpha-2": "AC",
			"alpha-3": "ANO",
			"country-code": "998",
			"iso_3166-2": "ISO 3166-2:AC",
			"region": "Test Region",
			"sub-region": "Another Sub-Region",
			"intermediate-region": "Test Intermediate",
			"region-code": "001",
			"sub-region-code": "003",
			"intermediate-region-code": "004"
		}
	]`)
}

func (t *TestDataProvider) GetSampleCurrencyData() []byte {
	return []byte(`[
		{
			"countryCode": "TC",
			"countryName": "Test Country",
			"currencyCode": "TST",
			"population": "1000000",
			"capital": "Test Capital",
			"continentName": "Test Continent"
		},
		{
			"countryCode": "AC",
			"countryName": "Another Country",
			"currencyCode": "ANO",
			"population": "2000000",
			"capital": "Another Capital",
			"continentName": "Another Continent"
		}
	]`)
}

func (t *TestDataProvider) GetSimpleTemplate() string {
	return `// Test Template
package countries

var (
	countries = []*Country{
	{{- range .Countries }}
		{
			Name: {{ printf "%q" .Name }},
			Alpha2: {{ printf "%q" .Alpha2 }},
		},
	{{- end }}
	}
)`
}

// Helper functions for tests
func NewTestGenerator() (*Generator, *MockDataLoader, *MockFileWriter, *MockTemplateProvider) {
	dataProvider := &TestDataProvider{}

	mockDataLoader := &MockDataLoader{
		ISO3166Data:  dataProvider.GetSampleISO3166Data(),
		CurrencyData: dataProvider.GetSampleCurrencyData(),
	}

	mockFileWriter := NewMockFileWriter()

	mockTemplateProvider := &MockTemplateProvider{
		Template: dataProvider.GetSimpleTemplate(),
	}

	config := GeneratorConfig{
		DataLoader:       mockDataLoader,
		FileWriter:       mockFileWriter,
		TemplateProvider: mockTemplateProvider,
		OutputPath:       "test_output.go",
		RepoURL:          "https://github.com/test/repo",
	}

	generator := NewGenerator(config)
	return generator, mockDataLoader, mockFileWriter, mockTemplateProvider
}
