package main

import "fmt"

func main() {
	fmt.Println(numTrees(3))
	fmt.Println(numTrees(5))
	fmt.Println(numTrees(100))
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	memo := map[[2]int]int{}
	var search func(int, int) int
	search = func(i, j int) int {
		if i >= j {
			return 1
		}
		key := [2]int{i, j}
		res, ok := memo[key]
		if ok {
			return res
		}
		for c := i; c <= j; c++ {
			ls, rs := search(i, c-1), search(c+1, j)
			res += ls * rs
		}
		memo[key] = res
		return res
	}
	return search(1, n)
}
