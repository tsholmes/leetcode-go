package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 1, 0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{{1}}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0},
		{0},
	}))
	fmt.Println(uniquePathsWithObstacles([][]int{
		{0, 0},
	}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	n := len(obstacleGrid)
	m := len(obstacleGrid[0])
	counts := make([]int, m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 && j == 0 {
				counts[j] = 1
			}
			if obstacleGrid[i][j] == 1 {
				counts[j] = 0
				continue
			}
			if j > 0 {
				counts[j] += counts[j-1]
			}
		}
	}
	return counts[m-1]
}
