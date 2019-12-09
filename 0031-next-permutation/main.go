package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(nums ...int) {
		nextPermutation(nums)
		fmt.Println(nums)
	}

	do(1, 1)
	do(1, 2, 3)
	do(1, 3, 2)
	do(2, 3, 1)
	do(3, 2, 1)
	do(1, 1, 5)
	do(2, 4, 3, 1)
}

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	end := len(nums) - 1
	for end > 0 && nums[end] <= nums[end-1] {
		end--
	}
	if end == 0 {
		sort.Ints(nums)
		return
	}

	start := len(nums) - 1
	for nums[end-1] >= nums[start] {
		start--
	}
	nums[end-1], nums[start] = nums[start], nums[end-1]
	sort.Ints(nums[end:])
}
