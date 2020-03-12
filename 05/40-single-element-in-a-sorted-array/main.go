package main

import "sort"

func main() {

}

func singleNonDuplicate(nums []int) int {
	i := sort.Search(len(nums)-1, func(i int) bool {
		if i&1 == 0 {
			return nums[i] != nums[i+1]
		} else {
			return nums[i] != nums[i-1]
		}
	})
	return nums[i]
}
