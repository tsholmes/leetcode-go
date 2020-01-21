package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(reconstructQueue([][]int{
		{7, 0},
		{4, 4},
		{7, 1},
		{5, 0},
		{6, 1},
		{5, 2},
	}))
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i int, j int) bool {
		pi, pj := people[i], people[j]
		if pi[0] < pj[0] {
			return true
		} else if pi[0] > pj[0] {
			return false
		}
		return pi[1] < pj[1]
	})

	res := make([][]int, len(people))
	used := make([]bool, len(people))
	for _, p := range people {
		count := 0
		i := 0
		for used[i] || count < p[1] {
			if !used[i] || res[i][0] >= p[0] {
				count++
			}
			i++
		}
		res[i] = p
		used[i] = true
	}
	return res
}
