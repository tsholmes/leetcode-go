package main

import "fmt"

func main() {
	fmt.Println(combine(4, 2))
	fmt.Println(combine(5, 1))
	fmt.Println(combine(5, 5))
	fmt.Println(combine(4, 0))
	fmt.Println(combine(4, 5))
}

func combine(n int, k int) [][]int {
	if k < 1 || k > n {
		return nil
	}
	var res [][]int
	work := make([]int, k)
	var search func(int)
	search = func(i int) {
		if i == k {
			res = append(res, append([]int{}, work...))
			return
		}
		start := 1
		if i > 0 {
			start = work[i-1] + 1
		}
		for j := start; j <= n; j++ {
			work[i] = j
			search(i + 1)
		}
	}

	search(0)

	return res
}
