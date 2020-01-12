package main

import (
	"fmt"
	"math"
)

func main() {
	m := 1
	for i := 0; i <= 19; i++ {
		fmt.Println(isPowerOfThree(m-1), isPowerOfThree(m), isPowerOfThree(m+1))
		m *= 3
	}
}

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	l := math.Log(float64(n)) / math.Log(3.0)
	return math.Abs(l-math.Floor(l)) < 1e-10
}
