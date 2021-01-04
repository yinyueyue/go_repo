package main

import "fmt"

func main() {
	arr := []int{0, 10, 20, 5, 89, 33, 55, 7, 8, 1}
	BubbleSort(&arr)
	fmt.Printf("冒泡排序后的结果：%v", arr)
}

// 冒泡排序算法 go实现
func BubbleSort(arr *[]int) {
	var tmp int
	LENGTH := len(*arr)
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				tmp = (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = tmp
			}
		}
	}
}
