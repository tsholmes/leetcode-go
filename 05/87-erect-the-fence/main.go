package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(outerTrees([][]int{
		{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2},
	}))
	fmt.Println(outerTrees([][]int{
		{1, 2}, {2, 2}, {4, 2},
	}))
}

func outerTrees(points [][]int) [][]int {
	if len(points) < 4 {
		return points
	}
	{
		mini := 0
		for i, p := range points {
			if p[0] < points[mini][0] {
				mini = i
			}
		}
		points[0], points[mini] = points[mini], points[0]
	}
	{
		maxi := 0
		for i, p := range points {
			if p[0] > points[maxi][0] {
				maxi = i
			}
		}
		if points[maxi][0] == points[0][0] {
			return points
		}
		points[1], points[maxi] = points[maxi], points[1]
	}

	res := [][]int{points[0], points[1]}
	var build func(p0, p1 []int, ps [][]int)
	build = func(p0, p1 []int, ps [][]int) {
		// filter to left
		lcount := 0
		for _, p := range ps {
			if cross(p0, p1, p) >= 0 {
				ps[lcount] = p
				lcount++
			}
		}
		ps = ps[:lcount]

		if len(ps) == 0 {
			return
		}

		// find max
		{
			maxi := 0
			maxc := cross(p0, p1, ps[0])
			for i, p := range ps {
				c := cross(p0, p1, p)
				if c > maxc {
					maxi, maxc = i, c
				}
			}
			ps[0], ps[maxi] = ps[maxi], ps[0]
		}

		mid := ps[0]
		ps = ps[1:]
		res = append(res, mid)

		// split by dot
		l, r := 0, len(ps)-1
		mdot := dot(p0, p1, mid)
		for l <= r {
			d := dot(p0, p1, ps[l])
			if d < mdot {
				l++
			} else {
				ps[l], ps[r] = ps[r], ps[l]
				r--
			}
		}

		build(p0, mid, ps[:l])
		build(mid, p1, ps[l:])
	}

	ps1, ps2 := make([][]int, len(points)-2), make([][]int, len(points)-2)
	copy(ps1, points[2:])
	copy(ps2, points[2:])

	build(points[0], points[1], ps1)
	build(points[1], points[0], ps2)

	sort.Slice(res, func(i, j int) bool {
		pi, pj := res[i], res[j]
		if pi[0] < pj[0] {
			return true
		} else if pi[0] > pj[0] {
			return false
		} else {
			return pi[1] < pj[1]
		}
	})

	last := []int{-1, -1}
	count := 0
	for _, p := range res {
		if p[0] != last[0] || p[1] != last[1] {
			res[count] = p
			count++
		}
		last = p
	}
	res = res[:count]

	return res
}

// (p1 - p0) x (p2 - p0)
func cross(p0, p1, p2 []int) int {
	return (p1[0]-p0[0])*(p2[1]-p0[1]) - (p2[0]-p0[0])*(p1[1]-p0[1])
}

// (p1 - p0) . (p2 - p0)
func dot(p0, p1, p2 []int) int {
	return (p1[0]-p0[0])*(p2[0]-p0[0]) + (p1[1]-p0[1])*(p2[1]-p0[1])
}
