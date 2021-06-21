package main

import "fmt"

type Pipe interface {
}

// 1. send num in list to chan
func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// 2. mul
func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// example_1 set up pipelines
func exam_1() {
	c := gen(2, 3, 4, 5, 6, 67, 78, 9, 20, 10)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
}
