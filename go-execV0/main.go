package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"

	"github.com/pkg/errors"
)

func main() {
begin:
	cmd := exec.Command("./print/print", "-P stratum1+ssl://0x7ff2bbb1b5174587580fe92d13e5abc7c85ca67f@asia1.ethermine.org:5555")

	var out bytes.Buffer
	cmd.Stdout = &out //os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Println(errors.WithStack(err))
		return
	}
	cur := time.Now()
	restTime := time.Now()
	fmt.Println("==========================start success======================")
	for {
		if time.Now().Sub(restTime) > time.Second*1 {
			log.Println("--------------reset buffer-------------")
			out.Reset()
			restTime = time.Now()
			continue
		}
		if time.Now().Sub(cur) > time.Second*5 {
			log.Println("-------------------exit loop-----------", time.Now())
			time.Sleep(time.Second * 2)
			log.Println("start")
			goto begin
		}
		str, err := out.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				continue
			}
			log.Println("读取错误", errors.WithStack(err))
			continue
		}
		fmt.Printf("cap=" + fmt.Sprintf("%d", out.Cap()) + str)
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(errors.WithStack(err))
	}
}
