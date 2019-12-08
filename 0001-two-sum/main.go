package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	diffs := make(map[int]int, len(nums))
	for i, n := range nums {
		diffs[target-n] = i
	}
	for i, n := range nums {
		j, ok := diffs[n]
		if ok && j != i {
			return []int{i, j}
		}
	}
	panic("not found")
}
