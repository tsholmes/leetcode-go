package main

import "fmt"

func main() {
	do := func(matrix [][]int) {
		rotate(matrix)
		for _, line := range matrix {
			fmt.Println(line)
		}
		fmt.Println()
	}
	do([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	do([][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	})
	do([][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	})
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			matrix[i][j], matrix[n-j-1][i],
				matrix[j][n-i-1], matrix[n-i-1][n-j-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1],
				matrix[i][j], matrix[j][n-i-1]
		}
	}
}
