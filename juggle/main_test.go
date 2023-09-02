package main

import "testing"

func Test_GetPoints(t *testing.T) {
	table := []struct {
		triesPerPerson int
		table          [][]rune
		expectedPoints int
	}{
		{
			triesPerPerson: 3,
			table: [][]rune{
				{'1', '2', '3', '1'},
				{'2', '.', '.', '2'},
				{'2', '.', '.', '2'},
				{'2', '.', '.', '2'},
			},
			expectedPoints: 2,
		},
		{
			triesPerPerson: 4,
			table: [][]rune{
				{'1', '1', '1', '1'},
				{'9', '9', '9', '9'},
				{'1', '1', '1', '1'},
				{'9', '9', '1', '1'},
			},
			expectedPoints: 1,
		},
		{
			triesPerPerson: 4,
			table: [][]rune{
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
				{'1', '1', '1', '1'},
			},
		},
	}

	for _, test := range table {
		received := getPoints(test.triesPerPerson, test.table)
		if received != test.expectedPoints {
			t.Logf("Invalid amount of points received, want: %d, got: %d",
				test.expectedPoints,
				received)
		}
	}
}
