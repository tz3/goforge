package cmd

import (
	"os"
	"testing"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
)

func TestNonInteractiveCommand(t *testing.T) {
	tests := []struct {
		name        string
		flagSetup   func() *pflag.FlagSet
		expectedCmd string
	}{
		{
			name: "No flags set",
			flagSetup: func() *pflag.FlagSet {
				return pflag.NewFlagSet("test", pflag.ContinueOnError)
			},
			expectedCmd: defaultProjectTitle,
		},
		{
			name: "One flag set",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "config.yaml", "config file")
				_ = fs.Set("config", "config.yaml")
				return fs
			},
			expectedCmd: "goforge --config config.yaml",
		},
		{
			name: "Multiple flags set",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "config.yaml", "config file")
				_ = fs.Set("config", "config.yaml")
				fs.Int("port", 8080, "port number")
				_ = fs.Set("port", "8080")
				return fs
			},
			expectedCmd: "goforge --config config.yaml --port 8080",
		},
		{
			name: "Help flag ignored",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "config.yaml", "config file")
				_ = fs.Set("config", "config.yaml")
				fs.Bool("help", false, "help flag")
				_ = fs.Set("help", "true")
				return fs
			},
			expectedCmd: "goforge --config config.yaml",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := tt.flagSetup()
			result := nonInteractiveCommand(flagSet)
			assert.Equal(t, tt.expectedCmd, result)
		})
	}
}

func TestHasChangedFlag(t *testing.T) {
	tests := []struct {
		name        string
		flagSetup   func() *pflag.FlagSet
		expectedRes bool
	}{
		{
			name: "No flags set",
			flagSetup: func() *pflag.FlagSet {
				return pflag.NewFlagSet("test", pflag.ContinueOnError)
			},
			expectedRes: false,
		},
		{
			name: "One flag set",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "", "config file")
				_ = fs.Set("config", "config.yaml")
				return fs
			},
			expectedRes: true,
		},
		{
			name: "Multiple flags set",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "", "config file")
				_ = fs.Set("config", "config.yaml")
				fs.Int("port", 0, "port number")
				_ = fs.Set("port", "8080")
				return fs
			},
			expectedRes: true,
		},
		{
			name: "Flags set to default values",
			flagSetup: func() *pflag.FlagSet {
				fs := pflag.NewFlagSet("test", pflag.ContinueOnError)
				fs.String("config", "config.yaml", "config file")
				fs.Int("port", 8080, "port number")
				// No Set() calls here, so no flags are actually changed from their default values
				return fs
			},
			expectedRes: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := tt.flagSetup()
			result := hasChangedFlag(flagSet)
			if result != tt.expectedRes {
				t.Errorf("expected %v, got %v", tt.expectedRes, result)
			}
		})
	}
}

func TestIsDirectoryNonEmpty(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() string
		expectedRes bool
	}{
		{
			name: "Directory does not exist",
			setup: func() string {
				return "nonexistent_dir"
			},
			expectedRes: false,
		},
		{
			name: "Directory exists but is empty",
			setup: func() string {
				dir, err := os.MkdirTemp("", "empty_dir")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}
				return dir
			},
			expectedRes: false,
		},
		{
			name: "Directory exists and is non-empty",
			setup: func() string {
				dir, err := os.MkdirTemp("", "non_empty_dir")
				if err != nil {
					t.Fatalf("failed to create temp dir: %v", err)
				}
				// Create a temporary file inside the directory
				file, err := os.CreateTemp(dir, "temp_file")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				file.Close()
				return dir
			},
			expectedRes: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setup()
			defer os.RemoveAll(dir) // Cleanup after the test
			result := isDirectoryNonEmpty(dir)
			if result != tt.expectedRes {
				t.Errorf("expected %v, got %v", tt.expectedRes, result)
			}
		})
	}
}

func TestIsValidProjectName(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "Test with Capital alphanumeric characters",
			arg:  "TEST",
			want: true,
		},
		{
			name: "Test with small alphanumeric characters",
			arg:  "test",
			want: true,
		},
		{
			name: "Test with numeric characters",
			arg:  "1234",
			want: true,
		},
		{
			name: "Test with mixed alphanumeric and numeric characters",
			arg:  "Test123",
			want: true,
		},
		{
			name: "Test with non-alphanumeric characters",
			arg:  "Test@123",
			want: false,
		},
		{
			name: "Test with empty string",
			arg:  "",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidProjectName(tt.arg); got != tt.want {
				t.Errorf("isValidInput() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestValidateFlags(t *testing.T) test not necessary -> integration test only

// func TestHandleInteractiveProjectName(t *testing.T) test not necessary -> integration test only

// func TestHandleInteractiveProjectType(t *testing.T) test not necessary -> integration test only

// func TestSetupProject(t *testing.T) test not necessary -> integration test only

// func TestSetFlagValue(t *testing.T) test not necessary -> integration test only
