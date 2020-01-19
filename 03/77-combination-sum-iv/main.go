package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
	fmt.Println(combinationSum4(
		[]int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 110, 120, 130, 140, 150, 160, 170, 180, 190, 200, 210, 220, 230, 240, 250, 260, 270, 280, 290, 300, 310, 320, 330, 340, 350, 360, 370, 380, 390, 400, 410, 420, 430, 440, 450, 460, 470, 480, 490, 500, 510, 520, 530, 540, 550, 560, 570, 580, 590, 600, 610, 620, 630, 640, 650, 660, 670, 680, 690, 700, 710, 720, 730, 740, 750, 760, 770, 780, 790, 800, 810, 820, 830, 840, 850, 860, 870, 880, 890, 900, 910, 920, 930, 940, 950, 960, 970, 980, 990, 111},
		999,
	))
}

func combinationSum4(nums []int, target int) int {
	counts := map[int]int{}
	h := intHeap{}
	for _, n := range nums {
		heap.Push(&h, n)
		counts[n] = 1
	}

	for len(h) > 0 {
		v := heap.Pop(&h).(int)
		if counts[v] > math.MaxInt32 {
			continue
		}
		if v == target {
			return counts[v]
		} else if v > target {
			break
		}

		for _, n := range nums {
			if v+n > target {
				continue
			}
			add := counts[v+n] == 0
			counts[v+n] += counts[v]
			if add {
				heap.Push(&h, v+n)
			}
		}
	}

	return 0
}

type intHeap []int

var _ heap.Interface = &intHeap{}

func (h *intHeap) Len() int           { return len(*h) }
func (h *intHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }
func (h *intHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
