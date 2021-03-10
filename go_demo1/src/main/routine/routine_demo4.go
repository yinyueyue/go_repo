package main

import "fmt"

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)
	exitChan := make(chan bool, 5)

	//定义匿名函数存入数据
	go putNum(1000, intChan)

	for i := 0; i < 5; i++ {
		go primeNum(intChan, exitChan, primeChan)
	}

	//遍历退出标识，判断是否全部退出
	go func() {
		for i := 0; i < 5; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	var primeNumSlice []int
	for {
		v, ok := <-primeChan
		if !ok {
			break
		}
		primeNumSlice = append(primeNumSlice, v)

	}
	fmt.Printf("所有的素数=%v", primeNumSlice)

}
func putNum(endNum int, intChan chan int) {
	for i := 1; i <= endNum; i++ {
		intChan <- i
	}
	close(intChan)
}

//计算是否是素数
func primeNum(intChan chan int, exitChan chan bool, primeChan chan int) {

	//捕获异常，协程如果出现异常，会导致其他协程也中断运行
	defer func() {

		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	for {

		isPrimeNum := true
		v, ok := <-intChan
		if !ok {
			exitChan <- true
			break
		}

		if v == 1 {
			continue
		}
		for i := 2; i <= v/2; i++ {
			if v%i == 0 {
				isPrimeNum = false
				break
			}
		}
		if isPrimeNum {
			primeChan <- v
		}

	}
}
