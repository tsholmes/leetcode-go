package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(nextGreaterElement(12))
	fmt.Println(nextGreaterElement(21))
	fmt.Println(nextGreaterElement(102))
	fmt.Println(nextGreaterElement(0))
	fmt.Println(nextGreaterElement(230241))
	fmt.Println(nextGreaterElement(12222333))
}

func nextGreaterElement(n int) int {
	if n == 0 {
		return -1
	}
	var ds []int
	for n > 0 {
		ds = append(ds, n%10)
		n /= 10
	}

	tail := 1
	for j := 1; j < len(ds); j++ {
		if ds[j] >= ds[j-1] {
			tail++
		} else {
			break
		}
	}
	if tail == len(ds) {
		return -1
	}

	for i := 0; i < tail/2; i++ {
		j := tail - i - 1
		ds[i], ds[j] = ds[j], ds[i]
	}

	nextv := math.MaxInt32
	nexti := 0
	for i := 0; i < tail; i++ {
		d := ds[i] - ds[tail]
		if d > 0 && d <= nextv {
			nextv = d
			nexti = i
		}
	}

	ds[nexti], ds[tail] = ds[tail], ds[nexti]

	for i := len(ds) - 1; i >= 0; i-- {
		n = (n * 10) + ds[i]
	}

	if n > math.MaxInt32 {
		return -1
	}

	return n
}
