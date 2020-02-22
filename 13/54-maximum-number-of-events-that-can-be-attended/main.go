package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(isPossible([]int{9, 3, 5}))
	fmt.Println(isPossible([]int{1, 1, 1, 2}))
	fmt.Println(isPossible([]int{8, 5}))
	fmt.Println(isPossible([]int{9, 9, 9}))
}

func isPossible(target []int) bool {
	h := &intHeap{}
	sum := 0
	for _, v := range target {
		heap.Push(h, v)
		sum += v
	}
	for h.Len() > 0 {
		v := heap.Pop(h).(int)
		if v == 1 {
			continue
		}
		if v < 1 {
			return false
		}
		rem := sum - v
		if rem < 0 {
			return false
		}
		miss := v - rem
		if miss < 1 {
			return false
		}
		heap.Push(h, miss)
		sum += miss - v
	}
	return true
}

type intHeap []int

var _ heap.Interface = &intHeap{}

func (h *intHeap) Len() int           { return len(*h) }
func (h *intHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *intHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
