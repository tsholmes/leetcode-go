package main

import "container/heap"

func main() {

}

func longestSubarray(nums []int, limit int) int {
	minh := &minHeap{}
	maxh := &maxHeap{}

	minrem := map[int]int{}
	maxrem := map[int]int{}

	left := 0

	maxlen := 0

	for i, n := range nums {
		heap.Push(minh, n)
		heap.Push(maxh, n)

		for {
			for minrem[(*minh)[0]] > 0 {
				v := heap.Pop(minh).(int)
				minrem[v]--
			}
			for maxrem[(*maxh)[0]] > 0 {
				v := heap.Pop(maxh).(int)
				maxrem[v]--
			}
			min := (*minh)[0]
			max := (*maxh)[0]

			if max-min <= limit {
				if i-left >= maxlen {
					maxlen = i - left + 1
				}
				break
			}

			minrem[nums[left]]++
			maxrem[nums[left]]++
			left++
		}
	}

	return maxlen
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
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
