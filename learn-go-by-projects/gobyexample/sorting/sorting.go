package main

import (
	"fmt"
	"sort"
)

func main() {

	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)
	fmt.Println("StringsAreSorted", sort.StringsAreSorted(strs))

	ints := []int{3, 5, 2, 56}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	fmt.Println("IntsAreSorted", sort.IntsAreSorted(ints))
}
