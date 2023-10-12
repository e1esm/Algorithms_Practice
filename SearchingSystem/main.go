package main

import (
	"fmt"
	"math"
	"strings"
)

const (
	q = 251
	r = 181
)

func GetIDsOfRelevantDocuments(arr []string, requests []string) []int {
	result := make([]int, 0)
	documentWords := make(map[int][]int)
	requestWords := make(map[int][]int)

	for i := 0; i < len(arr); i++ {
		documentWords[i] = generateHash(strings.Split(arr[i], " "))
	}
	for i := 0; i < len(requests); i++ {
		requestWords[i] = generateHash(strings.Split(requests[i], " "))
	}

	for i := 0; i < len(requests); i++ {
		for j := 0; j < len(documentWords); j++ {
			index := findElement(j+1, requestWords[i], documentWords[j])
			if index > -1 {
				result = append(result, index)
			}
		}
	}
	return result
}

func findElement(currIndex int, source, targets []int) int {

	for i := 0; i < len(targets); i++ {
		for j := 0; j < len(source); j++ {
			if source[j] == targets[i] {
				return currIndex
			}
		}
	}

	return -1
}

func generateHash(requestedStr []string) []int {
	hashedValues := make([]int, len(requestedStr))

	for i := 0; i < len(requestedStr); i++ {
		for j := 0; j < len(requestedStr[i]); j++ {
			hashedValues[i] += int(requestedStr[i][j]) * int(math.Pow(q, float64(len(requestedStr[i])-j-1)))
		}
		hashedValues[i] %= r
	}

	return hashedValues
}

func main() {
	fmt.Println(GetIDsOfRelevantDocuments([]string{
		"i love coffee",
		"coffee with milk and sugar",
		"free tea for everyone",
	}, []string{
		"i like black coffee without milk",
		"everyone loves new year",
		"mary likes black coffee without milk",
	}))
}
