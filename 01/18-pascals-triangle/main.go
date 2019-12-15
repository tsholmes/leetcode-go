package main

import "fmt"

func main() {
	fmt.Println(generate(5))
}

func generate(numRows int) [][]int {
	if numRows < 1 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				res[i][j] = 1
			} else {
				res[i][j] = res[i-1][j-1] + res[i-1][j]
			}
		}
	}
	return res
}
