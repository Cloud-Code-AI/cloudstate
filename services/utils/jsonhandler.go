package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func GetJSONFiles(region string) []string {
	directory := "./output/" + region
	filePaths := make([]string, 0)
	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return filePaths
	}
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			filePaths = append(filePaths, directory+"/"+file.Name())
		}
	}
	return filePaths
}

func ReadJSONFile(filePath string) (interface{}, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Unmarshal the JSON data
	var jsonData interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return jsonData, err
}
