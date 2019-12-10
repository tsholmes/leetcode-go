package main

import (
	"fmt"
	"math/bits"
)

func main() {
	do := func(n int) {
		res := solveNQueens(n)
		for _, b := range res {
			for _, s := range b {
				fmt.Println(s)
			}
			fmt.Println("---")
		}
		fmt.Println()
	}

	do(4)
	do(5)
}

func solveNQueens(n int) [][]string {
	work := make([]int, n)
	b := make([]byte, n)
	var search func(int)
	var res [][]string
	search = func(i int) {
		if i == n {
			sb := make([]string, n)
			for j := range b {
				for k := range b {
					if k == work[j] {
						b[k] = 'Q'
					} else {
						b[k] = '.'
					}
				}
				sb[j] = string(b)
			}
			res = append(res, sb)
			return
		}

		mask := 0
		for j := 0; j < i; j++ {
			bit := 1 << uint(work[j])
			mask |= bit
			mask |= bit << uint(i-j)
			mask |= bit >> uint(i-j)
		}

		mask = (1<<uint(n) - 1) &^ mask

		for mask > 0 {
			j := bits.TrailingZeros(uint(mask))
			bit := 1 << uint(j)
			work[i] = j
			mask = mask &^ bit
			search(i + 1)
		}
	}

	search(0)

	return res
}
