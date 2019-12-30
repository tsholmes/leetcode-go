package main

import "math"

func main() {

}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	for i := 1; i < len(costs); i++ {
		for j := 0; j < 3; j++ {
			c1 := costs[i][j] + costs[i-1][(j+1)%3]
			c2 := costs[i][j] + costs[i-1][(j+2)%3]
			if c1 < c2 {
				costs[i][j] = c1
			} else {
				costs[i][j] = c2
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
