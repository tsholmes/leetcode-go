package main

import "sort"

func main() {

}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	ds := make([]int, 0, 6)
	ps := [][]int{p1, p2, p3, p4}
	for i, pi := range ps {
		for _, pj := range ps[i+1:] {
			dx, dy := pi[0]-pj[0], pi[1]-pj[1]
			ds = append(ds, dx*dx+dy*dy)
		}
	}
	sort.Ints(ds)
	if ds[0] == 0 {
		return false
	}
	sl := ds[0]
	for i, d := range ds {
		if i < 4 && d != sl {
			return false
		} else if i >= 4 && d != sl*2 {
			return false
		}
	}
	return true
}
