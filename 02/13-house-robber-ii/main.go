package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 3, 2}))
	fmt.Println(rob([]int{1, 2, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) < 3 {
		max := 0
		for _, n := range nums {
			if n > max {
				max = n
			}
		}
		return max
	}
	N := len(nums)
	max := nums[0]

	// rob 0
	dp := make([]int, N)
	dp[0] = nums[0]
	dp[1] = nums[0]
	for i := 2; i < N-1; i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+nums[i] > dp[i] {
			dp[i] = dp[i-2] + nums[i]
		}
	}
	if dp[N-2] > max {
		max = dp[N-2]
	}

	// don't rob 0
	dp = make([]int, N)
	dp[0] = 0
	dp[1] = nums[1]
	for i := 2; i < N; i++ {
		dp[i] = dp[i-1]
		if dp[i-2]+nums[i] > dp[i] {
			dp[i] = dp[i-2] + nums[i]
		}
	}
	if dp[N-1] > max {
		max = dp[N-1]
	}

	return max
}
