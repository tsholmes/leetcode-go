package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getSkyline([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	}))
}

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return nil
	}
	var xs []int
	for _, b := range buildings {
		xs = append(xs, b[0], b[1])
	}
	sort.Ints(xs)
	count := 1
	for _, x := range xs[1:] {
		if x != xs[count-1] {
			xs[count] = x
			count++
		}
	}
	xs = xs[:count]

	max := make([]int, len(xs))
	i := 0
	for _, b := range buildings {
		for xs[i] < b[0] {
			i++
		}
		for j := i; j < len(xs) && xs[j] < b[1]; j++ {
			if b[2] > max[j] {
				max[j] = b[2]
			}
		}
	}

	var res [][]int
	lastHei := -1
	for i := range xs {
		x := xs[i]
		hei := max[i]
		if hei != lastHei {
			res = append(res, []int{x, hei})
		}
		lastHei = hei
	}

	return res
}
