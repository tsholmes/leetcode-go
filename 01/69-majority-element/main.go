package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}

func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
