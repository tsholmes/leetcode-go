package main

import "fmt"

func main() {
	fmt.Println(maximalRectangle([][]byte{
		[]byte(`10100`),
		[]byte(`10111`),
		[]byte(`11111`),
		[]byte(`10010`),
	}))
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n, m := len(matrix), len(matrix[0])
	dp := make([][]int, n)
	max := 0
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			dp[i][j] = 1
			if i > 0 {
				dp[i][j] += dp[i-1][j]
			}
			min := dp[i][j]
			for k := j; k >= 0; k-- {
				if dp[i][k] < min {
					min = dp[i][k]
				}
				a := min * (j - k + 1)
				if a > max {
					max = a
				}
			}
		}
	}
	for _, d := range dp {
		fmt.Println(d)
	}

	return max
}
