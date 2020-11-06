package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	fmt.Println(person{age: 10, name: "haimingwey"})
	fmt.Println(person{age: 10})
	fmt.Println(person{name: "jack"})
	// &前缀生成一个结构体指针
	fmt.Println(&person{name: "lucy", age: 10})

	s := person{name: "hell"}
	fmt.Println(s.name)
	p := &s
	p.age = 10
	fmt.Println(p)

}
