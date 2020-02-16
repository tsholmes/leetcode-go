package main

import (
	"math"
	"math/bits"
)

func main() {

}

func maxStudents(seats [][]byte) int {
	if len(seats) == 0 || len(seats[0]) == 0 {
		return 0
	}
	N, M := len(seats), len(seats[0])

	memo := map[[2]int]int{}
	var search func(i, mask int) int
	search = func(i, mask int) int {
		if i == N {
			return 0
		}
		k := [2]int{i, mask}
		if v, ok := memo[k]; ok {
			return v
		}

		max := math.MinInt32

		for m2 := 0; m2 < 1<<uint(M); m2++ {
			if !rowValid(m2, seats[i]) {
				continue
			}
			if !valid(mask, m2, M) {
				continue
			}
			v := bits.OnesCount(uint(m2)) + search(i+1, m2)
			if v > max {
				max = v
			}
		}

		memo[k] = max
		return max
	}
	return search(0, 0)
}

func rowValid(m int, row []byte) bool {
	for i, b := range row {
		if m&(1<<uint(i)) != 0 && b == '#' {
			return false
		}
	}
	return true
}

func valid(m1, m2, M int) bool {
	for i := uint(0); i < uint(M); i++ {
		if m2&(1<<i) == 0 {
			continue
		}
		if i > 0 && m1&(1<<(i-1)) != 0 {
			return false
		}
		if i > 0 && m2&(1<<(i-1)) != 0 {
			return false
		}
		if i < uint(M)-1 && m1&(1<<(i+1)) != 0 {
			return false
		}
		if i < uint(M)-1 && m2&(1<<(i+1)) != 0 {
			return false
		}
	}
	return true
}
