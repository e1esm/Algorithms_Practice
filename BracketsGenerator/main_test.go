package main

import (
	"reflect"
	"testing"
)

func TestGenerateBrackets(t *testing.T) {
	table := []struct {
		n              int
		expectedOutput []string
	}{
		{
			n: 3,
			expectedOutput: []string{
				"((()))",
				"(()())",
				"(())()",
				"()(())",
				"()()()",
			},
		},
		{
			n: 2,
			expectedOutput: []string{
				"(())",
				"()()",
			},
		},
	}

	for _, v := range table {
		result := GenerateBrackets(v.n)
		if !reflect.DeepEqual(result, v.expectedOutput) {
			t.Errorf("Invalid result. Want: %v, got: %v",
				v.expectedOutput,
				result)
		}
	}
}
