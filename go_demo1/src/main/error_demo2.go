package main

import (
	"errors"
	"fmt"
)

func main() {
	var a float64
	fmt.Println(a)
	//test01()
	getCircleArea(1)
	area2, err := getCircleArea2(-2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("area = %v", area2)
}
func test01() {

	//捕获异常
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到的错误：%s\n", r)
		}
	}()
	a := [5]int{0, 1, 2, 3, 4}
	a[1] = 123
	fmt.Println(a)
	//a[10] = 11
	index := 10
	a[index] = 10
	fmt.Println(a)

}

func getCircleArea(radius float32) (area float32) {
	if radius < 0 {
		//	return errors.New("半径不能是负数")
		panic("半径不能是负数")
	}

	return radius * radius * 3.14

}

func getCircleArea2(radius float32) (area float32, err error) {
	if radius < 0 {
		//手动返回错误信息
		return 0, errors.New("半径不能是负数")

	}

	return radius * radius * 3.14, nil

}
