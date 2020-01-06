package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minTotalDistance([][]int{
		{1, 0, 0, 0, 1},
		{0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0},
	}))
}

func minTotalDistance(grid [][]int) int {
	var pts [][2]int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				pts = append(pts, [2]int{i, j})
			}
		}
	}

	bi := sort.Search(len(grid), func(j int) bool {
		lt, gt := 0, 0
		for _, p := range pts {
			if p[0] <= j {
				lt++
			} else {
				gt++
			}
		}
		return lt >= gt
	})
	bj := sort.Search(len(grid[0]), func(j int) bool {
		lt, gt := 0, 0
		for _, p := range pts {
			if p[1] <= j {
				lt++
			} else {
				gt++
			}
		}
		return lt >= gt
	})

	abs := func(n int) int {
		if n < 0 {
			return -n
		}
		return n
	}

	res := 0
	for _, p := range pts {
		res += abs(p[0]-bi) + abs(p[1]-bj)
	}
	return res
}
