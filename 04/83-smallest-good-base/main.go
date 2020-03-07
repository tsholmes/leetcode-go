package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(smallestGoodBase("13"))
	fmt.Println(smallestGoodBase("4681"))
	fmt.Println(smallestGoodBase("1000000000000000000"))
}

func smallestGoodBase(ns string) string {
	n64, _ := strconv.ParseInt(ns, 10, 64)

	n := int(n64)

	min := math.MaxInt64

	for ds := 2; ds < 64; ds++ {
		b := sort.Search(n-2, func(i int) bool {
			return baseN(i+2, ds) >= n
		}) + 2
		if baseN(b, ds) == n && b < min {
			min = b
		}
	}

	return strconv.Itoa(min)
}

func baseN(base int, n int) int {
	r := 1
	b := 1
	for i := 1; i < n; i++ {
		if math.MaxInt64/b < base {
			return math.MaxInt64
		}
		b *= base
		if math.MaxInt64-b < r {
			return math.MaxInt64
		}
		r += b
	}
	return r
}
