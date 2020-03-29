package main

import "fmt"

func main() {
	fmt.Println(hasValidPath([][]int{
		{2, 4, 3},
		{6, 5, 2},
	}))
}

func hasValidPath(grid [][]int) bool {
	ps := map[[2]int][2]int{}
	var (
		get  func([2]int) [2]int
		join func([2]int, [2]int)
	)
	get = func(p [2]int) [2]int {
		if _, ok := ps[p]; !ok {
			ps[p] = p
		}
		pp := ps[p]
		if pp != p {
			pp = get(pp)
			ps[p] = pp
		}
		return pp
	}
	join = func(a, b [2]int) {
		pa, pb := get(a), get(b)
		ps[pa] = pb
	}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	/*
	  1 which means a street connecting the left cell and the right cell.
	  2 which means a street connecting the upper cell and the lower cell.
	  3 which means a street connecting the left cell and the lower cell.
	  4 which means a street connecting the right cell and the lower cell.
	  5 which means a street connecting the left cell and the upper cell.
	  6 which means a street connecting the right cell and the upper cell.
	*/
	// down, right, up, left
	conns := map[int][4]bool{
		1: {false, true, false, true},
		2: {true, false, true, false},
		3: {true, false, false, true},
		4: {true, true, false, false},
		5: {false, false, true, true},
		6: {false, true, true, false},
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			p := [2]int{i, j}
			cs := conns[grid[i][j]]
			for di, d := range dirs {
				if !cs[di] {
					continue
				}
				p2 := [2]int{i + d[0], j + d[1]}
				if p2[0] < 0 || p2[0] >= len(grid) || p2[1] < 0 || p2[1] >= len(grid[0]) {
					continue
				}
				c2 := conns[grid[p2[0]][p2[1]]]
				if !c2[di^2] {
					continue
				}
				join(p, p2)
			}
		}
	}

	return get([2]int{0, 0}) == get([2]int{len(grid) - 1, len(grid[0]) - 1})
}
