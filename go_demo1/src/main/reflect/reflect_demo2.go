package main

import (
	"fmt"
	"reflect"
)

func main() {
	intNum := 10

	reflectPointer(&intNum)
	fmt.Printf("反射修改后的值=%v", intNum)

	reflectFloatDemo(26.888)

}

func reflectPointer(b interface{}) {

	val := reflect.ValueOf(b)

	//val的实际类型为指针，不能直接调用SetInt()方法
	fmt.Printf("val的类型=%v\n", val.Kind())

	/**
	Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	如果v的Kind不是Interface或Ptr会panic；
	如果v持有的值为nil，会返回Value零值。
	*/

	// 这系列Set方法只能使用指针变量的Elem()对象调用，否则会报错
	val.Elem().SetInt(20)
	fmt.Printf("val=%v\n", val.Elem().Int())

}

func reflectFloatDemo(floatNum float64) {

	floatVal := reflect.ValueOf(floatNum)

	fmt.Printf("value = %v\n", floatVal.Float())
	fmt.Printf("kind = %v\n", floatVal.Kind())
	fmt.Printf("type = %v\n", floatVal.Type())

	floatInterface := floatVal.Interface()

	fmt.Printf("interface = %v\n", floatInterface)

	//interface转float64

	//类型断言
	floatVal2 := floatInterface.(float64)

	fmt.Printf("floatVal2 = %v\n", floatVal2)

}
