package main

import "container/heap"

func main() {

}

func shortestDistance(maze [][]int, start []int, destination []int) int {
	N, M := len(maze), len(maze[0])
	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	roll := func(p [2]int, dir int) ([2]int, int) {
		d := dirs[dir]
		dist := 0
		for {
			p2 := [2]int{p[0] + d[0], p[1] + d[1]}
			if p2[0] < 0 || p2[0] >= N || p2[1] < 0 || p2[1] >= M || maze[p2[0]][p2[1]] == 1 {
				break
			}
			p = p2
			dist++
		}
		return p, dist
	}

	sp := [2]int{start[0], start[1]}
	dp := [2]int{destination[0], destination[1]}
	seen := map[[2]int]bool{}
	queue := &stateHeap{{sp, 0}}
	for queue.Len() > 0 {
		s := heap.Pop(queue).(dstate)

		if seen[s.p] {
			continue
		}

		seen[s.p] = true

		if s.p == dp {
			return s.dist
		}

		for dir := 0; dir < 4; dir++ {
			p2, d := roll(s.p, dir)
			if d != 0 {
				heap.Push(queue, dstate{
					p:    p2,
					dist: s.dist + d,
				})
			}
		}
	}

	return -1
}

type dstate struct {
	p    [2]int
	dist int
}
type stateHeap []dstate

func (h *stateHeap) Len() int           { return len(*h) }
func (h *stateHeap) Less(i, j int) bool { return (*h)[i].dist < (*h)[j].dist }
func (h *stateHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.(dstate)) }
func (h *stateHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
