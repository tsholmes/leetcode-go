package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(increasingTriplet([]int{5, 4, 3, 2, 1}))
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	min1 := nums[0]
	min2 := math.MaxInt32
	for _, n := range nums[1:] {
		if n > min2 {
			return true
		}
		if n > min1 && n < min2 {
			min2 = n
		}
		if n < min1 {
			min1 = n
		}
	}
	return false
}
