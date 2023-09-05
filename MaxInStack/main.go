package main

import (
	"fmt"
	"math"
)

type Stack struct {
	arr []int
	i   int
	max int
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]int, 0),
		max: math.MinInt,
		i:   -1,
	}
}

func (s *Stack) push(val int) {
	s.arr = append(s.arr, val)
	if val > s.max {
		s.max = val
	}
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
	s.setMax()
	return res
}

func (s *Stack) setMax() {
	index := s.i
	for index >= 0 {
		if s.arr[s.i] > s.max {
			s.max = s.arr[s.i]
		}
		index--
	}
}

func (s *Stack) Max() int {
	if s.i >= 0 {
		return s.max
	}
	return math.MinInt
}

func main() {
	stack := NewStack()
	stack.push(10)
	stack.push(5)
	stack.pop()
	stack.push(25)
	fmt.Println(stack.Max())
}
