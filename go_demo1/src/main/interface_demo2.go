package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	name string
	age  int
}

type heroSlice []Hero

func (heroes heroSlice) Len() int {
	return len(heroes)
}

//指定排序规则，升序或者降序
func (heroes heroSlice) Less(i, j int) bool {
	return heroes[i].age < heroes[j].age
}

func (heroes heroSlice) Swap(i, j int) {

	//对象转换，下面使用简洁写法
	/*	 temp :=heroes[i]
		 heroes[i] = heroes[j]
		 heroes[j] = temp*/

	heroes[i], heroes[j] = heroes[j], heroes[i]
}

func main() {

	/**
	创建heroes切片，并对切片内的数据按照年龄排序，需要实现sort包内的type Interface interface接口
	*/
	var heroes heroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			//	name: "hero " + (string(rune(rand.Intn(100)))),
			name: fmt.Sprintf("英雄-%d", rand.Intn(100)),
			age:  rand.Intn(100),
		}
		heroes = append(heroes, hero)
	}

	//排序
	sort.Sort(heroes)
	fmt.Println("%V\n", heroes)

	//类型断言
	var a interface{}
	var point = Point{1, 2}
	a = point
	var b Point

	//这里使用类型断言来处理，把a转换为b，相当于java的强制类型转换
	//含义为：判断对象a是否是指向Point类型的变量，如果是就转换为Point类型，并赋值给b，否则会报错，抛出异常
	b = a.(Point)
	fmt.Println(b)

	var x interface{}
	var y float32 = 2.2
	x = y

	//类型转换，带转换结果，转换失败不会抛出异常
	y2, ok := x.(float32)

	if ok {
		fmt.Printf("转换成功，y2=%v的类型是%T", y2, y2)
	} else {
		fmt.Println("转换失败")
	}

	var c bool = true
	var c1 float32 = 9.9
	var c2 string = "haha"
	var c3 int32 = 333
	var c4 int64 = 666
	TypeJudge(c, c1, c2, c3, c4)

}

//可变参数，循环判断录入的参数类型， x.(type)为固定写法
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			{
				fmt.Printf("第%v个参数的类型是bool类型，值为%v\n", index, x)
				continue
			}
		case float32:
			{
				fmt.Printf("第%v个参数的类型是float32类型，值为%v\n", index, x)
				continue
			}
		case float64:
			{
				fmt.Printf("第%v个参数的类型是float64类型，值为%v\n", index, x)
				continue
			}
		case string:
			{
				fmt.Printf("第%v个参数的类型是string类型，值为%v\n", index, x)
				continue
			}
		case int32:
			{
				fmt.Printf("第%v个参数的类型是int32类型，值为%v\n", index, x)
				continue
			}
		default:
			fmt.Println("unknown type")
		}

	}
}

type Point struct {
	x int
	y int
}
