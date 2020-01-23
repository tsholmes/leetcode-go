package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}

func thirdMax(nums []int) int {
	sort.Ints(nums)

	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[count-1] {
			nums[count] = nums[i]
			count++
		}
	}
	nums = nums[:count]

	if len(nums) < 3 {
		return nums[len(nums)-1]
	}
	return nums[len(nums)-3]
}
