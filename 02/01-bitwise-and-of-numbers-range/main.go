package main

import "fmt"

func main() {
	fmt.Println(rangeBitwiseAnd(5, 7))
	fmt.Println(rangeBitwiseAnd(0, 1))
}

func rangeBitwiseAnd(m int, n int) int {
	res := 0
	for i := 31; i >= 0; i-- {
		b := 1 << uint(i)
		if (m & b) != (n & b) {
			break
		}
		res |= m & b
	}
	return res
}
