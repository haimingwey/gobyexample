package main

import "fmt"

func sum(nums ...int) int {
	fmt.Println("nums:", nums)
	total := 0
	for _, i := range nums {
		total = total + i
	}
	fmt.Println("total:", total)
	return total
}
func main() {

	sum(1)
	sum(1, 2)
	sum(1, 2, 4)

	m := []int{1, 2, 3, 4, 5}
	sum(m...)
}
