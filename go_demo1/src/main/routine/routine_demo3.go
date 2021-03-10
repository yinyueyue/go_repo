package main

import (
	"fmt"
	"time"
)

func main() {
	exitChan := make(chan bool, 1)
	intChan := make(chan int, 100)
	go writeData(intChan)
	go readData(intChan, exitChan)
	for {
		_, ok := <-exitChan
		if ok {
			break
		}
	}
}

func writeData(intChan chan int) {
	for i := 0; i < 100; i++ {
		intChan <- i
		fmt.Printf("写入数据=%v\n", i)
		time.Sleep(100)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {

	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Printf("读取到数据=%v\n", v)
		time.Sleep(100)
	}
	exitChan <- true
	close(exitChan)
}
