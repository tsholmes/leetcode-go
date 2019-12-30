package main

import "fmt"

func main() {
	fmt.Println(nums)
}

var nums []int

func init() {
	nums = append(nums, 1)
	a, b, c := 0, 0, 0
	for n := 1; n < 1690; n++ {
		na, nb, nc := nums[a]*2, nums[b]*3, nums[c]*5
		min := na
		if nb < min {
			min = nb
		}
		if nc < min {
			min = nc
		}
		nums = append(nums, min)
		if na == min {
			a++
		}
		if nb == min {
			b++
		}
		if nc == min {
			c++
		}
	}
}

func nthUglyNumber(n int) int {
	return nums[n-1]
}
