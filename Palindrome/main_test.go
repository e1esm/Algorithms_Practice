package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	table := []struct {
		input          string
		expectedOutput bool
	}{
		{
			input:          "A man, a plan, a canal: Panama",
			expectedOutput: true,
		}, {
			input:          "zo",
			expectedOutput: false,
		},
	}

	for _, test := range table {
		output := IsPalindrome(test.input)
		if output != test.expectedOutput {
			t.Logf("Invalid output. Want: %v, got: %v",
				test.expectedOutput, output)
		}
	}
}
