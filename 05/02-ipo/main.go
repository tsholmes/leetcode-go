package main

import (
	"container/heap"
	"sort"
)

func main() {

}

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	pairs := make([][2]int, len(Profits))
	for i := range pairs {
		pairs[i] = [2]int{Profits[i], Capital[i]}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] < pairs[j][1] })

	cap := W
	used := 0
	h := &pairHeap{}
	for used < k {
		for len(pairs) > 0 && pairs[0][1] <= cap {
			heap.Push(h, pairs[0])
			pairs = pairs[1:]
		}
		if h.Len() == 0 {
			break
		}
		p := heap.Pop(h).([2]int)
		cap += p[0]
		used++
	}

	return cap
}

type pairHeap [][2]int

func (h *pairHeap) Len() int           { return len(*h) }
func (h *pairHeap) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *pairHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.([2]int)) }
func (h *pairHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
