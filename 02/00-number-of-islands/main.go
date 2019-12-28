package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		[]byte(`11110`),
		[]byte(`11010`),
		[]byte(`11000`),
		[]byte(`00000`),
	}))
}

func numIslands(grid [][]byte) int {
	parent := map[[2]int][2]int{}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				k := [2]int{i, j}
				parent[k] = k
			}
		}
	}
	var get func(k [2]int) [2]int
	get = func(k [2]int) [2]int {
		p := parent[k]
		if p == k {
			return p
		}
		parent[k] = get(p)
		return parent[k]
	}
	join := func(k1 [2]int, k2 [2]int) {
		p1, p2 := get(k1), get(k2)
		parent[p1] = p2
	}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != '1' {
				continue
			}
			for _, dir := range dirs {
				i2, j2 := i+dir[0], j+dir[1]
				if i2 < 0 || i2 >= len(grid) || j2 < 0 || j2 >= len(grid[0]) {
					continue
				}
				if grid[i2][j2] != '1' {
					continue
				}
				join([2]int{i, j}, [2]int{i2, j2})
			}
		}
	}

	distinct := map[[2]int]struct{}{}
	for k := range parent {
		distinct[get(k)] = struct{}{}
	}
	return len(distinct)
}
