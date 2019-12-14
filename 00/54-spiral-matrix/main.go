package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}))
	fmt.Println(spiralOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}))
}

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}
	n, m := len(matrix), len(matrix[0])
	pos := [2]int{0, 0}
	dir := [2]int{0, 1}
	res := make([]int, 0, n*m)

	for len(res) < cap(res) {
		res = append(res, matrix[pos[0]][pos[1]])
		matrix[pos[0]][pos[1]] = math.MinInt64

		pos2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if pos2[0] < 0 || pos2[0] == n || pos2[1] < 0 || pos2[1] == m || matrix[pos2[0]][pos2[1]] == math.MinInt64 {
			dir = [2]int{dir[1], -dir[0]}
			pos2 = [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		}
		pos = pos2
	}

	return res
}
