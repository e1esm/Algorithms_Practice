package main

import "testing"

func TestGetLongestWord(t *testing.T) {
	table := []struct {
		length         int
		sentence       string
		expectedLength int
		expectedWord   string
	}{
		{
			length:         19,
			sentence:       "i love segment tree",
			expectedWord:   "segment",
			expectedLength: 7,
		},
		{
			length:         21,
			sentence:       "from jumps from river",
			expectedLength: 5,
			expectedWord:   "jumps",
		},
	}

	for _, test := range table {
		word, length := GetLongestWord(test.sentence)
		if word != test.expectedWord || length != test.expectedLength {
			t.Log("Invalid result")
		}
	}
}
