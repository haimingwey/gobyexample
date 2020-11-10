package main

import "fmt"

/**
TODO 理解闭包
把 “func() int" 当作一个整体，函数intSeq()返回的是一个函数，
这个函数是一个匿名函数（没有函数名的函数），并且这个函数的返回值是int。
*/

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	newSeq := intSeq()
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	fmt.Println(newSeq())
	anotherSeq := intSeq()
	println(anotherSeq())
}
