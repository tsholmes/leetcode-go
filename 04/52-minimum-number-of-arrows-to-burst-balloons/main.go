package main

import (
	"math"
	"sort"
)

func main() {

}

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})
	end := math.MinInt64
	count := 0
	for _, p := range points {
		if p[0] <= end {
			continue
		}
		end = p[1]
		count++
	}
	return count
}
