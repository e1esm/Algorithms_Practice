package main

import "strings"

func ReverseWords(s string) string {
	var result strings.Builder
	var word strings.Builder
	for i := 0; i < len(s); i++{
		if s[i] != 32{
			word.WriteByte(s[i])
		}
		if s[i] == 32 || i == len(s) - 1{
			CopyReversedWord(&result, &word)
			word.Reset()
		}
	}

	return strings.TrimSuffix(result.String(), " ")
}


func CopyReversedWord(target *strings.Builder, initial *strings.Builder){
	str := initial.String()
	for i := initial.Len() - 1; i >= 0; i--{
		target.WriteByte(str[i])
	}
	target.WriteString(" ")
}
