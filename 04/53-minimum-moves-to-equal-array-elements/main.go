package main

import "sort"

func main() {

}

func minMoves(nums []int) int {
	sort.Ints(nums)
	count := 0
	for _, n := range nums {
		count += n - nums[0]
	}
	return count
}
