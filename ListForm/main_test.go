package main

import (
	"reflect"
	"testing"
)

func TestResultOf(t *testing.T) {
	table := []struct {
		firstSlice     []int
		secondSlice    []int
		expectedOutput []int
	}{
		{
			firstSlice:     []int{1, 2, 0, 0},
			secondSlice:    []int{3, 4},
			expectedOutput: []int{1, 2, 3, 4},
		},
		{
			firstSlice:     []int{9, 5},
			secondSlice:    []int{1, 7},
			expectedOutput: []int{1, 1, 2},
		},
		{
			firstSlice:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1},
			secondSlice:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0, 9, 8, 7, 6, 5, 4, 3, 2},
			expectedOutput: []int{2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 3, 3, 3, 3, 3, 3, 3, 2, 3},
		},
	}
	for _, test := range table {
		result := ResultOf(test.firstSlice, test.secondSlice)
		if !reflect.DeepEqual(result, test.expectedOutput) {
			t.Errorf("Invalid resul. Want: %v, got: %v",
				test.expectedOutput, result)
		}
	}
}
