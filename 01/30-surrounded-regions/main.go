package main

import "fmt"

func main() {
	do := func(board [][]byte) {
		solve(board)
		for _, line := range board {
			fmt.Println(string(line))
		}
	}
	do([][]byte{
		[]byte(`XXXX`),
		[]byte(`XOOX`),
		[]byte(`XXOX`),
		[]byte(`XOXX`),
	})
	do([][]byte{
		[]byte(`XOXOXO`),
		[]byte(`OXOXOX`),
		[]byte(`XOXOXO`),
		[]byte(`OXOXOX`),
	})
}

func solve(board [][]byte) {
	parent := map[[2]int][2]int{}
	get := func(a [2]int) ([2]int, bool) {
		b, ok := parent[a]
		if !ok {
			return [2]int{}, false
		}
		for a != b {
			parent[a] = parent[b]
			a, b = b, parent[b]
		}
		return b, true
	}
	join := func(a, b [2]int) {
		pa, oka := get(a)
		pb, okb := get(b)
		if oka && okb {
			parent[pa] = parent[pb]
		}
	}

	parent[[2]int{-1, -1}] = [2]int{-1, -1}

	dirs := [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			key := [2]int{i, j}
			if board[i][j] == 'X' {
				continue
			}
			if _, ok := parent[key]; !ok {
				parent[key] = key
			}
			for _, dir := range dirs {
				k2 := [2]int{i + dir[0], j + dir[1]}
				if k2[0] < 0 || k2[0] >= len(board) || k2[1] < 0 || k2[1] >= len(board[0]) {
					k2 = [2]int{-1, -1}
				} else if board[k2[0]][k2[1]] == 'X' {
					continue
				}
				join(key, k2)
			}
		}
	}

	po, _ := get([2]int{-1, -1})

	for k := range parent {
		p, _ := get(k)
		if p == po {
			continue
		}
		board[k[0]][k[1]] = 'X'
	}
}
