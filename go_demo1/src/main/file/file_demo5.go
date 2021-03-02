package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

/**
读取文件写入到另外一个文件
*/
func main() {
	// 0666 文件模式，只有在linux和unix系统下有效，表示读取用户组等信息，windows系统下此参数无效
	//读写和追加模式
	sourceFile := "D://demo2.txt"
	targetFile := "D://demo3.txt"

	file, _ := ioutil.ReadFile(sourceFile)

	_ = ioutil.WriteFile(targetFile, file, 1)

	copyFile("D://demo1.txt", "D://demo4.txt")

}

/**
拷贝文件
*/
func copyFile(srcFilePath string, destFilePath string) (written int64, err error) {
	srcFile, err := os.Open(srcFilePath)

	if err != nil {
		fmt.Printf("open file error =%v", err)
		return
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 1)
	if err != nil {
		fmt.Printf("open file error =%v", err)
		return
	}
	defer destFile.Close()
	return io.Copy(bufio.NewWriter(destFile), bufio.NewReader(srcFile))
}
