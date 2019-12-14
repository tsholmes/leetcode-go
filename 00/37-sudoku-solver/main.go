package main

import (
	"fmt"
	"math/bits"
)

func main() {
	do := func(board [][]byte) {
		solveSudoku(board)
		for _, line := range board {
			fmt.Println(string(line))
		}
		fmt.Println()
	}

	do([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
	do([][]byte{
		{'.', '.', '9', '7', '4', '8', '.', '.', '.'},
		{'7', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '2', '.', '1', '.', '9', '.', '.', '.'},
		{'.', '.', '7', '.', '.', '.', '2', '4', '.'},
		{'.', '6', '4', '.', '1', '.', '5', '9', '.'},
		{'.', '9', '8', '.', '.', '.', '3', '.', '.'},
		{'.', '.', '.', '8', '.', '3', '.', '2', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '6'},
		{'.', '.', '.', '2', '7', '5', '9', '.', '.'},
	})
}

func solveSudoku(board [][]byte) {
	solveSudokuRecursive(board, 3)
}

func solveSudokuRecursive(board [][]byte, maxGuesses int) int {
	orig := board
	board = clone(board)

	lastMissing := 0
	for {
		rows := make([]int, 9)
		cols := make([]int, 9)
		boxes := make([]int, 9)

		// build up in-use bits and fill in with only one option
		for r := range board {
			for c := range board[r] {
				if board[r][c] == '.' {
					continue
				}
				bi := (r/3)*3 + c/3

				bit := 1 << uint(board[r][c]-'1')
				rows[r] |= bit
				cols[c] |= bit
				boxes[bi] |= bit
			}
		}

		missing := 0
		for r := range board {
			for c := range board[r] {
				if board[r][c] != '.' {
					continue
				}
				bi := (r/3)*3 + c/3

				mask := 0x1ff &^ (rows[r] | cols[c] | boxes[bi])

				if bits.OnesCount(uint(mask)) == 1 {
					num := bits.TrailingZeros(uint(mask))
					board[r][c] = '1' + byte(num)
					rows[r] |= mask
					cols[c] |= mask
					boxes[bi] |= mask
				} else {
					missing++
				}
			}
		}
		if missing == 0 {
			copyBack(orig, board)
			return 0
		} else if missing == lastMissing {
			// give up on this path
			if maxGuesses == 0 {
				return missing
			}

			for r := range board {
				for c := range board[r] {
					if board[r][c] != '.' {
						continue
					}
					bi := (r/3)*3 + c/3

					mask := 0x1ff &^ (rows[r] | cols[c] | boxes[bi])
					if bits.OnesCount(uint(mask)) == 2 {
						g1 := bits.TrailingZeros(uint(mask))
						g2 := bits.TrailingZeros(uint(mask &^ (1 << uint(g1))))

						board[r][c] = byte(g1) + '1'
						if solveSudokuRecursive(board, maxGuesses-1) == 0 {
							copyBack(orig, board)
							return 0
						}
						board[r][c] = byte(g2) + '1'
						if solveSudokuRecursive(board, maxGuesses-1) == 0 {
							copyBack(orig, board)
							return 0
						}
						board[r][c] = '.'
					}
				}
			}

			return missing
		}
		lastMissing = missing

		// fill rows that have only one option
		for r := range board {
			for d := 0; d < 9; d++ {
				mask := 1 << uint(d)
				if rows[r]&mask != 0 {
					continue
				}
				var ok int
				for c := range board[r] {
					if board[r][c] != '.' {
						continue
					}
					bit := 1 << uint(c)

					bi := (r/3)*3 + c/3
					if (rows[r]|cols[c]|boxes[bi])&mask == 0 {
						ok |= bit
					}
				}
				if bits.OnesCount(uint(ok)) == 1 {
					c := bits.TrailingZeros(uint(ok))
					bi := (r/3)*3 + c/3

					board[r][c] = '1' + byte(d)
					rows[r] |= mask
					cols[c] |= mask
					boxes[bi] |= mask
				}
			}
		}

		// fill cols that have only one option
		for c := range board[0] {
			for d := 0; d < 9; d++ {
				mask := 1 << uint(d)
				if cols[c]&mask != 0 {
					continue
				}
				var ok int
				for r := range board {
					if board[r][c] != '.' {
						continue
					}
					bit := 1 << uint(r)

					bi := (r/3)*3 + c/3
					if (rows[r]|cols[c]|boxes[bi])&mask == 0 {
						ok |= bit
					}
				}
				if bits.OnesCount(uint(ok)) == 1 {
					r := bits.TrailingZeros(uint(ok))
					bi := (r/3)*3 + c/3

					board[r][c] = '1' + byte(d)
					rows[r] |= mask
					cols[c] |= mask
					boxes[bi] |= mask
				}
			}
		}

		// fill in boxes that have only one option
		for bi := range board {
			br, bc := bi/3, bi%3
			for d := 0; d < 9; d++ {
				mask := 1 << uint(d)
				if boxes[bi]&mask != 0 {
					continue
				}
				var ok int
				for bj := range board {
					bit := 1 << uint(bj)
					r := br*3 + bj/3
					c := bc*3 + bj%3
					if board[r][c] != '.' {
						continue
					}
					if (rows[r]|cols[c]|boxes[bi])&mask == 0 {
						ok |= bit
					}
				}
				if bits.OnesCount(uint(ok)) == 1 {
					bj := bits.TrailingZeros(uint(ok))
					r := br*3 + bj/3
					c := bc*3 + bj%3

					board[r][c] = '1' + byte(d)
					rows[r] |= mask
					cols[c] |= mask
					boxes[bi] |= mask
				}
			}
		}
	}
}

func clone(board [][]byte) [][]byte {
	b2 := make([][]byte, 9)
	for r := range b2 {
		b2[r] = make([]byte, 9)
		copy(b2[r], board[r])
	}
	return b2
}

func copyBack(orig [][]byte, clone [][]byte) {
	for r := range orig {
		copy(orig[r], clone[r])
	}
}
