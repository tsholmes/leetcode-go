package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	held, notHeld := math.MinInt64, 0
	for _, p := range prices {
		h, n := held, notHeld
		if held+p > n {
			n = held + p
		}
		if notHeld-p > h {
			h = notHeld - p
		}
		held, notHeld = h, n
	}
	return notHeld
}
