package main

func main() {

}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	N, M := len(matrix), len(matrix[0])
	counts := make([][]int, N)
	for i := range counts {
		counts[i] = make([]int, M)
	}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	search := func(queue [][2]int) {
		seen := map[[2]int]bool{}
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			if seen[p] {
				continue
			}
			seen[p] = true
			counts[p[0]][p[1]]++
			for _, d := range dirs {
				p2 := [2]int{p[0] + d[0], p[1] + d[1]}
				if p2[0] < 0 || p2[0] >= N || p2[1] < 0 || p2[1] >= M {
					continue
				}
				if matrix[p2[0]][p2[1]] >= matrix[p[0]][p[1]] {
					queue = append(queue, p2)
				}
			}
		}
	}
	var pac, atl [][2]int
	for i := 0; i < N; i++ {
		pac = append(pac, [2]int{i, 0})
		atl = append(atl, [2]int{i, M - 1})
	}
	for j := 0; j < M; j++ {
		pac = append(pac, [2]int{0, j})
		atl = append(atl, [2]int{N - 1, j})
	}
	search(pac)
	search(atl)
	var res [][]int
	for i := range counts {
		for j := range counts[i] {
			if counts[i][j] >= 2 {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
