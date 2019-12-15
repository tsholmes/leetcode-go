package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{3, 3, 5, 0, 0, 3, 1, 4}))
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	held1, notHeld1 := math.MinInt64, 0
	held2, notHeld2 := math.MinInt64, 0
	for _, p := range prices {
		h1, n1 := held1, notHeld1
		h2, n2 := held2, notHeld2

		if -p > h1 {
			h1 = -p
		}
		if held1+p > n1 {
			n1 = held1 + p
		}
		if notHeld1-p > h2 {
			h2 = notHeld1 - p
		}
		if held2+p > n2 {
			n2 = held2 + p
		}

		held1, notHeld1 = h1, n1
		held2, notHeld2 = h2, n2
	}
	if notHeld1 > notHeld2 {
		notHeld2 = notHeld1
	}
	return notHeld2
}
