package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumGap([]int{3, 6, 9, 1}))
	fmt.Println(maximumGap([]int{10}))
}

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort.Ints(nums)
	max := 0
	for i := 1; i < len(nums); i++ {
		gap := nums[i] - nums[i-1]
		if gap > max {
			max = gap
		}
	}
	return max
}
