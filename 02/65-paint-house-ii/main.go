package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minCostII([][]int{
		{1, 5, 3},
		{2, 9, 4},
	}))
	fmt.Println(minCostII([][]int{
		{3, 14, 12, 2, 20, 16, 12, 2},
		{9, 6, 9, 8, 2, 9, 20, 18},
		{20, 2, 20, 4, 5, 12, 11, 11},
		{16, 3, 7, 5, 15, 2, 2, 4},
		{17, 3, 11, 1, 9, 5, 7, 11},
	}))
}

func minCostII(costs [][]int) int {
	if len(costs) == 0 || (len(costs) > 1 && len(costs[0]) < 2) {
		return 0
	}
	for i := 1; i < len(costs); i++ {
		min0, min1 := costs[i-1][0], costs[i-1][1]
		if min0 > min1 {
			min0, min1 = min1, min0
		}
		for j := 2; j < len(costs[i-1]); j++ {
			if costs[i-1][j] <= min0 {
				min0, min1 = costs[i-1][j], min0
			} else if costs[i-1][j] <= min1 {
				min1 = costs[i-1][j]
			}
		}
		for j := range costs[i] {
			if costs[i-1][j] == min0 {
				costs[i][j] += min1
			} else {
				costs[i][j] += min0
			}
		}
	}
	min := math.MaxInt64
	for _, c := range costs[len(costs)-1] {
		if c < min {
			min = c
		}
	}
	return min
}
