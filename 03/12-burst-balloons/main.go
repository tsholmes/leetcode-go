package main

import "fmt"

func main() {
	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
	fmt.Println(maxCoins([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}))

	ns := make([]int, 500)
	for i := range ns {
		ns[i] = i % 100
	}
	fmt.Println(maxCoins(ns))
}

func maxCoins(nums []int) int {
	memo := map[[4]int]int{}
	var search func(int, int, int, int) int
	search = func(i, j, lv, rv int) int {
		if i > j || j < i {
			return 0
		}
		if i == j {
			return nums[i] * lv * rv
		}
		key := [4]int{i, j, lv, rv}
		if v, ok := memo[key]; ok {
			return v
		}

		max := 0

		for k := i; k <= j; k++ {
			v := lv*nums[k]*rv + search(i, k-1, lv, nums[k]) + search(k+1, j, nums[k], rv)
			if v > max {
				max = v
			}
		}

		memo[key] = max
		return max
	}
	return search(0, len(nums)-1, 1, 1)
}
