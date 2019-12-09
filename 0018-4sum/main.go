package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	fmt.Println(fourSum([]int{0, 0, 0, 0}, 0))
	fmt.Println(fourSum([]int{-1, 2, 2, -5, 0, -1, 4}, 3))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	var res [][]int

	ends := make(map[int]map[[2]int]int)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := len(nums) - 3; j > i; j-- {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			sum := nums[i] + nums[j]
			m, ok := ends[sum]
			if !ok {
				m = make(map[[2]int]int)
				ends[sum] = m
			}

			p := [2]int{nums[i], nums[j]}
			if j2, ok := m[p]; !ok || j < j2 {
				m[p] = j
			}
		}
	}

	for i := 2; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i < j-1 && nums[i] == nums[i+1] {
				continue
			}
			if j < len(nums)-1 && nums[j] == nums[j+1] {
				continue
			}

			s := target - nums[i] - nums[j]

			for p, k := range ends[s] {
				if i <= k {
					continue
				}
				res = append(res, []int{p[0], p[1], nums[i], nums[j]})
			}
		}
	}

	return res
}
