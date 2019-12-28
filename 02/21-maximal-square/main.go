package main

func main() {

}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	hei, wid := len(matrix), len(matrix[0])

	max := 0

	dp := make([][]int, hei)
	for i := 0; i < hei; i++ {
		dp[i] = make([]int, wid)
		for j := 0; j < wid; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				dp[i][j] = 1
				if max == 0 {
					max = 1
				}
				continue
			}
			min := dp[i-1][j]
			if dp[i][j-1] < min {
				min = dp[i][j-1]
			}
			if dp[i-1][j-1] < min {
				min = dp[i-1][j-1]
			}
			if min >= max {
				max = min + 1
			}
			dp[i][j] = min + 1
		}
	}
	return max * max
}
