package main

import "sort"

func main() {

}

func minSubsequence(nums []int) []int {
	tsum := 0
	for i := range nums {
		tsum += nums[i]
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	var res []int
	sum := 0
	for sum <= tsum-sum {
		sum += nums[0]
		res = append(res, nums[0])
		nums = nums[1:]
	}

	return res
}
