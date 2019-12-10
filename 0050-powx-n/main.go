package main

import "fmt"

func main() {
	fmt.Println(myPow(2.00000, 10))
	fmt.Println(myPow(2.10000, 3))
	fmt.Println(myPow(2.00000, -2))
}

func myPow(x float64, n int) float64 {
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	v := 1.0

	for n > 0 {
		if n&1 != 0 {
			v *= x
		}
		x *= x
		n /= 2
	}
	if neg {
		v = 1.0 / v
	}
	return v
}
