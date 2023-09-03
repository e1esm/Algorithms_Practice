package main

import (
	"testing"
)

func TestBinaryOf(t *testing.T) {
	table := []struct {
		input          int
		expectedOutput string
	}{
		{
			input:          5,
			expectedOutput: "101",
		},
		{
			input:          14,
			expectedOutput: "1110",
		},
		{
			input:          2543,
			expectedOutput: "100111101111",
		},
		{
			input:          2563478,
			expectedOutput: "1001110001110110010110",
		},
	}

	for _, test := range table {
		output := BinaryOf(test.input)
		if output != test.expectedOutput {
			t.Logf("Invalid result. Want: %v, got: %v",
				test.expectedOutput, output)
		}
	}
}
