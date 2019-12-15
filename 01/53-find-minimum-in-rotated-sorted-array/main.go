package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{1, 2, 3, 4, 5}))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	offset := sort.Search(len(nums), func(i int) bool { return nums[i] < nums[0] }) % len(nums)
	return nums[offset]
}
