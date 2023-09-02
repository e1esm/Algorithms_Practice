package main

import (
	"reflect"
	"testing"
)

func Test_getDistanceArray(t *testing.T) {
	table := []struct {
		input  []int
		output []int
	}{
		{
			input:  []int{0, 1, 4, 9, 0},
			output: []int{0, 1, 2, 1, 0},
		},
		{
			input:  []int{0, 7, 9, 4, 8, 20},
			output: []int{0, 1, 2, 3, 4, 5},
		},
	}

	for _, test := range table {
		received := getDistanceArray(test.input)
		if !reflect.DeepEqual(test.output, received) {
			t.Errorf("Not equal arrays")
		}
	}
}
