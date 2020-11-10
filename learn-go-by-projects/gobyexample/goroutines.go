package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	go f("goroutine")

	// 匿名函数
	go func(str string) {
		fmt.Println(str)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
