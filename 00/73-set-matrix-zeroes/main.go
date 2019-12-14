package main

import "fmt"

func main() {
	do := func(matrix [][]int) {
		setZeroes(matrix)
		for _, line := range matrix {
			fmt.Println(line)
		}
		fmt.Println()
	}
	do([][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	})
	do([][]int{
		{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	})
	do([][]int{})
	do([][]int{{1}})
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	n := len(matrix)
	m := len(matrix[0])
	row1 := false
	col1 := false
	for i := 0; i < m; i++ {
		if matrix[0][i] == 0 {
			row1 = true
		}
	}
	for i := 0; i < n; i++ {
		if matrix[i][0] == 0 {
			col1 = true
		}
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < n; i++ {
		if matrix[i][0] != 0 {
			continue
		}
		for j := 1; j < m; j++ {
			matrix[i][j] = 0
		}
	}
	for j := 1; j < m; j++ {
		if matrix[0][j] != 0 {
			continue
		}
		for i := 1; i < n; i++ {
			matrix[i][j] = 0
		}
	}
	if row1 {
		for i := 0; i < m; i++ {
			matrix[0][i] = 0
		}
	}
	if col1 {
		for i := 0; i < n; i++ {
			matrix[i][0] = 0
		}
	}
}
