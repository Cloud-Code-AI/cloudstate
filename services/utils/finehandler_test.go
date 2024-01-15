package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestWriteJSONToFile tests the WriteJSONToFile function.
func TestWriteJSONToFile(t *testing.T) {
	// Create a temporary directory for testing
	tempDir, err := ioutil.TempDir("", "filewriter_test")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after the test

	// Define the test file path
	testFilePath := filepath.Join(tempDir, "test.json")

	// Test data
	testData := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, world!",
	}

	// Call the function under test
	if err := WriteJSONToFile(testFilePath, testData); err != nil {
		t.Errorf("WriteJSONToFile failed: %v", err)
	}

	// Read the file back
	data, err := ioutil.ReadFile(testFilePath)
	if err != nil {
		t.Fatalf("Failed to read test file: %v", err)
	}

	// Unmarshal and check the data
	var readData struct {
		Message string `json:"message"`
	}
	if err := json.Unmarshal(data, &readData); err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Compare the data
	if readData.Message != testData.Message {
		t.Errorf("Expected message %q, got %q", testData.Message, readData.Message)
	}
}
