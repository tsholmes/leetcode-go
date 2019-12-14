package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(grayCode(0))
	fmt.Println(grayCode(2))
	fmt.Println(grayCode(4))
}

func grayCode(n int) []int {
	res := make([]int, 1<<uint(n))
	for i := 1; i < len(res); i++ {
		if i == 1 {
			res[i] = 1
			continue
		}
		bit := 1 << uint(63-bits.LeadingZeros(uint(i)))
		res[i] = bit | res[2*bit-i-1]
	}
	return res
}
