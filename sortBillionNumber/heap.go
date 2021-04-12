package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Min() int {
	if len(*h) <= 0 {
		return 0
	}
	return (*h)[0]
}

func (h *IntHeap) Number() []int {
	return *h
}

func toArray(s string) ([]int, error) {
	num := strings.Split(s, ",")
	res := []int{}
	for _, n := range num {
		i, err := strconv.Atoi(n)
		if err != nil {
			fmt.Println(err)
			continue
		}
		res = append(res, i)
	}
	return res, nil
}

func HeapSortTop10(filePath string) ([]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)
	intHeap := &IntHeap{}
	heap.Init(intHeap)
	for {
		data, _ := reader.ReadString('\n')
		trim := strings.TrimSpace(data)
		if len(trim) <= 0 {
			break
		}
		array, err := toArray(trim)
		if err != nil {
			return nil, err
		}
		for _, n := range array {
			//fmt.Println("cur:", n, "min:", intHeap.Min())
			if n < intHeap.Min() && intHeap.Len() >= 10 {
				continue
			}
			heap.Push(intHeap, n)
			if intHeap.Len() > 10 {
				heap.Pop(intHeap)
			}
		}
	}
	return intHeap.Number(), nil
}
