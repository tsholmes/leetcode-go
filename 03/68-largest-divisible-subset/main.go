package main

import "sort"

func main() {

}

func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	sizes := make([]int, len(nums))
	backs := make([]int, len(nums))
	var max, maxi int
	for i := range nums {
		sizes[i] = 1
		backs[i] = 1
		for j := range nums[:i] {
			if nums[i]%nums[j] == 0 && sizes[j] >= sizes[i] {
				sizes[i] = sizes[j] + 1
				backs[i] = j
			}
		}
		if sizes[i] > max {
			max = sizes[i]
			maxi = i
		}
	}
	res := make([]int, max)
	for max > 0 {
		res[max-1] = nums[maxi]
		max--
		maxi = backs[maxi]
	}
	return res
}
