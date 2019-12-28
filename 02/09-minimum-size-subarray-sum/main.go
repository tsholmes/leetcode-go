package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen(s int, nums []int) int {
	min := math.MaxInt64
	i, j, sum := 0, 0, 0
	for j < len(nums) || sum >= s {
		if sum < s {
			if j == len(nums) {
				break
			}
			sum += nums[j]
			j++
		} else {
			sum -= nums[i]
			i++
		}
		if sum >= s && j-i < min {
			min = j - i
		}
	}

	if min == math.MaxInt64 {
		min = 0
	}
	return min
}
