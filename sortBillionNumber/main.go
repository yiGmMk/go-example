package main

import "fmt"

func main() {
	file := "./number.txt"
	top, err := HeapSortTop10(file)
	fmt.Println(top, err)
}
