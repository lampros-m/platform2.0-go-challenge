package test

import (
	"os"
	"testing"
)

// FileReader : Struct provides functionality to read json file from test folder.
type FileReader struct {
	t *testing.T
}

// NewFileReader : Constructor for file reader.
func NewFileReader(t *testing.T) *FileReader {
	t.Helper()
	return &FileReader{t: t}
}

// ReadFile reads a file from the current dir and returns the content as string.`.
func (f *FileReader) ReadFile(file string) string {
	f.t.Helper()
	path, err := os.Getwd()
	if err != nil {
		f.t.Fatalf("Error in resource file open: %s", err.Error())
	}

	c, err := os.ReadFile(path + "/" + file)
	if err != nil {
		f.t.Fatalf("Error in resource file open: %s", err.Error())
	}
	return string(c)
}
