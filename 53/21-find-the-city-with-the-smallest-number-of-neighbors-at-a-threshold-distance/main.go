package main

import (
	"container/heap"
	"math"
)

func main() {

}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	es := map[int][][2]int{}
	for _, e := range edges {
		es[e[0]] = append(es[e[0]], [2]int{e[1], e[2]})
		es[e[1]] = append(es[e[1]], [2]int{e[0], e[2]})
	}

	count := func(i int) int {
		queue := stateHeap{[2]int{i, 0}}
		res := 0
		seen := map[int]bool{}
		for len(queue) > 0 {
			s := heap.Pop(&queue).(dstate)
			k := s[0]
			if seen[k] {
				continue
			}
			seen[k] = true

			if s[1] > distanceThreshold {
				break
			}

			res++

			for _, e := range es[k] {
				heap.Push(&queue, dstate{e[0], s[1] + e[1]})
			}
		}
		return res
	}

	min := math.MaxInt64
	mini := 0
	for i := 0; i < n; i++ {
		c := count(i)
		if c <= min {
			min = c
			mini = i
		}
	}
	return mini
}

type dstate [2]int
type stateHeap []dstate

func (h *stateHeap) Len() int           { return len(*h) }
func (h *stateHeap) Less(i, j int) bool { return (*h)[i][1] < (*h)[j][1] }
func (h *stateHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.(dstate)) }
func (h *stateHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
