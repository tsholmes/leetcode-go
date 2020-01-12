package main

import "fmt"

func main() {
	fmt.Println(countBits(16))
}

func countBits(num int) []int {
	hpow, pow := 0, 1
	res := make([]int, num+1)
	for n := 1; n <= num; n++ {
		if n == pow {
			res[n] = 1
			hpow, pow = pow, pow<<1
		} else {
			res[n] = res[n-hpow] + 1
		}
	}
	return res
}
