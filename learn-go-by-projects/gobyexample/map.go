package main

import "fmt"

func main() {

	m := make(map[string]int)
	m["k1"] = 1
	m["k2"] = 20

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("m", m)

	// prs 表示是否包含key k2
	_, prs := m["k2"]
	fmt.Println("has key k2 ?:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
