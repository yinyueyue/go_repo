package main

import (
	"flag"
	"fmt"
)

/**
使用flag包来解析命令行参数
*/
func main() {

	var user string
	var pwd string
	var host string
	var port int

	/**
	  p *string 使用什么指针对象接收
	  name string  对应命令行参数名称 -user
	  value string 默认值
	  usage string 参数说明
	*/
	flag.StringVar(&user, "user", "", "用户名，默认为空")
	flag.StringVar(&pwd, "pwd", "", "密码，默认为空")
	flag.StringVar(&host, "host", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 8080, "端口，默认为8080")

	//此方法必须调用
	flag.Parse()

	fmt.Printf("user=%v,pwd=%v,host=%v,port=%v", user, pwd, host, port)

	/**
	 执行方法
	 1.使用命令行编译
	  go build -o file_demo6.exe file_demo5.go

	2.运行exe文件
	file_demo6.exe -user root -pwd 12345 -host 127.0.0.1 -port 3306


	*/

}
