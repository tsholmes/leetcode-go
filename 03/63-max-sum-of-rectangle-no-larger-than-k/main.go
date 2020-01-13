package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSumSubmatrix([][]int{
		{1, 0, 1},
		{0, -2, 3},
	}, 2))
}

func maxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	N, M := len(matrix), len(matrix[0])
	sums := make([][]int, N)
	for i := range matrix {
		sums[i] = make([]int, M)
		for j := range matrix[i] {
			switch {
			case i > 0 && j > 0:
				sums[i][j] = matrix[i][j] + sums[i-1][j] + sums[i][j-1] - sums[i-1][j-1]
			case i > 0:
				sums[i][j] = matrix[i][j] + sums[i-1][j]
			case j > 0:
				sums[i][j] = matrix[i][j] + sums[i][j-1]
			default:
				sums[i][j] = matrix[i][j]
			}
		}
	}

	max := math.MinInt64
	for i := range matrix {
		for j := range matrix[i] {
			for i2 := 0; i2 <= i; i2++ {
				for j2 := 0; j2 <= j; j2++ {
					var sum int
					switch {
					case i2 > 0 && j2 > 0:
						sum = sums[i][j] - sums[i2-1][j] - sums[i][j2-1] + sums[i2-1][j2-1]
					case i2 > 0:
						sum = sums[i][j] - sums[i2-1][j]
					case j2 > 0:
						sum = sums[i][j] - sums[i][j2-1]
					default:
						sum = sums[i][j]
					}
					if sum <= k && sum > max {
						max = sum
					}
				}
			}
		}
	}

	return max
}
