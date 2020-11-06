package main

import "fmt"

func main() {

	var a string = "121234"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// := 语法是声明并初始化变量的简写, 比如 var f string = "just test" 简写为如下；
	f := "just test"
	fmt.Println(f)
}
