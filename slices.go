package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("after set:", s)

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")

	fmt.Println("after append", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy s to c, c:", c)

	fmt.Println(s[3])
	fmt.Println(s[3:4])
	fmt.Println(s[:5])

	t := []int{1, 2, 3, 4}
	fmt.Println(t)

	twoD := make([][]int, 3)

	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD", twoD)
}
