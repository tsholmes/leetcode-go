package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	}))
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	N, M := len(heightMap), len(heightMap[0])

	heights := make([][]int, N)
	for i := 0; i < N; i++ {
		heights[i] = make([]int, M)
		for j := 0; j < M; j++ {
			if i == 0 || i == N-1 || j == 0 || j == M-1 {
				heights[i][j] = heightMap[i][j]
			} else {
				heights[i][j] = math.MaxInt32
			}
		}
	}

	changed := true
	for changed {
		changed = false

		// down
		for j := 1; j < M-1; j++ {
			curHeight := heights[0][j]
			for i := 1; i < N-1; i++ {
				if curHeight < heightMap[i][j] {
					curHeight = heightMap[i][j]
				}
				if curHeight > heights[i][j] {
					curHeight = heights[i][j]
				}
				if curHeight < heights[i][j] {
					heights[i][j] = curHeight
					changed = true
				}
			}
		}
		// up
		for j := 1; j < M-1; j++ {
			curHeight := heights[N-1][j]
			for i := N - 1; i > 0; i-- {
				if curHeight < heightMap[i][j] {
					curHeight = heightMap[i][j]
				}
				if curHeight > heights[i][j] {
					curHeight = heights[i][j]
				}
				if curHeight < heights[i][j] {
					heights[i][j] = curHeight
					changed = true
				}
			}
		}
		// right
		for i := 1; i < N-1; i++ {
			curHeight := heights[i][0]
			for j := 1; j < M-1; j++ {
				if curHeight < heightMap[i][j] {
					curHeight = heightMap[i][j]
				}
				if curHeight > heights[i][j] {
					curHeight = heights[i][j]
				}
				if curHeight < heights[i][j] {
					heights[i][j] = curHeight
					changed = true
				}
			}
		}
		// left
		for i := 1; i < N-1; i++ {
			curHeight := heights[i][M-1]
			for j := M - 2; j > 0; j-- {
				if curHeight < heightMap[i][j] {
					curHeight = heightMap[i][j]
				}
				if curHeight > heights[i][j] {
					curHeight = heights[i][j]
				}
				if curHeight < heights[i][j] {
					heights[i][j] = curHeight
					changed = true
				}
			}
		}
	}

	res := 0
	for i := 1; i < N-1; i++ {
		for j := 1; j < M-1; j++ {
			res += heights[i][j] - heightMap[i][j]
		}
	}

	return res
}
