package main

import "fmt"

func main() {
	do := func(knows [][]int) {
		f := solution(func(a, b int) bool { return knows[a][b] == 1 })
		fmt.Println(f(len(knows)))
	}
	do([][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 1},
	})
	do([][]int{
		{1, 0, 1},
		{0, 1, 1},
		{0, 0, 1},
	})
	do([][]int{
		{1, 1, 1},
		{0, 1, 1},
		{0, 0, 1},
	})
}

/**
 * The knows API is already defined for you.
 *     knows := func(a int, b int) bool
 */
func solution(knows func(a int, b int) bool) func(n int) int {
	return func(n int) int {
		cand := make([]int, n)
		for i := range cand {
			cand[i] = i
		}
		for len(cand) > 1 {
			count := 0
			for i := 0; i+1 < len(cand); i += 2 {
				j := i + 1
				ci, cj := cand[i], cand[j]
				f, b := knows(ci, cj), knows(cj, ci)
				if b && !f {
					cand[count] = ci
					count++
				}
				if f && !b {
					cand[count] = cj
					count++
				}
			}
			if len(cand)&1 != 0 {
				cand[count] = cand[len(cand)-1]
				count++
			}
			cand = cand[:count]
		}
		if len(cand) == 0 {
			return -1
		}
		for i := 0; i < n; i++ {
			if i == cand[0] {
				continue
			}
			if knows(cand[0], i) || !knows(i, cand[0]) {
				return -1
			}
		}
		return cand[0]
	}
}
