package main

import "fmt"

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
	fmt.Println(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i < c {
				continue
			}
			if i-c != 0 && res[i-c] == 0 {
				continue
			}
			if res[i] == 0 || res[i-c]+1 < res[i] {
				res[i] = res[i-c] + 1
			}
		}
	}
	if res[amount] == 0 {
		return -1
	}
	return res[amount]
}
