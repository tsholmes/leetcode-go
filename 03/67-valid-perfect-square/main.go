package main

import "fmt"

func main() {
	for i := 0; i <= 1e6; i++ {
		if !isPerfectSquare(i * i) {
			fmt.Println(i)
		}
	}
}

func isPerfectSquare(num int) bool {
	sqrt := float64(num) / 2
	for i := 0; i < 32 && sqrt > 0; i++ {
		sqrt2 := float64(num) / sqrt
		sqrt = (sqrt + sqrt2) / 2
	}
	isqrt := int(sqrt + 0.5)
	return isqrt*isqrt == num
}
