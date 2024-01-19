package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
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

func ReadJSONFile(filePath string) (map[string]interface{}, error) {
	// Read the JSON file
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Unmarshal the JSON data
	var jsonData map[string]interface{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil, err
	}
	return jsonData, err
}

func collectFilePaths(path string, info os.FileInfo, err error, files *[]string) error {
	if err != nil {
		return err
	}
	if !info.IsDir() {
		*files = append(*files, path)
	}
	return nil
}

func visit(root string, path string, f os.FileInfo) (interface{}, error) {
	var files []string
	fileerr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		return collectFilePaths(root, info, err, &files)
	})

	if fileerr != nil {
		fmt.Printf("Error walking the path %q: %v\n", root, fileerr)
		return nil, fileerr
	}
	if !f.IsDir() {
		fmt.Printf("Visited file or path: %q\n", path)
	}
	return files, nil
}

func GetMostRecentDirectory(root string) (string, error) {
	var dirs []os.FileInfo
	entries, err := ioutil.ReadDir(root)
	if err != nil {
		return "", err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry)
		}
	}

	sort.Slice(dirs, func(i, j int) bool {
		ti, _ := time.Parse("2006-01-02T15:04:05", dirs[i].Name())
		tj, _ := time.Parse("2006-01-02T15:04:05", dirs[j].Name())
		return ti.After(tj)
	})

	if len(dirs) > 0 {
		return filepath.Join(root, dirs[0].Name()), nil
	}

	return "", fmt.Errorf("no directories found")
}
