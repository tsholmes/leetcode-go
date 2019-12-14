package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSubArray([]int{-2, -1}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	cumsum := 0
	mincumsum := 0

	for _, n := range nums {
		cumsum += n
		d := cumsum - mincumsum
		if d > max {
			max = d
		}
		if cumsum < mincumsum {
			mincumsum = cumsum
		}
	}

	return max
}
