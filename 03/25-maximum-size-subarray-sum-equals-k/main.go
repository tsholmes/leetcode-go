package main

import "fmt"

func main() {
	fmt.Println(maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3))
	fmt.Println(maxSubArrayLen([]int{-2, -1, 2, 1}, 1))
}

func maxSubArrayLen(nums []int, k int) int {
	sum := 0
	max := 0
	sums := map[int]int{}
	for i := range nums {
		sum += nums[i]
		if sum == k {
			max = i + 1
		}
		if j, ok := sums[sum-k]; ok && (i-j) > max {
			max = i - j
		}
		if _, ok := sums[sum]; !ok {
			sums[sum] = i
		}
	}
	return max
}
