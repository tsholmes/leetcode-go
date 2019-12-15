package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
	fmt.Println(maxProduct([]int{-2, -3, 0, -4, -1}))
	fmt.Println(maxProduct([]int{-2}))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	maxNeg := math.MinInt64
	prod := 1
	for _, n := range nums {
		if n > max {
			max = n
		}
		prod *= n
		if prod == 0 {
			maxNeg = math.MinInt64
			prod = 1
			continue
		}
		p := max
		if prod < 0 {
			if maxNeg > math.MinInt64 {
				p = prod / maxNeg
			}
			if prod > maxNeg {
				maxNeg = prod
			}
		} else if prod > 0 {
			p = prod
		}
		if p > max {
			max = p
		}
	}
	return max
}
