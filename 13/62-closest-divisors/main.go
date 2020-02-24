package main

import "fmt"

func main() {
	fmt.Println(closestDivisors(1e9))
}

func closestDivisors(num int) []int {
	f1, f2 := find(num+1), find(num+2)
	if f1[1]-f1[0] < f2[1]-f2[0] {
		return f1[:]
	}
	return f2[:]
}

func find(n int) [2]int {
	min := [2]int{1, n}
	mind := n - 1
	for i := 1; i*i <= n; i++ {
		if n%i != 0 {
			continue
		}
		j := n / i
		if j-i < mind {
			mind = j - 2
			min = [2]int{i, j}
		}
	}
	return min
}
