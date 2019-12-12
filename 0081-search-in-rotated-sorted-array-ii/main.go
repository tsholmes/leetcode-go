package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0))
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3))
	fmt.Println(search([]int{1, 2, 3, 4, 1}, 3))
	fmt.Println(search([]int{1, 1, 1, 1, 1}, 3))
	fmt.Println(search([]int{1, 2, 2, 2, 0, 1, 1}, 0))
	fmt.Println(search([]int{3, 1}, 1))
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	off := sort.Search(len(nums), func(i int) bool {
		return nums[i] <= nums[0]
	}) % len(nums)
	min := nums[off]
	off2 := off
	for {
		if nums[(off2+len(nums)-1)%len(nums)] <= min {
			off2 = (off2 + len(nums) - 1) % len(nums)
			min = nums[off2%len(nums)]
			if off2 == off {
				break
			}
		} else {
			break
		}
	}
	idx := sort.Search(len(nums), func(i int) bool {
		return nums[(i+off2)%len(nums)] >= target
	})
	return nums[(idx+off2)%len(nums)] == target
}
