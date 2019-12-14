package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minPathSum([][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	for i := range grid {
		for j := range grid[i] {
			if i == 0 && j == 0 {
				continue
			}
			min := math.MaxInt64
			if i > 0 {
				min = grid[i][j] + grid[i-1][j]
			}
			if j > 0 && grid[i][j-1]+grid[i][j] < min {
				min = grid[i][j-1] + grid[i][j]
			}
			grid[i][j] = min
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
