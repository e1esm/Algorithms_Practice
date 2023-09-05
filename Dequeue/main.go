package main

import (
	"fmt"
	"math"
)

type Dequeue struct {
	arr []int
}

func NewDequeue() *Dequeue {
	return &Dequeue{
		arr: make([]int, 0),
	}
}

func (d *Dequeue) push_back(value int) {
	d.arr = append(d.arr, value)
}

func (d *Dequeue) pop_back() int {
	result := d.arr[len(d.arr)-1]
	d.arr = d.arr[:len(d.arr)-1]
	return result
}

func (d *Dequeue) push_front(value int) {
	d.arr = append(d.arr, 0)
	for i := len(d.arr) - 1; i > 0; i-- {
		d.arr[i], d.arr[i-1] = d.arr[i-1], d.arr[i]
	}
	d.arr[0] = value
}

func (d *Dequeue) pop_front() int {
	if len(d.arr) == 0 {
		return math.MaxInt
	}
	res := d.arr[0]
	d.arr = d.arr[1:]
	return res
}

func main() {
	dequeue := NewDequeue()
	dequeue.push_back(10)
	dequeue.push_back(23)
	dequeue.push_front(1)
	fmt.Println(dequeue.pop_front())
	fmt.Println(dequeue.pop_back())
}
