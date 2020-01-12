package main

import "fmt"

func main() {
	fmt.Println(isSelfCrossing([]int{2, 1, 1, 2}))
	fmt.Println(isSelfCrossing([]int{1, 2, 3, 4}))
	fmt.Println(isSelfCrossing([]int{1, 1, 1, 1}))
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 1, 1}))
	fmt.Println(isSelfCrossing([]int{3, 3, 4, 2, 2}))
	fmt.Println(isSelfCrossing([]int{1, 1, 2, 2, 1, 1}))
}

func isSelfCrossing(x []int) bool {
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var diri int
	var tail [][2]int
	var pos [2]int
	tail = append(tail, pos)

	between := func(a, b, c int) bool {
		if b > c {
			b, c = c, b
		}
		return a >= b && a <= c
	}

	cross := func(a0, a1, b0, b1 [2]int) bool {
		if a0[0] == a1[0] {
			// a-vert b-horiz
			return between(a0[0], b0[0], b1[0]) && between(b0[1], a0[1], a1[1])
		} else {
			// a-horiz b-vert
			return between(a0[1], b0[1], b1[1]) && between(b0[0], a0[0], a1[0])
		}
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	max := func(a, b int) int {
		return -min(-a, -b)
	}

	over := func(a0, a1, b0, b1 [2]int) bool {
		if a0[0] == a1[0] {
			// a-vert b-vert
			return a0[0] == b0[0] && min(a0[1], a1[1]) <= max(b0[1], b1[1]) && max(a0[1], a1[1]) >= min(b0[1], b1[1])
		} else {
			// a-horiz b-horiz
			return a0[1] == b0[1] && min(a0[0], a1[0]) <= max(b0[0], b1[0]) && max(a0[0], a1[0]) >= min(b0[0], b1[0])
		}
	}

	for _, dist := range x {
		dir := dirs[diri]
		diri = (diri + 1) & 3
		pos = [2]int{pos[0] + dist*dir[0], pos[1] + dist*dir[1]}
		for o := 3; o+1 <= len(tail); o++ {
			a0, a1, b0 := tail[len(tail)-o], tail[len(tail)-o-1], tail[len(tail)-1]
			var ok bool
			if o&1 == 1 {
				ok = cross(a0, a1, b0, pos)
			} else {
				ok = over(a0, a1, b0, pos)
			}
			if ok {
				return true
			}
		}
		tail = append(tail, pos)
	}
	return false
}
