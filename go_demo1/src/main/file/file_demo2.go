package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	//使用ioutil读取文件，一次性把文件加载到内存，适合读取小文件，大文件存在效率低和内存泄漏的风险
	//此方法内部存在关闭，不需要手动处理
	fileBytes, _ := ioutil.ReadFile("D://demo.txt")

	fmt.Printf("读取到的文件内容：%v", string(fileBytes))

}
