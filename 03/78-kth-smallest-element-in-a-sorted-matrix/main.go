package main

import "container/heap"

func main() {

}

func kthSmallest(matrix [][]int, k int) int {
	h := intHeap{}
	for _, row := range matrix {
		for _, v := range row {
			heap.Push(&h, v)
			if len(h) > k {
				heap.Pop(&h)
			}
		}
	}
	return heap.Pop(&h).(int)
}

type intHeap []int

var _ heap.Interface = &intHeap{}

func (h *intHeap) Len() int           { return len(*h) }
func (h *intHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *intHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
