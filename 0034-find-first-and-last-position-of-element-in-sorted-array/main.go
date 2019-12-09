package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(target int, nums ...int) {
		fmt.Println(searchRange(nums, target))
	}

	do(8, 5, 7, 7, 8, 8, 10)
	do(6, 5, 7, 7, 8, 8, 10)
}

func searchRange(nums []int, target int) []int {
	left := sort.Search(len(nums), func(i int) bool { return nums[i] >= target })
	if left == len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	invRight := sort.Search(len(nums), func(i int) bool { return nums[len(nums)-i-1] <= target })
	right := len(nums) - invRight - 1
	return []int{left, right}
}
