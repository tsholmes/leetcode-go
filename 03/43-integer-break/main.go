package main

import (
	"fmt"
)

func main() {
	fmt.Println(integerBreak(6))
	fmt.Println(integerBreak(20))
}

func integerBreak(n int) int {
	if n < 3 {
		return 1
	}
	if n == 3 {
		return 2
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	a, b, c := 1, 2, 3
	for i := 4; i <= n; i++ {
		a, b, c = b, c, max(max(a*3, b*2), c*1)
	}
	return c
}
