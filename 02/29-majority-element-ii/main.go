package main

import "sort"

func main() {

}

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	var res []int
	i := 0
	for i < len(nums) {
		j := i
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		if j-i > len(nums)/3 {
			res = append(res, nums[i])
		}
		i = j
	}
	return res
}
