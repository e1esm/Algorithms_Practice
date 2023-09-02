package main

import "testing"

func TestGetParityStatus(t *testing.T) {
	table := []struct {
		arr            []int
		expectedResult string
	}{
		{
			arr:            []int{1, 2, -3},
			expectedResult: DEFEAT,
		},
		{
			arr:            []int{7, 11, 7},
			expectedResult: WIN,
		},
		{
			arr:            []int{6, -2, 0},
			expectedResult: WIN,
		},
	}

	for _, test := range table {
		result := GetParityStatus(test.arr)
		if test.expectedResult != result {
			t.Logf("Invalid received result. Want: %s, got: %s",
				test.expectedResult, result)
		}
	}
}
