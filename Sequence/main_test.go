package main

import "testing"



func TestIsSequence(t *testing.T){
	table := []struct{
			sequenceToFind string
			sentence string
			expectedOutput bool
	}{
		{
			sequenceToFind: "abc",
			sentence: "ahbgdcu",
			expectedOutput: true,
		},{
			sequenceToFind: "abcp",
			sentence: "ahcp",
			expectedOutput: false,
		},
	}

	for _, v := range table{
		result := IsSequence(v.sequenceToFind, v.sentence)
		if result != v.expectedOutput{
			t.Errorf("Invalid result. Want: %v, got: %v",
		v.expectedOutput,
		result)
		}
	}
}