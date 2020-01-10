package main

import (
	"fmt"
)

func main() {
	fmt.Println(multiply(
		[][]int{
			{1, -5},
		},
		[][]int{
			{12},
			{-1},
		},
	))
}

func multiply(A [][]int, B [][]int) [][]int {
	if len(A) == 0 || len(A[0]) == 0 || len(B) == 0 {
		return nil
	}
	L, N := len(A), len(B[0])

	aIdxs := make([][]int, L)
	for i := range A {
		for j := range A[i] {
			if A[i][j] != 0 {
				aIdxs[i] = append(aIdxs[i], j)
			}
		}
	}

	bIdxs := make([][]int, N)
	for i := range B {
		for j := range B[i] {
			if B[i][j] != 0 {
				bIdxs[j] = append(bIdxs[j], i)
			}
		}
	}

	res := make([][]int, L)

	for i := 0; i < L; i++ {
		res[i] = make([]int, N)
		for j := 0; j < N; j++ {
			pair(aIdxs[i], bIdxs[j], func(k int) {
				res[i][j] += A[i][k] * B[k][j]
			})
		}
	}

	return res
}

func pair(a []int, b []int, f func(int)) {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		default:
			f(a[i])
			i++
			j++
		}
	}
}
