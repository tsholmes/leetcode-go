package main

import (
	"math/rand"
)

func main() {
}

type Solution struct {
	rows, cols int
	removed    int
	mapping    map[int][2]int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		rows:    n_rows,
		cols:    n_cols,
		removed: 0,
		mapping: map[int][2]int{},
	}
}

func (s *Solution) get(i int) [2]int {
	if v, ok := s.mapping[i]; ok {
		return v
	}
	row, col := i/s.cols, i%s.cols
	return [2]int{row, col}
}

func (s *Solution) Flip() []int {
	remi := s.rows*s.cols - s.removed - 1
	i := rand.Intn(remi + 1)

	res, swap := s.get(i), s.get(remi)
	s.mapping[i] = swap
	s.mapping[remi] = res
	s.removed++

	return res[:]
}

func (s *Solution) Reset() {
	s.removed = 0
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(n_rows, n_cols);
 * param_1 := obj.Flip();
 * obj.Reset();
 */
