package main

import "fmt"

//声明全局匿名函数
var (
	MultiplyFun = func(a, b int) int {
		return a * b
	}
)

//闭包 闭包内定义的变量会累计，只会在闭包声明时初始化执行一次
func AddUpper() func(int) int {
	var n int = 1
	return func(x int) int {
		n = n + x
		return n
	}
}

var map2 = make(map[int]int)

// 匿名函数demo
func main() {

	//声明匿名函数并立即调用
	res1 := func(a, b int) int {
		return a + b
	}(10, 20)
	fmt.Println(res1)

	//定义匿名函数，并赋值给一个对象
	subtractFun := func(a, b int) int {
		return a - b
	}

	fmt.Println(subtractFun(20, 30))

	fmt.Println(MultiplyFun(10, 6))

	f := AddUpper()
	fmt.Println(f(10))
	fmt.Println(f(2))
	fmt.Println(f(3))

}
