package project

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// func Test_ExitCLI(t *testing.T) test not necessary -> integration test only

// func Test_createFrameworkMap(t *testing.T) test not necessary -> integration test only

// func Test_createDatabaseDriverMap(t *testing.T) test not necessary -> integration test only

// func Test_CreateMainFile(t *testing.T) test not necessary -> integration test only

func Test_createPath(t *testing.T) {
	pc := &ProjectConfig{}

	// Create a temporary directory for testing
	tempDir := t.TempDir()

	// Define test cases
	tests := []struct {
		pathToCreate string
		projectPath  string
		expectError  bool
	}{
		{"newdir", tempDir, false},      // Non-existent directory
		{"existingdir", tempDir, false}, // Already existing directory
	}

	// Create an existing directory for the test
	if err := os.Mkdir(filepath.Join(tempDir, "existingdir"), 0751); err != nil {
		t.Fatalf("Error setting up test: %v", err)
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("pathToCreate=%s", tt.pathToCreate), func(t *testing.T) {
			err := pc.createPath(tt.pathToCreate, tt.projectPath)
			if (err != nil) != tt.expectError {
				t.Errorf("createPath() error = %v, expectError %v", err, tt.expectError)
			}

			// Verify the directory was created if there was no error expected
			if !tt.expectError {
				if _, err := os.Stat(filepath.Join(tt.projectPath, tt.pathToCreate)); os.IsNotExist(err) {
					t.Errorf("Expected directory %s to be created", filepath.Join(tt.projectPath, tt.pathToCreate))
				}
			}
		})
	}
}

// func Test_createFileAndWriteTemplate(t *testing.T) test not necessary -> integration test only

func Test_IsValidWebFramework(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"chi", true},
		{"echo", true},
		{"fiber", true},
		{"gin", true},
		{"gorilla/mux", true},
		{"httprouter", true},
		{"standard-library", true},
		{"unknown-framework", false},
		{"", false},
		{"Chi", false}, // case-sensitive check
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsValidWebFramework(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidWebFramework(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}

func Test_IsValidDatabaseDriver(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"mysql", true},
		{"postgres", true},
		{"sqlite", true},
		{"mongo", true},
		{"none", true},
		{"unknown-driver", false},
		{"", false},
		{"Mysql", false}, // case-sensitive check
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := IsValidDatabaseDriver(tt.input)
			if result != tt.expected {
				t.Errorf("IsValidDatabaseDriver(%q) = %v; expected %v", tt.input, result, tt.expected)
			}
		})
	}
}
