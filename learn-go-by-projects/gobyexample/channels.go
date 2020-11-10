package main

import "fmt"

func main() {

	// 通道是连接多个协程的管道， 用make(chan type)创建通道
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

}
