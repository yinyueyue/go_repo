package model

type Student struct {
	Id   int
	Name string
	Age  int
}

//如果Student为小写的student，在其他包内生成student对象需要使用工厂模式

func NewStudent(id int, name string, age int) *Student {

	return &Student{id, name, age}
}
