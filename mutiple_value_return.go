package main

import "fmt"

func get2Int() (int, int) {
	return 2, 3
}

func main() {

	i, j := get2Int()
	fmt.Println(i, j)
	_, k := get2Int()
	fmt.Println(k)

}
