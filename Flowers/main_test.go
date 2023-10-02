package main

import (
	"reflect"
	"testing"
)

func TestGetGarednBeds(t *testing.T){
	table := []struct{
		spreads []Spread
		amount int
		expectedOutput []Spread
	}{
		{
			spreads: []Spread{
				{start: 7, end: 8},
				{start: 7, end: 8},
				{start: 2, end: 3},
				{start: 6, end: 10},
			},
			amount: 4,
			expectedOutput: []Spread{
				{2, 3},
				{6, 10},
			},
		},
		{
			spreads: []Spread{
				{2, 3},
				{5, 6},
				{3, 4},
				{3, 4},
			},
			amount: 4,
			expectedOutput: []Spread{
				{2, 4},
				{5, 6},
			},
		},
		{
			spreads: []Spread{
				{1, 3},
				{3, 5},
				{4, 6},
				{5, 6},
				{2, 4},
				{7, 10},
			},
			amount: 6,
			expectedOutput: []Spread{
				{1, 6},
				{7, 10},
			},
		},
	}

	for _, v := range table{
		result := GetGardenBeds(v.amount, v.spreads)
		if !reflect.DeepEqual(result, v.expectedOutput){
			t.Errorf("Invalid result. Want: %v, got: %v",
			 v.expectedOutput,
			result)
		}
	}
}