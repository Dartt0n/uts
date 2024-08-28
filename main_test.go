package main

import (
	"os"
	"testing"
)

func TestReadFromArgs(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{
			name:     "No arguments",
			args:     []string{"uts"},
			expected: "",
		},
		{
			name:     "One argument",
			args:     []string{"uts", "1724692825"},
			expected: "1724692825",
		},
		{
			name:     "Multiple arguments",
			args:     []string{"uts", "1724692825", "extra"},
			expected: "1724692825",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			result := readFromArgs()
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}

func TestReadFromStdin(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: "",
		},
		{
			name:     "Single line input",
			input:    "1724692825",
			expected: "1724692825",
		},
		{
			name:     "Input with spaces and newlines",
			input:    "1724692825 \n",
			expected: "1724692825",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdin = r

			go func() {
				w.Write([]byte(tt.input))
				w.Close()
			}()

			result := readFromStdin()
			if result != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, result)
			}
		})
	}
}
