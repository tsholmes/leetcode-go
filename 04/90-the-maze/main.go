package main

func main() {

}

func hasPath(maze [][]int, start []int, destination []int) bool {
	N, M := len(maze), len(maze[0])

	dirs := [4][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	roll := func(p [2]int, dir int) [2]int {
		d := dirs[dir]
		for {
			p2 := [2]int{p[0] + d[0], p[1] + d[1]}
			if p2[0] < 0 || p2[0] >= N || p2[1] < 0 || p2[1] >= M || maze[p2[0]][p2[1]] == 1 {
				break
			}
			p = p2
		}
		return p
	}

	seen := map[[2]int]bool{}
	var search func([2]int) bool
	search = func(p [2]int) bool {
		if p == [2]int{destination[0], destination[1]} {
			return true
		}
		if seen[p] {
			return false
		}
		seen[p] = true
		for d := 0; d < 4; d++ {
			if search(roll(p, d)) {
				return true
			}
		}
		return false
	}
	return search([2]int{start[0], start[1]})
}
