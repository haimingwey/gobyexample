package main

import (
	"fmt"
	"sort"
)

type sortByLen []string

func (s sortByLen) Len() int {
	return len(s)
}
func (s sortByLen) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortByLen) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {

	fruits := []string{"apple", "kiwi", "banana", "peach"}
	sort.Sort(sortByLen(fruits))
	fmt.Println(fruits)
}
