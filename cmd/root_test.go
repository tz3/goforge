// cmd/execute_test.go
package cmd

import (
	"bytes"
	"testing"
)

// TestExecute tests the Execute function
func TestExecute(t *testing.T) {
	tests := []struct {
		name        string
		setup       func()
		expectError bool
	}{
		{
			name: "Successful execution",
			setup: func() {
				rootCmd.SetArgs([]string{})
			},
			expectError: false,
		},
		{
			name: "Non-existing subcommand",
			setup: func() {
				rootCmd.SetArgs([]string{"nonexistent"})
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			tt.setup()

			// Capture output
			var buf bytes.Buffer
			rootCmd.SetOutput(&buf)

			// Execute
			err := rootCmd.Execute()
			if (err != nil) != tt.expectError {
				t.Fatalf("expected error: %v, got: %v", tt.expectError, err)
			}
		})
	}
}
