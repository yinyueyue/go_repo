package main

import (
	"fmt"
	"reflect"
)

/**
golang 反射
*/

func main() {
	student := Student{
		Id:    1,
		Name:  "Mask",
		class: "class 1",
	}

	/**
	Type 是类型
	Kind 是类别
	*/
	typeA := reflect.TypeOf(student)
	v := reflect.ValueOf(student)

	//类型断言，interface转换为原始对象
	student1 := v.Interface().(Student)

	//不会改变student1
	student.Name = "Jack"

	fmt.Println(student1)
	fmt.Println(student1 == student)
	fmt.Println(v)
	fmt.Println(typeA.Name())
	fmt.Printf("typeA 的分类1=%v\n", typeA.Kind())
	fmt.Printf("typeA 的分类2=%v\n", v.Kind())
	fmt.Printf("typeA 的类别=%v\n", v.Type())

	/**
		kind()是对象的分类，返回是一个常量值，
	const (
		Invalid Kind = iota
		Bool
		Int
		Int8
		Int16
		Int32
		Int64
		Uint
		Uint8
		Uint16
		Uint32
		Uint64
		Uintptr
		Float32
		Float64
		Complex64
		Complex128
		Array
		Chan
		Func
		Interface
		Map
		Ptr
		Slice
		String
		Struct
		UnsafePointer
	)

	Type()是对象的具体类别，返回
	typeA 的类别=main.Student
	*/

	a := 10
	aType := reflect.TypeOf(a)
	aVal := reflect.ValueOf(a)
	fmt.Println(aType)
	fmt.Println(aVal.Int())
	aInterface := aVal.Interface()

	a += 2
	//类型断言
	a2, ok := aInterface.(int)
	if ok {
		fmt.Println(a2)
	}

	//常量定义
	const (
		b    = iota       // 每次 const 出现时，都会让 iota 初始化为0
		c                 // 1
		d                 // 2
		e, f = iota, iota //不换行，值相同
		//f = iota
	)
	fmt.Printf("b=%v,c=%v,d=%v,e=%v,f=%v\n", b, c, d, e, f)

}

type Student struct {
	Id    int
	Name  string
	class string
}
