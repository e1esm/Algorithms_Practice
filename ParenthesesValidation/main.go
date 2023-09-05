package main

import (
	"fmt"
	"time"
)

type Stack struct {
	arr []rune
	i   int
}

func NewStack() *Stack {
	return &Stack{
		arr: make([]rune, 0),
		i:   -1,
	}
}

func (s *Stack) Push(val rune) {
	s.arr = append(s.arr, val)
	s.i++
}

func (s *Stack) pop() rune {
	if s.IsEmpty() {
		return ' '
	}
	popped := s.arr[s.i]
	s.arr = s.arr[:s.i]
	s.i--
	return popped
}

func (s *Stack) IsEmpty() bool {
	if s.i == -1 {
		return true
	}
	return false
}

func (s *Stack) Peek() rune {
	return s.arr[s.i]
}

func IsValid(parentheses string, stack *Stack) bool {
	if parentheses == "" {
		return false
	}
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()
	for _, val := range parentheses {
		if val == '(' || val == '[' || val == '{' {
			stack.Push(val)
		} else {
			if getPair(stack.Peek()) != val {
				return false
			}
			stack.pop()
		}
	}
	if stack.IsEmpty() {
		return true
	} else {
		return false
	}
}

func getPair(part rune) rune {
	switch part {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	default:
		return ' '
	}
}
