package main

import "fmt"

const defaultHello = "Hello, World"
const englishHelloPrefix = "Hello, "
const chineseHelloPrefix = "你好, "
const frenchHelloPrefix = "Hola, "

func Hello(name string, language string) string {
	switch language {
	case "chinese":
		return chineseHelloPrefix + name
	case "english":
		return englishHelloPrefix + name
	case "french":
		return frenchHelloPrefix + name
	default:
		return defaultHello
	}
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
