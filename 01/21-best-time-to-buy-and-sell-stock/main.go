package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	max := 0
	min := prices[0]
	for _, p := range prices {
		pro := p - min
		if pro > max {
			max = pro
		}
		if p < min {
			min = p
		}
	}
	return max
}
