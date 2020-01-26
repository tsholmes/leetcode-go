package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minDifficulty([]int{11, 111, 22, 222, 33, 333, 44, 444}, 6))
	fmt.Println(minDifficulty([]int{9, 9, 9}, 4))
}

func minDifficulty(jobDifficulty []int, d int) int {
	type key struct {
		idx  int
		drem int
	}
	memo := map[key]int{}
	var search func(int, int) int
	search = func(i int, drem int) int {
		if len(jobDifficulty)-i < drem {
			return math.MaxInt32
		}
		if drem < 0 {
			return math.MaxInt32
		}
		if i == len(jobDifficulty) {
			return 0
		}

		k := key{idx: i, drem: drem}
		if v, ok := memo[k]; ok {
			return v
		}

		min := math.MaxInt32
		max := 0
		for j := i; j < len(jobDifficulty); j++ {
			if jobDifficulty[j] > max {
				max = jobDifficulty[j]
			}
			s := max + search(j+1, drem-1)
			if s < min {
				min = s
			}
		}

		memo[k] = min

		return min
	}

	r := search(0, d)
	if r >= math.MaxInt32 {
		return -1
	}
	return r
}
