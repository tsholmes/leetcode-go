package main

func main() {

}

func findLonelyPixel(picture [][]byte) int {
	N, M := len(picture), len(picture[0])
	rows := make([]int, N)
	cols := make([]int, M)
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if picture[i][j] == 'B' {
				rows[i]++
				cols[j]++
			}
		}
	}
	res := 0
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if picture[i][j] == 'B' && rows[i] == 1 && cols[j] == 1 {
				res++
			}
		}
	}
	return res
}
