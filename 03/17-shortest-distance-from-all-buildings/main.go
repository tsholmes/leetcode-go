package main

import "math"

func main() {

}

func shortestDistance(grid [][]int) int {
	type stateKey struct {
		source [2]int
		pos    [2]int
	}
	type state struct {
		key  stateKey
		dist int
	}
	var houses int
	var queue []state

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				houses++
				p := [2]int{i, j}
				queue = append(queue, state{
					key:  stateKey{source: p, pos: p},
					dist: 0,
				})
			}
		}
	}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

	seenBy := map[[2]int]int{}
	seen := map[stateKey]bool{}
	dists := map[[2]int]int{}
	for len(queue) > 0 {
		s := queue[0]
		queue = queue[1:]

		if seen[s.key] {
			continue
		}
		seen[s.key] = true
		seenBy[s.key.pos]++
		dists[s.key.pos] += s.dist

		for _, d := range dirs {
			p2 := [2]int{s.key.pos[0] + d[0], s.key.pos[1] + d[1]}
			if p2[0] < 0 || p2[0] >= len(grid) || p2[1] < 0 || p2[1] >= len(grid[0]) || grid[p2[0]][p2[1]] != 0 {
				continue
			}
			queue = append(queue, state{
				key:  stateKey{source: s.key.source, pos: p2},
				dist: s.dist + 1,
			})
		}
	}

	minDist := math.MaxInt64

	for i := range grid {
		for j := range grid[i] {
			k := [2]int{i, j}
			if grid[i][j] == 0 && seenBy[k] == houses && dists[k] < minDist {
				minDist = dists[k]
			}
		}
	}

	if minDist == math.MaxInt64 {
		return -1
	}

	return minDist
}
