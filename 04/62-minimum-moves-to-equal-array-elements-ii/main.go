package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(minMoves2([]int{1, 0, 0, 8, 6}))
}

func minMoves2(nums []int) int {
	sort.Ints(nums)
	sums := make([]int, len(nums))
	for i, n := range nums {
		if i > 0 {
			sums[i] += sums[i-1]
		}
		sums[i] += n
	}

	min := math.MaxInt64

	for i := range nums {
		lc := i
		gc := len(nums) - i - 1

		ls := 0
		if i > 0 {
			ls = sums[i-1]
		}
		gs := sums[len(sums)-1] - sums[i]

		cost := (nums[i] * lc) - ls + gs - (nums[i] * gc)
		if cost < min {
			min = cost
		}
	}
	return min
}
