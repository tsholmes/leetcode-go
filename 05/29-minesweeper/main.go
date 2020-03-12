package main

func main() {

}

func updateBoard(board [][]byte, click []int) [][]byte {
	N, M := len(board), len(board[0])
	dirs := [][2]int{
		{1, 0}, {1, 1}, {0, 1}, {-1, 1},
		{-1, 0}, {-1, -1}, {0, -1}, {1, -1},
	}
	count := func(i, j int) int {
		res := 0
		for _, d := range dirs {
			i2, j2 := i+d[0], j+d[1]
			if i2 >= 0 && i2 < N && j2 >= 0 && j2 < M {
				if board[i2][j2] == 'M' {
					res++
				}
			}
		}
		return res
	}
	var search func(i, j int, reveal bool)
	search = func(i, j int, reveal bool) {
		if i < 0 || i >= N || j < 0 || j >= M {
			return
		}
		switch board[i][j] {
		case 'M':
			if reveal {
				board[i][j] = 'X'
			}
			return
		case 'E':
			c := count(i, j)
			if c > 0 {
				board[i][j] = byte(c) + '0'
			} else {
				board[i][j] = 'B'
				for _, d := range dirs {
					search(i+d[0], j+d[1], false)
				}
			}
		}
	}
	search(click[0], click[1], true)
	return board
}
