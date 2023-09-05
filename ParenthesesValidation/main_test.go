package main

import "testing"

func TestIsValid(t *testing.T) {
	table := []struct {
		parentheses    string
		expectedOutput bool
	}{
		{parentheses: "{[()]}", expectedOutput: true},
		{parentheses: "()", expectedOutput: true},
		{parentheses: "[()]", expectedOutput: true},
		{parentheses: "[}{]", expectedOutput: false},
		{parentheses: "", expectedOutput: false},
		{parentheses: "[{}{}[][][][][]([[{{[()]}}]])]", expectedOutput: true},
	}

	for _, test := range table {
		result := IsValid(test.parentheses, NewStack())
		if result != test.expectedOutput {
			t.Errorf("Invalid result. Want: %v, got: %v", test.expectedOutput, result)
		}
	}
}
