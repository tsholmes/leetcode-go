package main

func main() {

}

func findBlackPixel(picture [][]byte, N int) int {
	R, C := len(picture), len(picture[0])
	rows := make([]int, R)
	cols := make([]int, C)
	type rowKey struct {
		row string
		col int
	}
	rowCounts := map[rowKey]int{}
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
				rowCounts[rowKey{string(picture[i]), j}]++
			}
		}
	}
	res := 0
	for i := 0; i < R; i++ {
		if rows[i] != N {
			continue
		}
		for j := 0; j < C; j++ {
			if cols[j] == N && rowCounts[rowKey{string(picture[i]), j}] == N {
				res++
			}
		}
	}
	return res
}
