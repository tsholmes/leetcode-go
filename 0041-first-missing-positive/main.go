package main

import "fmt"

func main() {
	fmt.Println(firstMissingPositive([]int{1, 2, 0}))
	fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	fmt.Println(firstMissingPositive([]int{7, 8, 9, 11, 12}))
	fmt.Println(firstMissingPositive([]int{3, 2, 1}))
	fmt.Println(firstMissingPositive([]int{1, 1}))
}

func firstMissingPositive(nums []int) int {
	for i := range nums {
		for {
			n := nums[i]
			if n <= 0 || n > len(nums) || n == i+1 || nums[n-1] == n {
				break
			}
			nums[i], nums[n-1] = nums[n-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
