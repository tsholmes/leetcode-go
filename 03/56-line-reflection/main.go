package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isReflected([][]int{{1, 1}, {-1, 1}}))
	fmt.Println(isReflected([][]int{{1, 1}, {-1, -1}}))
	fmt.Println(isReflected([][]int{{0, 0}, {1, 0}, {3, 0}}))
}

func isReflected(points [][]int) bool {
	if len(points) == 0 {
		return true
	}
	rows := map[int][]int{}
	for _, p := range points {
		rows[p[1]] = append(rows[p[1]], p[0]*2)
	}
	mid := func(xs []int) (int, bool) {
		sort.Ints(xs)
		count := 1
		for i := 1; i < len(xs); i++ {
			if xs[i] != xs[count-1] {
				xs[count] = xs[i]
				count++
			}
		}
		xs = xs[:count]
		i, j := 0, len(xs)-1
		m := (xs[i] + xs[j]) / 2
		i++
		j--
		for i <= j {
			m2 := (xs[i] + xs[j]) / 2
			if m2 != m {
				return 0, false
			}
			i++
			j--
		}
		return m, true
	}
	mids := map[int]bool{}
	for _, xs := range rows {
		m, ok := mid(xs)
		if !ok {
			return false
		}
		mids[m] = true
	}
	return len(mids) == 1
}
