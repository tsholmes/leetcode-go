package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(constrainedSubsetSum([]int{10, -2, -10, -5, 20}, 2))
	fmt.Println(constrainedSubsetSum([]int{10, 2, -10, 5, 20}, 2))
	fmt.Println(constrainedSubsetSum([]int{-1, -2, -3}, 1))
}

func constrainedSubsetSum(nums []int, k int) int {
	var nums2 []int

	{
		var curneg []int
		var cursum int

		buildneg := func() {
			if len(curneg) < k {
				return
			}

			h := &pairHeap{}
			for i := 0; i < k; i++ {
				heap.Push(h, pair{curneg[i], i})
			}
			for i := k; i < len(curneg); i++ {
				for i-(*h)[0][1] > k {
					heap.Pop(h)
				}
				min := (*h)[0]
				heap.Push(h, pair{min[0] + curneg[i], i})
			}
			for len(curneg)-(*h)[0][1] > k {
				heap.Pop(h)
			}
			nums2 = append(nums2, (*h)[0][0])
		}

		for _, n := range nums {
			if n >= 0 {
				if len(curneg) > 0 {
					buildneg()
					curneg = curneg[:0]
				}
				cursum += n
				continue
			}
			if cursum > 0 {
				nums2 = append(nums2, cursum)
				cursum = 0
			}
			curneg = append(curneg, n)
		}
		if cursum > 0 {
			nums2 = append(nums2, cursum)
		}
		if len(curneg) > 0 {
			buildneg()
		}
	}

	if len(nums2) == 1 && nums2[0] < 0 {
		max := math.MinInt32
		for _, n := range nums {
			if n > max {
				max = n
			}
		}
		return max
	}

	cum := 0
	mincum := 0

	max := 0

	for _, n := range nums2 {
		cum += n
		if cum-mincum > max {
			max = cum - mincum
		}
		if cum < mincum {
			mincum = cum
		}
	}

	return max
}

type pair [2]int
type pairHeap []pair

func (h *pairHeap) Len() int           { return len(*h) }
func (h *pairHeap) Less(i, j int) bool { return (*h)[i][0] > (*h)[j][0] }
func (h *pairHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *pairHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *pairHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
