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
}
