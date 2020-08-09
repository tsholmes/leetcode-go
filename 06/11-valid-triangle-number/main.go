package main

import "sort"

func main() {

}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			res += sort.Search(len(nums)-j-1, func(k int) bool {
				return nums[j+k+1] >= nums[i]+nums[j]
			})
		}
	}
	return res
}
