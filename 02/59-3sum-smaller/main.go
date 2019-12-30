package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSumSmaller([]int{-2, 0, 1, 3}, 2))
	fmt.Println(threeSumSmaller([]int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 4))
}

func threeSumSmaller(nums []int, target int) int {
	res := 0

	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			if nums[i]+nums[j]+nums[k] >= target {
				k--
				continue
			}
			for nums[i]+nums[j]+nums[k] < target && j < k {
				res += k - j
				j++
			}
		}
	}

	return res
}
