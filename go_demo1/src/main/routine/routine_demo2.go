package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	nums = make(map[int]int)
	lock = sync.Mutex{}
)

func main() {
	for i := 0; i < 100; i++ {
		go square(i)
	}

	time.Sleep(2000)
	fmt.Println(nums)
}

func square(num int) {
	time.Sleep(500)
	num2 := num * num
	lock.Lock()
	nums[num] = num2
	lock.Unlock()
}
