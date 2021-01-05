package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 5, 6, 8, 10, 22, 34, 36, 55}
	index := BinarySearch(&arr, 6)
	fmt.Printf("找到目标元素，index=%v", index)

}

//二分查找 实现
func BinarySearch(arr *[]int, target int) int {
	length := len(*arr)
	left := 0
	right := length - 1
	for left <= right {
		mid := (left + right) / 2
		if (*arr)[mid] > target {
			right = mid - 1
		} else if (*arr)[mid] < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
