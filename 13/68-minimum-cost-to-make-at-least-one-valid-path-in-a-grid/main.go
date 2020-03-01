package main

import "fmt"

func main() {
	fmt.Println(minCost([][]int{
		{1, 1, 1, 1},
		{2, 2, 2, 2},
		{1, 1, 1, 1},
		{2, 2, 2, 2},
	}))
}

func minCost(grid [][]int) int {
	N, M := len(grid), len(grid[0])
	vis := make([][]bool, N)
	for i := range vis {
		vis[i] = make([]bool, M)
	}

	min := N + M
	dirs := [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	mins := map[[2]int]int{}
	var search func(i, j, cost int)
	search = func(i, j, cost int) {
		if cost >= min {
			return
		}
		if vis[i][j] {
			return
		}
		if i == N-1 && j == M-1 {
			min = cost
			return
		}
		k := [2]int{i, j}
		if v, ok := mins[k]; ok && v <= cost {
			return
		}
		vis[i][j] = true
		mins[k] = cost
		for di, d := range dirs {
			i2, j2, c2 := i+d[0], j+d[1], cost
			if di != grid[i][j]-1 {
				c2++
			}
			if i2 < 0 || i2 >= N || j2 < 0 || j2 >= M {
				continue
			}
			search(i2, j2, c2)
		}
		vis[i][j] = false
	}
	search(0, 0, 0)
	return min
}
