package main

import "testing"

func TestGetDaysIndexes(t *testing.T) {
	table := []struct {
		n              int
		arr            []int
		price          int
		firstExpected  int
		secondExpected int
	}{
		{
			n:              6,
			arr:            []int{1, 2, 4, 4, 6, 8},
			price:          3,
			firstExpected:  3,
			secondExpected: 5,
		},
		{
			n:              6,
			arr:            []int{1, 2, 4, 4, 4, 4},
			price:          3,
			firstExpected:  3,
			secondExpected: -1,
		},
		{
			n:              6,
			arr:            []int{1, 2, 4, 4, 4, 4},
			price:          10,
			firstExpected:  -1,
			secondExpected: -1,
		},
	}

	for _, v := range table {
		first, second := GetDaysIndexes(v.arr, v.price)
		if first != v.firstExpected || second != v.secondExpected {
			t.Errorf("Invalid returned values. Want: %v, %v. Got: %v, %v",
				v.firstExpected,
				v.secondExpected,
				first,
				second)
		}
	}
}
