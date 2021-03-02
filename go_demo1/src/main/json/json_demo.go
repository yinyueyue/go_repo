package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	employee := Employee{
		Id:   100,
		Name: "Blade",
		age:  25,
		No:   "01234",
		Job:  "java engineer",
	}
	jsonBytes, _ := json.Marshal(employee)
	fmt.Printf("json序列化后=%v\n", string(jsonBytes))

	var employee2 Employee
	json.Unmarshal(jsonBytes, &employee2)

	fmt.Printf("json反序列化后=%v\n", employee2)

}

type Employee struct {
	Id   int `json:id`
	Name string
	age  int
	No   string
	Job  string `json:job`
}
