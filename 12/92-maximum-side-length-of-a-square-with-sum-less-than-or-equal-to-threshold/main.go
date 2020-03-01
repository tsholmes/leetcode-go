package main

import "fmt"

func main() {
	fmt.Println(maxSideLength([][]int{
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
		{2, 2, 2, 2, 2},
	}, 1))
}

func maxSideLength(mat [][]int, threshold int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	m := len(mat)
	n := len(mat[0])

	max := 0

	sums := make([][][]int, m)
	for i := 0; i < m; i++ {
		sums[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			sums[i][j] = make([]int, n+m)
			for s := 1; s <= i+1 && s <= j+1; s++ {
				sum := mat[i][j]
				if s > 1 {
					sum += sums[i-1][j-1][s-2]
				}
				for k := 1; k < s; k++ {
					sum += mat[i-k][j]
					sum += mat[i][j-k]
				}
				sums[i][j][s-1] = sum
				if sum <= threshold && s > max {
					max = s
				}
			}
		}
	}
	return max
}
