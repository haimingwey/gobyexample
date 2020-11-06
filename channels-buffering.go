package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "message 1"
	messages <- "message 2"
	messages <- "message 3"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}
