package main

import "testing"

func TestCommitsToDo(t *testing.T) {
	table := []struct {
		n               int
		expectedCommits int
	}{
		{
			n:               3,
			expectedCommits: 3,
		},
		{
			n:               0,
			expectedCommits: 1,
		},
		{
			n:               10,
			expectedCommits: 89,
		},
	}
	for _, test := range table {
		result := CommitsToDo(test.n)
		if result != test.expectedCommits {
			t.Errorf("Invalid result. Want: %d, got: %d", test.expectedCommits, result)
		}
	}
}
