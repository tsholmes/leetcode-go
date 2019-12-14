package main

import "fmt"

func main() {
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCCED"))
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "SEE"))
	fmt.Println(exist([][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}, "ABCB"))
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	var search func(int, [2]int) bool
	search = func(i int, pos [2]int) bool {
		if i == len(word) {
			return true
		}
		cc := board[pos[0]][pos[1]]
		board[pos[0]][pos[1]] = 0
		c := word[i]
		for _, dir := range dirs {
			p2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
			if p2[0] < 0 || p2[0] >= len(board) || p2[1] < 0 || p2[1] >= len(board[0]) {
				continue
			}
			if board[p2[0]][p2[1]] == c && search(i+1, p2) {
				return true
			}
		}
		board[pos[0]][pos[1]] = cc
		return false
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && search(1, [2]int{i, j}) {
				return true
			}
		}
	}
	return false
}
