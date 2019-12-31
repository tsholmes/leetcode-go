package main

import "fmt"

func main() {
	fmt.Println(numWays(3, 2))
}

func numWays(n int, k int) int {
	if n == 0 || k < 1 {
		return 0
	}
	if k == 1 {
		if n > 2 {
			return 0
		}
		return 1
	}
	r1, r2 := k, 0
	for i := 1; i < n; i++ {
		r1, r2 = (r1+r2)*(k-1), r1
	}
	return r1 + r2
}
