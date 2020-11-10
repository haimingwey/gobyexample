package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	sum := 0
	for _, i := range a {
		fmt.Println("value in array:", i)
		sum = sum + i
	}
	fmt.Println("sum:", sum)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for key, value := range m {
		fmt.Printf("key:%s -> value:%d \n", key, value)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
