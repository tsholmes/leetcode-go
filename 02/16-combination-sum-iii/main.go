package main

import "fmt"

func main() {
	fmt.Println(combinationSum3(3, 7))
	fmt.Println(combinationSum3(3, 9))
}

func combinationSum3(k int, n int) [][]int {
	var res [][]int

	work := make([]int, k)
	var search func(int, int, int)
	search = func(i int, j int, s int) {
		if i == k {
			if s == n {
				res = append(res, append([]int{}, work...))
			}
			return
		}
		if s >= n {
			return
		}
		for j2 := j; j2 <= 9; j2++ {
			work[i] = j2
			search(i+1, j2+1, s+j2)
		}
	}
	search(0, 1, 0)

	return res
}
