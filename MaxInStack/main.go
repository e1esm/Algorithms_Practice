package main

import (
	"fmt"
	"math"
)

type Stack struct {
	arr []int
	i   int
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]int, 0),
		i:   -1,
	}
}

func (s *Stack) push(val int) {
	s.arr = append(s.arr, val)
	s.i++
}

func (s *Stack) pop() int {
	if s.i < 0 {
		fmt.Println("Error")
		return s.i
	}
	res := s.arr[s.i]
	s.arr = s.arr[:s.i]
	s.i--
	return res
}

func (s *Stack) Max() int {
	indCopy := s.i
	max := math.MinInt
	for indCopy >= 0 {
		if s.arr[indCopy] > max {
			max = s.arr[indCopy]
		}
		indCopy--
	}
	return max
}

func main() {
	stack := NewStack()
	stack.push(10)
	stack.push(5)
	stack.pop()
	stack.push(25)
	fmt.Println(stack.Max())
}
