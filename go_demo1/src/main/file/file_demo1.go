package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("D://demo.txt")
	if err != nil {
		fmt.Println("open file error", err)
	}

	//当函数退出时调用，关闭io流，避免内存溢出
	defer file.Close()

	fmt.Printf("file = %v\n", file)

	//将文件读取到缓冲区
	reader := bufio.NewReader(file)
	//方式1
	/*	for {
		line, err2 := reader.ReadString('\n')

		//读到文件末尾
		if err2 == io.EOF {
			break
		}
		fmt.Printf("读取到行：%v",line)
	}*/

	//方式2
	for {
		lineBytes, _, err2 := reader.ReadLine()

		//读到文件末尾
		if err2 == io.EOF {
			break
		}
		fmt.Printf("读取到行：%v\n", string(lineBytes))
	}

}
