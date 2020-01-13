package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(getMoneyAmount(10))
}

func getMoneyAmount(n int) int {
	memo := map[[2]int]int{}
	var search func(int, int) int
	search = func(i, j int) int {
		if i == j {
			return 0
		}
		k := [2]int{i, j}
		if v, ok := memo[k]; ok {
			return v
		}

		min := math.MaxInt64
		for k := i; k <= j; k++ {
			el, er := 0, 0
			if k > i {
				el = search(i, k-1)
			}
			if k < j {
				er = search(k+1, j)
			}
			v := k
			if el > er {
				v += el
			} else {
				v += er
			}
			if v < min {
				min = v
			}
		}
		memo[k] = min

		return min
	}
	return search(1, n)
}
