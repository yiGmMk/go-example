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

//
func mergeDone(done <-chan struct{}, cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}
	wg.Add(len(cs))
	for _, in := range cs {
		go output(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func sqDone(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func FanOutStopReceiving() {
	done := make(chan struct{})
	defer close(done)

	in := gen(1, 23, 4, 5, 8, 434, 12, 16, 23, 4, 5, 8, 67, 132)
	c1 := sqDone(done, in)
	c2 := sqDone(done, in)

	// can't for range out => error
	//
	out := mergeDone(done, c1, c2)
	fmt.Println(<-out)
}

//  => fan-in
func FanIn() {

}
