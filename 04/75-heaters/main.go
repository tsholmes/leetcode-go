package main

import (
	"math"
	"sort"
)

func main() {

}

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)

	max := 0

	for _, h := range houses {
		idx := sort.Search(len(heaters), func(i int) bool { return heaters[i] >= h })
		min := math.MaxInt64
		if idx < len(heaters) {
			min = heaters[idx] - h
		}
		if idx > 0 {
			m2 := h - heaters[idx-1]
			if m2 < min {
				min = m2
			}
		}
		if min > max {
			max = min
		}
	}

	return max
}
