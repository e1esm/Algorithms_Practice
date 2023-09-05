package main

import (
	"fmt"
	"testing"
)

func TestRotateMatrix(t *testing.T) {
	table := []struct {
		rows           int
		columns        int
		matrix         [][]int
		expectedMatrix string
	}{
		{
			rows:    4,
			columns: 3,
			matrix: [][]int{
				{1, 2, 3},
				{0, 2, 6},
				{7, 4, 1},
				{2, 7, 0},
			},
		},
		{
			rows:    9,
			columns: 5,
			matrix: [][]int{
				{-7, -1, 0, -4, -9},
				{5, -1, 2, 2, 9},
				{3, 1, -8, -1, -7},
				{9, 0, 8, -8, -1},
				{2, 4, 5, 2, 8},
				{-7, 10, 0, -4, -8},
				{-3, 10, -7, 10, 3},
				{1, 6, -7, -5, 9},
				{-1, 9, 9, 1, 9},
			},
		},
	}
	for _, test := range table {
		result := RotateMatrix(test.rows, test.columns, test.matrix)
		fmt.Println(result)
	}

}
