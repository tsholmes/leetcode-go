package main

func main() {

}

type TicTacToe struct {
	rows  []int
	cols  []int
	diag1 int
	diag2 int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	return TicTacToe{
		rows: make([]int, n),
		cols: make([]int, n),
	}
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (t *TicTacToe) Move(row int, col int, player int) int {
	off := (player * 2) - 3
	N := len(t.rows)
	win := 0

	t.rows[row] += off
	if t.rows[row] == N*off {
		win = player
	}

	t.cols[col] += off
	if t.cols[col] == N*off {
		win = player
	}

	if row == col {
		t.diag1 += off
		if t.diag1 == N*off {
			win = player
		}
	}

	if row == N-col-1 {
		t.diag2 += off
		if t.diag2 == N*off {
			win = player
		}
	}

	return win
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
