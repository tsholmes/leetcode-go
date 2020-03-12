package main

import "container/heap"

func main() {

}

func updateMatrix(matrix [][]int) [][]int {
	N, M := len(matrix), len(matrix[0])

	res := make([][]int, N)
	for i := range res {
		res[i] = make([]int, M)
	}

	queue := &stateHeap{}
	seen := map[[2]int]bool{}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	for i, row := range matrix {
		for j, n := range row {
			if n == 0 {
				heap.Push(queue, [3]int{i, j, 0})
			}
		}
	}

	for queue.Len() > 0 {
		s := heap.Pop(queue).([3]int)
		k := [2]int{s[0], s[1]}
		if seen[k] {
			continue
		}
		seen[k] = true
		res[s[0]][s[1]] = s[2]
		for _, d := range dirs {
			i2, j2 := s[0]+d[0], s[1]+d[1]
			if i2 < 0 || i2 >= N || j2 < 0 || j2 >= M {
				continue
			}
			heap.Push(queue, [3]int{i2, j2, s[2] + 1})
		}
	}

	return res
}

type stateHeap [][3]int

func (h *stateHeap) Len() int           { return len(*h) }
func (h *stateHeap) Less(i, j int) bool { return (*h)[i][2] < (*h)[j][2] }
func (h *stateHeap) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *stateHeap) Push(x interface{}) { *h = append(*h, x.([3]int)) }
func (h *stateHeap) Pop() interface{}   { v := (*h)[h.Len()-1]; *h = (*h)[:h.Len()-1]; return v }
