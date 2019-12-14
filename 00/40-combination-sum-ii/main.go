package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	fmt.Println(combinationSum2([]int{2, 5, 2, 1, 2}, 5))
}

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	cpairs := make([][2]int, 0, len(candidates))
	last := candidates[0]
	lastCount := 1
	for i := 1; i < len(candidates); i++ {
		if candidates[i] == last {
			lastCount++
		} else {
			cpairs = append(cpairs, [2]int{last, lastCount})
			last = candidates[i]
			lastCount = 1
		}
	}
	cpairs = append(cpairs, [2]int{last, lastCount})

	var res [][]int

	var work []int
	var search func(int, int)
	search = func(t int, idx int) {
		if idx == len(cpairs) {
			if t == target {
				res = append(res, append([]int{}, work...))
			}
			return
		}

		search(t, idx+1)

		l := len(work)
		n := cpairs[idx]
		for i := 1; i*n[0]+t <= target && i <= n[1]; i++ {
			work = append(work, n[0])
			search(t+i*n[0], idx+1)
		}
		work = work[:l]
	}

	search(0, 0)

	return res
}
