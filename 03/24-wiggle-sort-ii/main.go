package main

import (
	"fmt"
	"sort"
)

func main() {
	do := func(nums []int) {
		wiggleSort(nums)
		fmt.Println(nums)
	}
	do([]int{1, 5, 1, 1, 6, 4})
	do([]int{1, 3, 2, 2, 3, 1})
	do([]int{1, 3, 2, 2, 3})
	do([]int{1, 2, 2, 3})
	do([]int{1, 2, 3, 3, 3, 5})
}

func wiggleSort(nums []int) {
	if len(nums) < 2 {
		return
	}
	sort.Ints(nums)
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i += 2 {
		res[i] = nums[(len(nums)+1)/2-1-i/2]
	}
	for i := 1; i < len(nums); i += 2 {
		res[i] = nums[len(nums)-i/2-1]
	}
	copy(nums, res)
}
