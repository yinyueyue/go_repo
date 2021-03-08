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

	/*num4 := <-intChan
	fmt.Println(num4)*/

	intChan <- 100
	intChan <- 200
	intChan <- 300
	intChan <- 400
	intChan <- 500

	//关闭管道， 通道关闭后不能再写入数据，但是可以继续读取
	close(intChan)

	//遍历管道，如果管道没有关闭，可以遍历完所有的数据，但是会报deadline错误

	for v := range intChan {

		fmt.Printf("v=%v", v)
	}

}
