package main

import "math"

func main() {

}
func luckyNumbers(matrix [][]int) []int {
	N, M := len(matrix), len(matrix[0])
	minr, maxc := make([]int, N), make([]int, M)
	for i := range minr {
		minr[i] = math.MaxInt32
	}
	for i := range maxc {
		maxc[i] = math.MinInt32
	}
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] < minr[i] {
				minr[i] = matrix[i][j]
			}
			if matrix[i][j] > maxc[j] {
				maxc[j] = matrix[i][j]
			}
		}
	}
	var res []int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == minr[i] && matrix[i][j] == maxc[j] {
				res = append(res, matrix[i][j])
			}
		}
	}
	return res
}
