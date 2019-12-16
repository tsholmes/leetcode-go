package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{1, 3, 5}))
	fmt.Println(findMin([]int{2, 2, 2, 0, 1}))
	fmt.Println(findMin([]int{1, 1, 2, 3, 4, 1, 1}))
	fmt.Println(findMin([]int{3, 1, 3, 3}))
}

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
