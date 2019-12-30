package main

import (
	"container/heap"
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minMeetingRooms([][]int{
		{0, 30}, {5, 10}, {15, 20},
	}))
	fmt.Println(minMeetingRooms([][]int{
		{7, 10}, {2, 4},
	}))
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	var ends intHeap
	max := 0
	count := 0

	for _, iv := range intervals {
		for len(ends) > 0 && ends[0] <= iv[0] {
			heap.Pop(&ends)
			count--
		}
		heap.Push(&ends, iv[1])
		count++
		if count > max {
			max = count
		}
	}

	return max
}

type intHeap []int

var _ heap.Interface = &intHeap{}

func (h *intHeap) Len() int           { return len(*h) }
func (h *intHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *intHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
