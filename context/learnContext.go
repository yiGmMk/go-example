package test_context

import (
	"context"
	"fmt"
	"time"
)

func main() {
	Ctx()
}

func Ctx() {
	ctx := context.Background()
	cancelCtx, _ := context.WithTimeout(ctx, time.Minute)
	//defer cancel()
	valCtx := context.WithValue(ctx, "province", "上海")
	fmt.Println(valCtx.Value("province"))

	go func() {
		for {
			fmt.Println("enter goroutine")
			select {
			case <-cancelCtx.Done():
				return
			default:
				fmt.Println(time.Now().Format("2006-01-02 150405"))
			}
		}
	}()
}
