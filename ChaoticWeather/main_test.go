package main

import (
	"testing"
)

func TestGetChaoticLength(t *testing.T) {
	table := []struct {
		length         int
		arr            []int
		expectedOutput int
	}{
		{
			length:         7,
			arr:            []int{-1, 10, -8, 0, 2, 0, 5},
			expectedOutput: 3,
		},
		{
			length:         5,
			arr:            []int{1, 2, 5, 4, 8},
			expectedOutput: 2,
		},
	}

	for _, test := range table {
		output := GetChaoticLength(test.arr)
		if output != test.expectedOutput {
			t.Errorf("Invalid return, want: %d got: %d",
				test.expectedOutput, output)
		}
	}
}
