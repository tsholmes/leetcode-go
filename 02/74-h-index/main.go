package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(hIndex([]int{3, 0, 6, 1, 5}))
}

func hIndex(citations []int) int {
	sort.Ints(citations)

	for i := range citations {
		h := len(citations) - i
		if citations[i] >= h && (i == 0 || citations[i-1] <= h) {
			return h
		}
	}

	return 0
}
