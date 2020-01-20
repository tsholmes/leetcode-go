package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(findNthDigit(i))
	}
}

func findNthDigit(n int) int {
	count := 0
	bidx := 1
	base := 1
	for {
		rcount := (base * 10) - base
		if n <= count+rcount*bidx {
			dn := n - count - 1
			ni := dn / bidx
			di := dn % bidx

			v := base + ni
			for k := 0; k < bidx-di-1; k++ {
				v /= 10
			}

			return v % 10
		}
		count += rcount * bidx
		bidx++
		base *= 10
	}
}
