package main

import (
	"container/heap"
	"math"
)

func main() {

}

type MedianFinder struct {
	lt maxHeap
	gt minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{}
}

func (m *MedianFinder) AddNum(num int) {
	if float64(num) >= m.FindMedian() {
		heap.Push(&m.gt, num)
	} else {
		heap.Push(&m.lt, num)
	}
	for len(m.gt) > len(m.lt) {
		heap.Push(&m.lt, heap.Pop(&m.gt).(int))
	}
	for len(m.lt) > len(m.gt)+1 {
		heap.Push(&m.gt, heap.Pop(&m.lt).(int))
	}
}

func (m *MedianFinder) FindMedian() float64 {
	if len(m.lt) == 0 && len(m.gt) == 0 {
		return math.MaxFloat64
	}
	if len(m.lt) > len(m.gt) {
		return float64(m.lt[0])
	}
	return float64(m.lt[0]+m.gt[0]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

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
