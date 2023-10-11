package main

import "testing"

func TestGetSortedWardrobe(t *testing.T) {
	table := []struct {
		arr            []int
		expectedOutput string
	}{
		{
			arr:            []int{0, 2, 1, 2, 0, 0, 1},
			expectedOutput: "0 0 0 1 1 2 2",
		},
		{
			arr:            []int{2, 1, 2, 0, 1},
			expectedOutput: "0 1 1 2 2",
		},
		{
			arr:            []int{2, 1, 1, 2, 0, 2},
			expectedOutput: "0 1 1 2 2 2",
		},
	}
	for _, test := range table {
		result := GetSortedWardrobe(test.arr)
		if result != test.expectedOutput {
			t.Errorf("Invalid result. Got: %v, want: %v",
				result,
				test.expectedOutput)
		}
	}
}
