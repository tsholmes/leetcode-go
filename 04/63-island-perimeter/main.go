package main

func main() {

}

func islandPerimeter(grid [][]int) int {
	count := 0
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] != 1 {
				continue
			}
			for _, d := range dirs {
				i2, j2 := i+d[0], j+d[1]
				if i2 < 0 || i2 >= len(grid) || j2 < 0 || j2 >= len(grid[i]) || grid[i2][j2] == 0 {
					count++
				}
			}
		}
	}
	return count
}
