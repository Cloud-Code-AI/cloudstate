package utils

import (
	"encoding/json"
	"os"
)

// WriteJSONToFile writes the given data as JSON to the specified file path.
// It creates the directory if it doesn't exist.
func WriteJSONToFile(filePath string, data interface{}) error {
	// Marshal data to JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Create directory if it doesn't exist
	if err := os.MkdirAll(getDir(filePath), 0755); err != nil {
		return err
	}

	// Write data to file
	return os.WriteFile(filePath, jsonData, 0644)
}

// getDir returns the directory part of the file path.
func getDir(filePath string) string {
	lastSlashIndex := -1
	for i := range filePath {
		if filePath[i] == '/' || filePath[i] == '\\' {
			lastSlashIndex = i
		}
	}
	if lastSlashIndex == -1 {
		return "."
	}
	return filePath[:lastSlashIndex]
}
