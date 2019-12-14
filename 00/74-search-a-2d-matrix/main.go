package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 3))
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 50},
	}, 13))
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	n, m := len(matrix), len(matrix[0])
	idx := sort.Search(n*m, func(i int) bool {
		r, c := i/m, i%m
		return matrix[r][c] >= target
	})
	if idx == n*m {
		return false
	}
	r, c := idx/m, idx%m
	return matrix[r][c] == target
}
