package main

import "fmt"

func main() {
	fmt.Println(maxJumps([]int{6, 4, 14, 6, 8, 13, 9, 7, 10, 6, 12}, 2))
}

func maxJumps(arr []int, d int) int {
	edges := make([][]int, len(arr))
	for i := 0; i < len(arr); i++ {
		for x := 1; i-x >= 0 && x <= d; x++ {
			if arr[i-x] >= arr[i] {
				break
			}
			edges[i] = append(edges[i], i-x)
		}
		for x := 1; i+x < len(arr) && x <= d; x++ {
			if arr[i+x] >= arr[i] {
				break
			}
			edges[i] = append(edges[i], i+x)
		}
	}

	memo := map[int]int{}
	var count func(p int) int
	count = func(p int) int {
		if v, ok := memo[p]; ok {
			return v
		}
		max := 0
		for _, e := range edges[p] {
			c := count(e)
			if c > max {
				max = c
			}
		}
		memo[p] = max + 1
		return max + 1
	}

	max := 0
	for i := range arr {
		c := count(i)
		if c > max {
			max = c
		}
	}
	return max
}
