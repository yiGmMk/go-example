package main

import (
	"fmt"
	"sync"
)

func merge(in ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int, 0)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(in))
	for _, i := range in {
		go output(i)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// mutiple read single write => fan-out
func FanOut() {
	in := gen(1, 23, 4, 5, 8, 434, 12, 16, 23, 4, 5, 8, 67, 132)
	c1 := sq(in)
	c2 := sq(in)
	for n := range merge(c1, c2) {
		fmt.Println("fan-out", n)
	}
}
