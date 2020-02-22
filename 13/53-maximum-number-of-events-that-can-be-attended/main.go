package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func maxEvents(events [][]int) int {
	var starts [100001][]int
	for _, e := range events {
		starts[e[0]] = append(starts[e[0]], e[1])
	}
	for i := 1; i <= 100000; i++ {
		sort.Ints(starts[i])
	}

	count := 0
	ends := &intHeap{}
	for i := 1; i <= 100000; i++ {
		for _, e := range starts[i] {
			heap.Push(ends, e)
		}
		for ends.Len() > 0 {
			e := heap.Pop(ends).(int)
			if e < i {
				continue
			}
			count++
			break
		}
	}
	return count
}

type intHeap []int

var _ heap.Interface = &intHeap{}

func (h *intHeap) Len() int           { return len(*h) }
func (h *intHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *intHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
