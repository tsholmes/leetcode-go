package main

import "sort"

func main() {

}

func findUnsortedSubarray(nums []int) int {
	snums := make([]int, len(nums))
	for i, n := range nums {
		snums[i] = n
	}
	sort.Ints(snums)
	seen := map[[2]int]bool{}
	for i, n := range snums {
		seen[[2]int{i, n}] = true
	}
	min, max := len(nums), -1
	for i, n := range nums {
		if !seen[[2]int{i, n}] {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
	}
	if max < min {
		return 0
	}
	return max - min + 1
}
