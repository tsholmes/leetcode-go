package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{
		{1, 3}, {6, 9},
	}, []int{2, 5}))
	fmt.Println(insert([][]int{
		{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
	}, []int{4, 8}))
}

func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	for len(intervals) > 0 && intervals[0][1] < newInterval[0] {
		res = append(res, intervals[0])
		intervals = intervals[1:]
	}
	start, end := newInterval[0], newInterval[1]
	for len(intervals) > 0 && intervals[0][0] <= end {
		if intervals[0][0] < start {
			start = intervals[0][0]
		}
		if intervals[0][1] > end {
			end = intervals[0][1]
		}
		intervals = intervals[1:]
	}
	res = append(res, []int{start, end})
	res = append(res, intervals...)
	return res
}
