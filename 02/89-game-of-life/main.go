package main

func main() {

}

func gameOfLife(board [][]int) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	N, M := len(board), len(board[0])
	get := func(i, j int) int {
		if i < 0 || i >= N || j < 0 || j >= M {
			return 0
		}
		return board[i][j] & 1
	}
	neighbors := func(i, j int) int {
		return get(i+1, j) +
			get(i+1, j+1) +
			get(i, j+1) +
			get(i-1, j+1) +
			get(i-1, j) +
			get(i-1, j-1) +
			get(i, j-1) +
			get(i+1, j-1)
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			board[i][j] = board[i][j] | (neighbors(i, j) << 1)
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			n, v := board[i][j]>>1, board[i][j]&1
			if n == 3 || (v == 1 && n == 2) {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}
