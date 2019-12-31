package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1e5))
}

func numSquares(n int) int {
	memo := map[int]int{}
	var search func(int) int
	search = func(n int) int {
		if n == 0 {
			return 0
		}
		sqrt := int(math.Sqrt(float64(n)))
		if sqrt*sqrt == n {
			return 1
		}
		if v, ok := memo[n]; ok {
			return v
		}
		min := math.MaxInt64
		for f := sqrt; f*f >= sqrt; f-- {
			v := 1 + search(n-f*f)
			if v < min {
				min = v
			}
		}
		memo[n] = min
		return min
	}

	return search(n)
}
