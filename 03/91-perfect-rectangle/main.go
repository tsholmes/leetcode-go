package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(isRectangleCover([][]int{
		{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4},
	}))
}

func isRectangleCover(rectangles [][]int) bool {
	yBounds := map[int]bool{}
	minX := math.MaxInt32
	maxX := math.MinInt32
	for _, r := range rectangles {
		yBounds[r[0]] = true
		yBounds[r[2]] = true
		if r[1] < minX {
			minX = r[1]
		}
		if r[3] > maxX {
			maxX = r[3]
		}
	}
	var ys []int
	for y := range yBounds {
		ys = append(ys, y)
	}
	sort.Ints(ys)

	y1 := ys[0]
	var rs [][]int
	for _, y2 := range ys[1:] {
		rs = rs[:0]
		for _, r := range rectangles {
			if r[0] < y2 && r[2] > y1 {
				rs = append(rs, r)
			}
		}
		sort.Slice(rs, func(i, j int) bool { return rs[i][1] < rs[j][1] })
		x1 := minX
		for _, r := range rs {
			if r[1] != x1 {
				return false
			}
			x1 = r[3]
		}
		if x1 != maxX {
			return false
		}
		y1 = y2
	}
	return true
}
