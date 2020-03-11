package main

import (
	"container/heap"
	"fmt"
)

func main() {
	fmt.Println(findShortestWay([][]int{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 1, 0},
		{0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 1},
	}, []int{0, 4}, []int{3, 5},
	))
}

func findShortestWay(maze [][]int, ball []int, hole []int) string {
	N, M := len(maze), len(maze[0])
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	dirNames := "drul"
	hp := [2]int{hole[0], hole[1]}
	sp := [2]int{ball[0], ball[1]}
	roll := func(p [2]int, dir int) ([2]int, int) {
		d := dirs[dir]
		dist := 0
		for p != hp {
			p2 := [2]int{p[0] + d[0], p[1] + d[1]}
			if p2[0] < 0 || p2[0] >= N || p2[1] < 0 || p2[1] >= M || maze[p2[0]][p2[1]] == 1 {
				break
			}
			p = p2
			dist++
		}
		return p, dist
	}

	seen := map[[2]int]bool{}
	queue := &stateHeap{{sp, 0, ""}}
	res := "impossible"
	for queue.Len() > 0 {
		s := heap.Pop(queue).(dstate)

		if seen[s.p] {
			continue
		}

		seen[s.p] = true

		if s.p == hp {
			res = s.back
			break
		}

		for dir := 0; dir < 4; dir++ {
			p2, d := roll(s.p, dir)
			if d != 0 {
				heap.Push(queue, dstate{
					p:    p2,
					dist: s.dist + d,
					back: s.back + dirNames[dir:dir+1],
				})
			}
		}
	}

	return res
}

type dstate struct {
	p    [2]int
	dist int
	back string
}
type stateHeap []dstate

func (h *stateHeap) Len() int { return len(*h) }
func (h *stateHeap) Less(i, j int) bool {
	di, dj := (*h)[i].dist, (*h)[j].dist
	bi, bj := (*h)[i].back, (*h)[j].back
	if di < dj {
		return true
	} else if di > dj {
		return false
	}
	return bi < bj
}
func (h *stateHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.(dstate)) }
func (h *stateHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
