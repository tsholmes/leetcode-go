package main

import (
	"math"
	"strconv"
)

func main() {

}

func optimalDivision(nums []int) string {
	type node struct {
		v    int
		l, r *node
	}
	type key struct {
		i, j int
		max  bool
	}
	type res struct {
		n *node
		v float64
	}
	ens := make([]node, len(nums))
	for i, n := range nums {
		ens[i] = node{v: n}
	}
	memo := map[key]res{}
	var search func(i, j int, max bool) (*node, float64)
	search = func(i, j int, max bool) (*node, float64) {
		if i == j {
			return &ens[i], float64(nums[i])
		}
		k := key{i, j, max}
		if v, ok := memo[k]; ok {
			return v.n, v.v
		}

		bestv := math.MaxFloat64
		if max {
			bestv = -math.MaxFloat64
		}
		var bestn *node
		for split := i; split < j; split++ {
			ln, lv := search(i, split, max)
			rn, rv := search(split+1, j, !max)
			dv := lv / rv
			if (max && dv > bestv) || (!max && dv < bestv) {
				bestv = dv
				bestn = &node{l: ln, r: rn}
			}
		}
		memo[k] = res{n: bestn, v: bestv}
		return bestn, bestv
	}

	root, _ := search(0, len(nums)-1, true)
	var work []byte
	var build func(n *node)
	build = func(n *node) {
		if n.v != 0 {
			work = strconv.AppendInt(work, int64(n.v), 10)
			return
		}
		if n.r.v != 0 {
			build(n.l)
			work = append(work, '/')
			build(n.r)
			return
		}
		build(n.l)
		work = append(work, '/')
		if n.r.v == 0 {
			work = append(work, '(')
		}
		build(n.r)
		if n.r.v == 0 {
			work = append(work, ')')
		}
	}
	build(root)

	return string(work)
}
