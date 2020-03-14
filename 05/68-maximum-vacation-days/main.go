package main

import "fmt"

func main() {
	fmt.Println(maxVacationDays([][]int{
		{0, 1, 1},
		{1, 0, 1},
		{1, 1, 0},
	}, [][]int{
		{1, 3, 1},
		{6, 0, 3},
		{3, 3, 3},
	}))
}

func maxVacationDays(flights [][]int, days [][]int) int {
	N, K := len(days), len(days[0])

	memo := map[[2]int]int{}
	var search func(i, k int) int
	search = func(i, k int) int {
		if k == K {
			return 0
		}
		key := [2]int{i, k}
		if v, ok := memo[key]; ok {
			return v
		}

		max := days[i][k] + search(i, k+1)
		for j := 0; j < N; j++ {
			if flights[i][j] == 0 {
				continue
			}
			if mv := search(j, k+1) + days[j][k]; mv > max {
				max = mv
			}
		}
		memo[key] = max
		return max
	}
	return search(0, 0)
}
