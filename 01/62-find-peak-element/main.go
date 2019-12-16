package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 3, 1}))
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}

func findPeakElement(nums []int) int {
	return sort.Search(len(nums), func(i int) bool {
		return i == len(nums)-1 || nums[i] > nums[i+1]
	})
}
