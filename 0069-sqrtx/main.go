package main

import (
	"fmt"
	"sort"
)

func main() {
	for i := 0; i <= 16; i++ {
		fmt.Println(mySqrt(i))
	}
}

func mySqrt(x int) int {
	v := sort.Search(x, func(i int) bool {
		return i*i >= x
	})
	if v*v > x {
		v--
	}
	return v
}
