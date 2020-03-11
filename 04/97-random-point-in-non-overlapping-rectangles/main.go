package main

import (
	"math/rand"
	"sort"
)

func main() {

}

type Solution struct {
	rects  [][]int
	cumas  []int
	totala int
}

func Constructor(rects [][]int) Solution {
	totala := 0
	cumas := make([]int, len(rects))
	for i, r := range rects {
		a := (r[2] - r[0] + 1) * (r[3] - r[1] + 1)
		totala += a
		cumas[i] = totala
	}
	return Solution{
		rects:  rects,
		cumas:  cumas,
		totala: totala,
	}
}

func (s *Solution) Pick() []int {
	ai := rand.Intn(s.totala)
	i := sort.Search(len(s.cumas), func(i int) bool { return s.cumas[i] > ai })
	r := s.rects[i]
	x := r[0] + rand.Intn(r[2]-r[0]+1)
	y := r[1] + rand.Intn(r[3]-r[1]+1)
	return []int{x, y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
