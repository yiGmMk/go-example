package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestExampleintHeap(t *testing.T) {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 2222)
	heap.Push(h, 6)
	heap.Push(h, 0)
	fmt.Printf("minimum: %d\n", (*h)[0])
	fmt.Println(h.Min())
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func TestWriteNumber(t *testing.T) {
	file, err := os.OpenFile("./number.txt", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()
	if err != nil {
		log.Println(err)
		t.Fail()
	}
	writer := bufio.NewWriter(file)
	// 10000,0000
	for i := 0; i < 10000; i++ {
		for j := 0; j < 10000; j++ {
			num, writeNum := 0, 0
			if i < 10 && j < 1 {
				num = rand.Int()%100000 + 200000

			} else {
				num = rand.Int() % 100000
			}
			if j == 10000 {
				writeNum, err = writer.WriteString(fmt.Sprintf("%d", num))
			} else {
				writeNum, err = writer.WriteString(fmt.Sprintf("%d,", num))
			}

			log.Println(writeNum, err)
		}
		writer.WriteString("\n")
	}
	err = writer.Flush()
	log.Println(err)
}

func TestSort(t *testing.T) {
	st := time.Now()
	res, err := HeapSortTop10("./number.txt")

	ed := time.Now()
	fmt.Println(res, err)
	useTime := ed.Sub(st)
	ns, s, m := useTime.Milliseconds(), useTime.Seconds(), useTime.Minutes()
	log.Println("use time", st.Sub(ed).Seconds(), "秒", ns, s, m)
}
