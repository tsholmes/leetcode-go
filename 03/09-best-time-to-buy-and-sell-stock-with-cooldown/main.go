package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	held := make([]int, len(prices))
	notHeld := make([]int, len(prices))
	held[0] = -prices[0]
	for i := 1; i < len(prices); i++ {
		held[i] = held[i-1]
		nheld := notHeld[i-1] - prices[i]
		if i > 1 {
			nheld = notHeld[i-2] - prices[i]
		}
		if nheld > held[i] {
			held[i] = nheld
		}
		notHeld[i] = notHeld[i-1]
		if held[i-1]+prices[i] > notHeld[i] {
			notHeld[i] = held[i-1] + prices[i]
		}
	}
	return notHeld[len(prices)-1]
}
