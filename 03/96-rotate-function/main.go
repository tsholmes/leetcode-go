package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxRotateFunction([]int{4, 3, 2, 6}))
}

func maxRotateFunction(A []int) int {
	if len(A) == 0 {
		return 0
	}
	N := len(A)
	var sum int
	cur := 0
	for i, n := range A {
		sum += n
		cur += i * n
	}
	max := math.MinInt64
	for i := range A {
		if cur > max {
			max = cur
		}
		cur += sum
		cur -= A[N-i-1] * N
	}
	return max
}
