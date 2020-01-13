package main

func main() {

}

func maxKilledEnemies(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	N, M := len(grid), len(grid[0])
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	max := 0
	for i := range grid {
		for j := range grid[0] {
			if grid[i][j] != '0' {
				continue
			}
			count := 0
			for _, d := range dirs {
				p := [2]int{i, j}
				for {
					p = [2]int{p[0] + d[0], p[1] + d[1]}
					if p[0] < 0 || p[0] >= N || p[1] < 0 || p[1] >= M || grid[p[0]][p[1]] == 'W' {
						break
					}
					if grid[p[0]][p[1]] == 'E' {
						count++
					}
				}
			}
			if count > max {
				max = count
			}
		}
	}
	return max
}
