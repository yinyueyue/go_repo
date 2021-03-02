package main

import (
	"factory/model"
	"fmt"
)

func main() {

	var student = model.Student{1, "abc", 20}
	fmt.Println(student)

	var student2 = model.NewStudent(10, "haha", 30)
	fmt.Println(*student2)

	//这里可以直接使用student2来取值 ，编译器进行了优化，相当于（*student2）
	fmt.Printf("id=%v,name=%v,age=%v", student2.Id, (*student2).Name, student2.Age)
}
