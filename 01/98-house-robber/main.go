package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	if dp[0] > dp[1] {
		dp[1] = dp[0]
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = dp[i-1]
		if s := dp[i-2] + nums[i]; s > dp[i] {
			dp[i] = s
		}
	}
	return dp[len(nums)-1]
}
