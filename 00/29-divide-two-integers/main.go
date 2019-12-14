package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(divide(10, 3))
	fmt.Println(divide(7, -3))
}

func divide(dividend int, divisor int) int {
	neg := false
	if dividend < 0 {
		dividend = -dividend
		neg = !neg
	}
	if divisor < 0 {
		divisor = -divisor
		neg = !neg
	}

	if dividend < divisor {
		return 0
	}

	res := 0
	mul := 1
	for divisor*mul*10 <= dividend {
		mul *= 10
	}

	for mul > 0 {
		for dividend >= divisor*mul {
			dividend -= divisor * mul
			res += mul
		}
		mul /= 10
	}

	if neg {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}
