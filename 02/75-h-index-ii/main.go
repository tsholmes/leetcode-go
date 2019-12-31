package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{0, 1, 3, 5, 6}))
	fmt.Println(hIndex([]int{5, 5, 5, 5, 5}))
	fmt.Println(hIndex([]int{4, 5, 5, 5, 5}))
}

func hIndex(citations []int) int {
	return len(citations) - sort.Search(len(citations), func(i int) bool {
		h := len(citations) - i
		return citations[i] >= h
	})
}
