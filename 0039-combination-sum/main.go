package main

import "fmt"

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	var work []int
	var search func(int, int)
	search = func(t int, idx int) {
		if idx == len(candidates) {
			if t == target {
				res = append(res, append([]int{}, work...))
			}
			return
		}

		search(t, idx+1)

		l := len(work)
		n := candidates[idx]
		for i := 1; i*n+t <= target; i++ {
			work = append(work, n)
			search(t+i*n, idx+1)
		}
		work = work[:l]
	}

	search(0, 0)

	return res
}
