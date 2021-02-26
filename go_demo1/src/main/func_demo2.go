package main

import (
	"fmt"
	"time"
)

//  通过匿名字段，可获得和继承类似的复用能力。依据编译器查找次序，只需在外层定义同名方法，就可以实现 "override"。

type Employee struct {
	id        int
	name      string
	age       int
	entryDate string
}

func (e Employee) showInfo() {

}

func showInfo(t *Employee) string {
	return fmt.Sprint("\n", &t)
}

type JavaEngineer struct {
	Employee
	//	hobby string
}

func (t JavaEngineer) testT() {
	fmt.Println("如类型 S 包含匿名字段 T，则 S 和 *S 方法集包含 T 方法。")
}

func main() {
	s1 := JavaEngineer{Employee{1, "jack", 20, time.Now().Format(time.RFC3339)}}
	s1.showInfo()
	s2 := &s1
	fmt.Printf("s1 is : %v\n", s1)
	s1.testT()
	fmt.Printf("s2 is : %v\n", s2)
	s2.testT()

}
