package main

import (
	"fmt"
	"time"
)

func main() {
	i := int64(0)
	cur := time.Now()
	for {
		fmt.Println(time.Now().Format("20060102-150405"), " ", i)
		i++
		if time.Now().Sub(cur) > time.Second {
			fmt.Println("stop after 1 minute", time.Now())
			break
		}
	}
}
