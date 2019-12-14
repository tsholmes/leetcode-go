package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(3, 2))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(100, 2))
}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	a, b := m+n-2, n-1
	if b > a-b {
		b = a - b
	}

	res := 1

	for k := 0; k < b; k++ {
		res *= a - k
		res /= k + 1
	}

	return res
}
