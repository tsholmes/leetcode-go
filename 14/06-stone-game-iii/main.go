package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(stoneGameIII([]int{1, 2, 3, 7}))
	fmt.Println(stoneGameIII([]int{1, 2, 3, -9}))
	fmt.Println(stoneGameIII([]int{1, 2, 3, 6}))
	fmt.Println(stoneGameIII([]int{1, 2, 3, -1, -2, -3, 7}))
	fmt.Println(stoneGameIII([]int{-1, -2, -3}))
}

func stoneGameIII(stoneValue []int) string {
	type state struct {
		i    int
		turn bool
	}
	memo := map[state]int{}
	var search func(i int, turn bool) int
	search = func(i int, turn bool) int {
		if i == len(stoneValue) {
			return 0
		}

		key := state{i, turn}
		if v, ok := memo[key]; ok {
			return v
		}

		var res int

		if turn {
			// Bob

			res = math.MaxInt32

			nlead := 0
			for j := i; j <= i+2 && j < len(stoneValue); j++ {
				nlead -= stoneValue[j]
				sres := search(j+1, false) + nlead
				if sres < res {
					res = sres
				}
			}
		} else {
			// Alice

			res = math.MinInt32

			nlead := 0
			for j := i; j <= i+2 && j < len(stoneValue); j++ {
				nlead += stoneValue[j]
				sres := nlead + search(j+1, true)
				if sres > res {
					res = sres
				}
			}
		}
		memo[key] = res

		return res
	}

	res := search(0, false)
	if res > 0 {
		return "Alice"
	} else if res == 0 {
		return "Tie"
	}
	return "Bob"
}

func signum(n int) int {
	if n > 0 {
		return 1
	} else if n < 0 {
		return -1
	}
	return 0
}
