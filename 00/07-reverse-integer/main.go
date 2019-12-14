package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
}

func reverse(x int) int {
	if x < 0 {
		return -reverse(-x)
	}
	y := 0
	for x > 0 {
		y *= 10
		y += x % 10
		x /= 10
	}
	if y > math.MaxInt32 {
		return 0
	}
	return y
}
