package main

import "math"

func main() {

}

func minDistance(height int, width int, tree []int, squirrel []int, nuts [][]int) int {
	halfs := make([]int, len(nuts))
	total := 0
	for i, np := range nuts {
		halfs[i] = abs(np[0]-tree[0]) + abs(np[1]-tree[1])
		total += 2 * halfs[i]
	}

	min := math.MaxInt32
	for i, np := range nuts {
		d := total - halfs[i] + abs(np[0]-squirrel[0]) + abs(np[1]-squirrel[1])
		if d < min {
			min = d
		}
	}

	return min
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
