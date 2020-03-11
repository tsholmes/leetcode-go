package main

import "container/heap"

func main() {

}

func numTimesAllBlue(light []int) int {
	var h maxHeap
	res := 0
	for _, n := range light {
		heap.Push(&h, n)
		if h.Len() == h[0] {
			res++
		}
	}
	return res
}

type maxHeap []int

func (h *maxHeap) Len() int           { return len(*h) }
func (h *maxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *maxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
