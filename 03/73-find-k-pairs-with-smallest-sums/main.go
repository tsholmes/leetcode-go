package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(kSmallestPairs([]int{1, 7, 11}, []int{2, 4, 6}, 3))
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var h pairHeap
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			heap.Push(&h, []int{n1, n2})
			for h.Len() > k {
				heap.Pop(&h)
			}
		}
	}
	return [][]int(h)
}

type pairHeap [][]int

func (h *pairHeap) Len() int           { return len(*h) }
func (h *pairHeap) Less(i, j int) bool { return (*h)[i][0]+(*h)[i][1] > (*h)[j][0]+(*h)[j][1] }
func (h *pairHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *pairHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
