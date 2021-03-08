package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

// 调用外部可执行文件,并将输出重定向到标准输出
func main() {
begin:
	cmd := exec.Command("./print/print")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Println(errors.WithStack(err))
		return
	}
	cur := time.Now()
	fmt.Println("==========================start success======================")
	for {
		if time.Now().Sub(cur) > time.Second*5 {
			log.Println("exit loop", time.Now())
			time.Sleep(time.Second * 2)
			log.Println("start")
			goto begin
		}
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(errors.WithStack(err))
	}
}
