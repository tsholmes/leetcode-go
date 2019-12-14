package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(target int, nums ...int) {
		fmt.Println(search(nums, target))
	}

	do(0, 4, 5, 6, 7, 0, 1, 2)
	do(3, 4, 5, 6, 7, 0, 1, 2)
	do(3, 0, 1, 2, 4, 5, 6, 7)
	do(1, 3, 1)
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	offset := sort.Search(len(nums), func(i int) bool { return nums[i] < nums[0] })
	pos := sort.Search(len(nums), func(i int) bool { return nums[(i+offset)%len(nums)] >= target })
	index := (pos + offset) % len(nums)
	if nums[index] == target {
		return index
	}
	return -1
}
