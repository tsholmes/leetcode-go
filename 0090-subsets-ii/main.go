package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	work := make([]int, len(nums))
	var search func(int, int)
	search = func(i int, count int) {
		if i == len(nums) {
			res = append(res, append([]int{}, work[:count]...))
			return
		}

		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		work[count] = nums[i]
		search(i+1, count+1)
		search(j, count)
	}
	search(0, 0)
	return res
}
