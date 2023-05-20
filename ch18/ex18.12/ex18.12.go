package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 6, 1, 3, 4}
	sort.Ints(s)
	fmt.Println(s)
}
