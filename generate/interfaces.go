package main

import (
	"io"
)

// DataLoader provides access to embedded JSON data sources
type DataLoader interface {
	LoadISO3166Data() ([]byte, error)
	LoadCurrencyData() ([]byte, error)
}

// FileWriter handles file operations for output generation
type FileWriter interface {
	Create(filename string) (io.WriteCloser, error)
}

// TemplateProvider provides access to code generation templates
type TemplateProvider interface {
	GetPackageTemplate() (string, error)
}
