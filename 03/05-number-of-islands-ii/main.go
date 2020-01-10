package main

import "fmt"

func main() {
	fmt.Println(numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}}))
	fmt.Println(numIslands2(3, 3, [][]int{{0, 0}, {0, 1}, {1, 2}, {2, 1}, {1, 1}}))
}

func numIslands2(m int, n int, positions [][]int) []int {
	ps := map[[2]int][2]int{}
	roots := 0
	var (
		get  func([2]int) [2]int
		join func([2]int, [2]int)
	)

	get = func(v [2]int) [2]int {
		if _, ok := ps[v]; !ok {
			ps[v] = v
			roots++
			return v
		}
		p := ps[v]
		if v == p {
			return v
		}
		p = get(p)
		ps[v] = p
		return p
	}

	join = func(a [2]int, b [2]int) {
		pa, pb := get(a), get(b)
		if pa != pb {
			ps[pa] = pb
			roots--
		}
	}

	land := map[[2]int]bool{}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	res := make([]int, 0, len(positions))

	for _, p := range positions {
		k := [2]int{p[0], p[1]}
		land[k] = true
		get(k)
		for _, d := range dirs {
			k2 := [2]int{k[0] + d[0], k[1] + d[1]}
			if land[k2] {
				join(k, k2)
			}
		}
		res = append(res, roots)
	}

	return res
}
