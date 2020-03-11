package main

import "math"

func main() {

}

func findRotateSteps(ring string, key string) int {
	cpos := map[byte][]int{}
	for i := 0; i < len(ring); i++ {
		cpos[ring[i]] = append(cpos[ring[i]], i)
	}

	dist := func(i, j int) int {
		if i > j {
			i, j = j, i
		}
		d := i + len(ring) - j
		if j-i < d {
			d = j - i
		}
		return d
	}

	memo := map[[2]int]int{}
	var search func(i, j int) int
	search = func(i, j int) int {
		if j == len(key) {
			return 0
		}
		k := [2]int{i, j}
		if v, ok := memo[k]; ok {
			return v
		}
		min := math.MaxInt32
		for _, i2 := range cpos[key[j]] {
			d := 1 + dist(i, i2) + search(i2, j+1)
			if d < min {
				min = d
			}
		}
		memo[k] = min
		return min
	}
	return search(0, 0)
}
