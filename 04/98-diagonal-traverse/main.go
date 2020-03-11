package main

import "fmt"

func main() {
	fmt.Println(findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}

func findDiagonalOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	N, M := len(matrix), len(matrix[0])

	var res []int

	passes := N + M - 1

	fwd := func(pass int) {
		i, j := pass, 0
		if i >= N {
			i, j = N-1, i-N+1
		}

		for i >= 0 && j < M {
			res = append(res, matrix[i][j])
			i, j = i-1, j+1
		}
	}

	back := func(pass int) {
		i, j := 0, pass
		if j >= M {
			i, j = j-M+1, M-1
		}

		for i < N && j >= 0 {
			res = append(res, matrix[i][j])
			i, j = i+1, j-1
		}
	}

	for pass := 0; pass < passes; pass++ {
		if pass%2 == 0 {
			fwd(pass)
		} else {
			back(pass)
		}
	}

	return res
}
