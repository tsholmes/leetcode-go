package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(target int, nums ...int) {
		fmt.Println(searchInsert(nums, target))
	}
	do(5, 1, 3, 5, 6)
	do(2, 1, 3, 5, 6)
	do(7, 1, 3, 5, 6)
	do(0, 1, 3, 5, 6)
}

func searchInsert(nums []int, target int) int {
	return sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
}
