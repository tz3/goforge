package cmd

import (
	"testing"
)

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
