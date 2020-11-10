package main

func fact(num int) int {
	if num == 0 {
		return 1
	}
	return num * fact(num-1)
}

func main() {
	println(fact(3))
}
