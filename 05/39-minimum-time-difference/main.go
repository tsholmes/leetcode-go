package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	fmt.Println(findMinDifference([]string{"12:12", "12:13"}))
}

func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	times := make([][2]int, len(timePoints))
	for i, s := range timePoints {
		h, _ := strconv.Atoi(s[:2])
		m, _ := strconv.Atoi(s[3:])
		times[i] = [2]int{h, m}
	}

	min := 24 * 60
	for i := range times {
		ti := times[i]
		tj := times[(i+1)%len(times)]
		if tj[0] < ti[0] || (tj[0] == ti[0] && tj[1] < ti[1]) {
			tj[0] += 24
		}
		mins := (tj[0]*60 + tj[1]) - (ti[0]*60 + ti[1])
		if mins < min {
			min = mins
		}
	}

	return min
}
