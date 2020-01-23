package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	fmt.Println(splitArray(
		[]int{10, 5, 13, 4, 8, 4, 5, 11, 14, 9, 16, 10, 20, 8},
		8,
	))
	nums := make([]int, 1000)
	for i := range nums {
		nums[i] = i
	}
	fmt.Println(splitArray(nums, 50))
}

func splitArray(nums []int, m int) int {
	memo := map[[3]int]int{}
	var search func(i, j, k int) int
	search = func(i, j, k int) int {
		if k > j-i || k < 0 {
			return math.MaxInt64
		}
		if i == j {
			return 0
		}
		key := [3]int{i, j, k}
		if v, ok := memo[key]; ok {
			return v
		}
		min := math.MaxInt64
		sum := 0
		for i2 := i; i2 < j; i2++ {
			sum += nums[i2]
			if sum > min {
				break
			}
			max := sum
			if v := search(i2+1, j, k-1); v > max {
				max = v
			}
			if max < min {
				min = max
			}
		}
		memo[key] = min
		return min
	}
	return search(0, len(nums), m)
}
