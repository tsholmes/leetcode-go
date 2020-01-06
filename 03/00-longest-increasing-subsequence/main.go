package main

func main() {

}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	max := 0
	for _, n := range dp {
		if n > max {
			max = n
		}
	}
	return max
}
