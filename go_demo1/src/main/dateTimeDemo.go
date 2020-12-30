package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("now = %v and now type = %T\n", now, now)

	fmt.Printf("year = %v and month = %v and day = %v and hour = %v and min = %v"+
		" and sec = %v\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	//时间格式化
	fmt.Printf("Now is %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	nowStr := fmt.Sprintf("Now is %d-%d-%d %d:%d:%d\n", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())
	fmt.Println(nowStr)

	// 格式化方式2，
	// 2006-01-02 15:04:05 为固定值，可以自由组合，使用方法类似yyyyMMdd这种
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format(time.RFC3339))
	//unix 秒数时间戳
	fmt.Println(now.Unix())
	//unix 纳秒时间戳
	fmt.Println(now.UnixNano())
}
