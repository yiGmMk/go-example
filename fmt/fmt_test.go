package test_fmt

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

// args没有按照与格式--对应输出,err占位了
func ErrorfMiss(err error, format string, args ...interface{}) {
	fmt.Printf(" ERROR:%v"+format, err, args)
}

func CorrectErrorf(err error, format string, args ...interface{}) {
	args = append(args, err)
	fmt.Printf(format+" ERROR:%v", args...)
}

func TestFmt(t *testing.T) {
	valStruct := struct {
		Name        string
		PhoneNumber string
	}{
		Name:        "张三",
		PhoneNumber: "18812345678",
	}
	ErrorfMiss(errors.New("输出错误(missing)"), "提示%d,%v", 12345, valStruct)
	fmt.Printf("\n------------------------\n")
	CorrectErrorf(errors.New("输出错误(normal) "), "提示%d,%v", 12345, valStruct)
}
