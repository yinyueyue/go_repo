package main

import (
	"fmt"
	"strconv"
)

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

	//会产生 deadlock
	/*num4 := <-intChan
	fmt.Println(num4)*/

	stringChan := make(chan string, 10)
	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i
		stringChan <- "hello " + strconv.Itoa(i)
	}

	/**

	 使用传统的for遍历管道时,如果不关闭管道会出现deadlock错误
	 在实际开发中,有时候不能确定什么时候应该关闭管道
	这个时候可以使用select case方式来进行管道的遍历
	*/

	var breakFlag bool
	for {
		select {
		case v := <-stringChan:
			{
				fmt.Printf("从stringChan读取到数据=%v\n", v)
			}

		case v := <-intChan2:
			{
				fmt.Printf("从intChan2读取到数据=%v\n", v)
			}
		default:
			fmt.Println("什么都没有取到,退出循环")
			breakFlag = true
		}
		if breakFlag {
			break
		}

	}

	//默认情况下,管道是双向的,可读可写
	//1.默认的读写管道声明方式: var chan1 chan int

	//2.声明为只写
	var chan2 chan<- int
	chan2 = make(chan int, 3)
	chan2 <- 20

	//3.声明为只读
	var chan3 <-chan int
	chan3 = make(chan int, 3)
	num5 := <-chan3
	fmt.Println(num5)
}
