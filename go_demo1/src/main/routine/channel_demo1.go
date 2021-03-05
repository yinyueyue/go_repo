package main

import "fmt"

/**
管道的创建和数据读取
*/
func main() {

	var intChan chan int
	//创建一个可以存放10个元素，且存放数据类型为int的管道
	intChan = make(chan int, 10)

	fmt.Printf("intChan 的值=v%,intChan的地址=%p\n", intChan, intChan)

	//向管道写数据

	intChan <- 10

	num := 20
	intChan <- num

	//获取管道的容量和长度

	fmt.Printf("intChan的容量为=%v,intChan的长度为=%v\n", cap(intChan), len(intChan))

	//读取管道中的数据，读取完了不能重复读，否则会报deadlock错误
	num2 := <-intChan
	num3 := <-intChan

	fmt.Printf("num2=%v,num3=%v\n", num2, num3)

	num4 := <-intChan
	fmt.Println(num4)
}
