package main

import "testing"

func TestReverseWords(t *testing.T){
	type test struct{
		input string
		expectedOutput string
	}
	table := []test{
			{input: "Let's take LeetCode contest",
			expectedOutput: "s'teL ekat edoCteeL tsetnoc",
		},
		{
			input: "God Ding",
			expectedOutput: "doG gniD",
		},
	}

	for _, test := range table{
		result := ReverseWords(test.input)
		if result != test.expectedOutput{
			t.Errorf("Invalid result. Want: %v, got: %v", test.expectedOutput, result)
		}
	}
}