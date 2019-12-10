package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	fmt.Println(merge([][]int{
		{1, 4}, {4, 5},
	}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var res [][]int
	start, end := intervals[0][0], intervals[0][1]
	for _, iv := range intervals[1:] {
		if iv[0] > end {
			res = append(res, []int{start, end})
			start, end = iv[0], iv[1]
		} else if iv[1] > end {
			end = iv[1]
		}
	}
	res = append(res, []int{start, end})

	return res
}
