package main

func main() {

}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	memo := map[[2]int]int{}
	var search func([2]int) int
	search = func(pos [2]int) int {
		if v, ok := memo[pos]; ok {
			return v
		}
		max := 0
		for _, d := range dirs {
			p2 := [2]int{pos[0] + d[0], pos[1] + d[1]}
			if p2[0] < 0 || p2[0] >= len(matrix) || p2[1] < 0 || p2[1] >= len(matrix[0]) {
				continue
			}
			if matrix[p2[0]][p2[1]] <= matrix[pos[0]][pos[1]] {
				continue
			}
			if v := search(p2); v > max {
				max = v
			}
		}
		memo[pos] = 1 + max
		return 1 + max
	}

	max := 0
	for i := range matrix {
		for j := range matrix[i] {
			if v := search([2]int{i, j}); v > max {
				max = v
			}
		}
	}
	return max
}
