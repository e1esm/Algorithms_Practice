package main

func bracketsGenerator(n, open, close int, str string, arr *[]string) {

	if open == n && close == n {
		*arr = append(*arr, str)
	}

	if open < n {
		bracketsGenerator(n, open+1, close, str+"(", arr)
	}
	if open > close {
		bracketsGenerator(n, open, close+1, str+")", arr)
	}
}

func GenerateBrackets(amount int) []string {
	arr := make([]string, 0)
	bracketsGenerator(amount, 0, 0, "", &arr)
	return arr
}
