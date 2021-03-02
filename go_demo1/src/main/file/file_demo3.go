package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
写文件
*/
func main() {
	// 0666 文件模式，只有在linux和unix系统下有效，表示读取用户组等信息，windows系统下此参数无效
	file, err := os.OpenFile("D://demo2.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}

	str1 := "hello go lang"

	//使用带缓冲的读写
	writer := bufio.NewWriter(file)

	defer file.Close()
	for i := 0; i < 10; i++ {
		_, _ = writer.WriteString(str1 + strconv.Itoa(i) + "\r\n")
	}

	//刷新文件缓冲区到文件
	_ = writer.Flush()

}
