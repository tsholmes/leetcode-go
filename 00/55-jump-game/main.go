package main

import "fmt"

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < i {
			return false
		}
		if i+nums[i] > max {
			max = i + nums[i]
		}
	}
	return max >= len(nums)-1
}
