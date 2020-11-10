package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan string, 1)

	select {
	case msg := <-messages:
		fmt.Println("received messages", msg)
	default:
		fmt.Println("no messages received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent messages", msg)
	default:
		fmt.Println("no messages sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message ", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
