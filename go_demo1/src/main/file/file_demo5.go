package main

import (
	"io/ioutil"
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

}
