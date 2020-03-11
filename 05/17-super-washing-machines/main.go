package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMinMoves([]int{1, 0, 5}))
	fmt.Println(findMinMoves([]int{0, 3, 0}))
	fmt.Println(findMinMoves([]int{0, 2, 0}))
	fmt.Println(findMinMoves([]int{1, 7, 1, 1, 7, 1}))
	fmt.Println(findMinMoves([]int{1, 9, 1, 1, 5, 1}))
	fmt.Println(findMinMoves([]int{100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0}))
	fmt.Println(findMinMoves([]int{7, 6, 5, 4, 3, 2, 1}))
	fmt.Println(findMinMoves([]int{0, 0, 14, 0, 10, 0, 0, 0}))
	fmt.Println(findMinMoves([]int{9, 1, 8, 8, 9}))
}

func findMinMoves(machines []int) int {
	sum := 0
	for _, n := range machines {
		sum += n
	}
	if sum%len(machines) != 0 {
		return -1
	}
	exp := sum / len(machines)

	maxTime := 0
	diff := 0
	for _, n := range machines {
		ndiff := n - exp
		diff += ndiff
		maxTime = max3(maxTime, abs(diff), ndiff)
	}

	return maxTime
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func max3(a, b, c int) int {
	if a > b && a > c {
		return a
	} else if b > c {
		return b
	}
	return c
}
