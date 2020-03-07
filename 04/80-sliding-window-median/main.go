package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(medianSlidingWindow([]int{
		6, 5, 9, 5, 4, 9, 1, 7, 5, 5,
	}, 4))
}

func medianSlidingWindow(nums []int, k int) []float64 {
	var left maxHeap
	var right minHeap

	for i := 0; i < k-1; i++ {
		heap.Push(&left, nums[i])
	}

	rem := map[int]int{}

	lex, rex := 0, 0

	var res []float64
	for i := k - 1; i < len(nums); i++ {
		if len(left) > 0 && nums[i] <= left[0] {
			heap.Push(&left, nums[i])
		} else {
			heap.Push(&right, nums[i])
		}
		if i-k >= 0 {
			rem[nums[i-k]]++
			if nums[i-k] <= left[0] {
				lex++
			} else {
				rex++
			}
		}

		for len(left)-lex+1 > len(right)-rex {
			n := heap.Pop(&left).(int)
			if rem[n] > 0 {
				rem[n]--
				lex--
				continue
			}
			heap.Push(&right, n)
		}

		for len(left) > 0 && rem[left[0]] > 0 {
			n := heap.Pop(&left).(int)
			rem[n]--
			lex--
		}

		for len(right)-rex > len(left)-lex {
			n := heap.Pop(&right).(int)
			if rem[n] > 0 {
				rem[n]--
				rex--
				continue
			}
			heap.Push(&left, n)
		}

		for len(right) > 0 && rem[right[0]] > 0 {
			n := heap.Pop(&right).(int)
			rem[n]--
			rex--
		}

		if k&1 == 1 {
			res = append(res, float64(left[0]))
		} else {
			res = append(res, float64(left[0]+right[0])/2.0)
		}
	}

	return res
}

type minHeap []int

func (h *minHeap) Len() int           { return len(*h) }
func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *minHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }

type maxHeap []int

func (h *maxHeap) Len() int           { return len(*h) }
func (h *maxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *maxHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
