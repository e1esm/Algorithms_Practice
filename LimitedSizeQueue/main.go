package main

import "fmt"

type Queue struct {
	arr   []int
	i     int
	limit int
}

func NewQueue(size int) *Queue {
	return &Queue{
		arr:   make([]int, 0),
		i:     -1,
		limit: size,
	}
}

func (q *Queue) push(value int) error {
	if q.arr == nil {
		q.arr = make([]int, q.limit)
	}
	if q.i+1 >= q.limit {
		return fmt.Errorf("Limit's met")
	}
	q.arr = append(q.arr, value)
	q.i++
	return nil
}

func (q *Queue) peek() (*int, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}
	return &q.arr[q.i], nil
}

func (q *Queue) pop() (*int, error) {
	if q.IsEmpty() {
		return nil, fmt.Errorf("Queue is empty")
	}
	var popped int
	if q.i > 0 {
		popped = q.arr[0]
		q.arr = q.arr[1:]
	}
	if q.i == 0 {
		popped = q.arr[0]
		q.arr = nil
	}
	q.i--
	return &popped, nil
}

func (q *Queue) IsEmpty() bool {
	return q.i == -1
}
