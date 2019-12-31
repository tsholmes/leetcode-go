package main

import "sort"

func main() {

}

func wiggleSort(nums []int) {
	sort.Ints(nums)
	for i := 2; i < len(nums); i += 2 {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
}
