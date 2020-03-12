package main

import (
	"math/rand"
	"sort"
)

func main() {

}

type Solution struct {
	cw    []int
	total int
}

func Constructor(w []int) Solution {
	cw := make([]int, len(w))
	total := 0
	for i, ww := range w {
		total += ww
		cw[i] = total
	}
	return Solution{cw: cw, total: total}
}

func (s *Solution) PickIndex() int {
	w := rand.Intn(s.total)
	return sort.Search(len(s.cw), func(i int) bool {
		return w < s.cw[i]
	})
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */
