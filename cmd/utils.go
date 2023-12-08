package main

import (
	"os"
)

var (
	// Map of file name to source name if they don't match the sourceName
	overrides = map[string]string{
		"Error":         "Errors",
		"OrderExecutor": "",
	}
)

func getCorrectFile(file string, directories []string) *os.File {
	for _, directory := range directories {
		f, err := os.Open(directory + "/" + file)
		if err == nil {
			return f
		}
	}
	return nil
}
