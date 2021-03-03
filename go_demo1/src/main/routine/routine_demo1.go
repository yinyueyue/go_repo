package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() execute ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	cupNum := runtime.NumCPU() //系统cpu个数

	fmt.Println("cpuNum=", cupNum)

	//设置使用多个cpu,golang 1.8之后不需要设置
	runtime.GOMAXPROCS(cupNum - 1)
	go test() //使用go 开启协程

	for i := 0; i < 10; i++ {
		fmt.Println("main() execute ", strconv.Itoa(i))
		time.Sleep(time.Second)
	}

}
