package main

import (
	"fmt"
	"math/bits"
)

func main() {
	for i := 1; i < 16; i++ {
		fmt.Println(totalNQueens(i))
	}
}

func totalNQueens(n int) int {
	work := make([]int, n)
	var search func(int)
	res := 0
	search = func(i int) {
		if i == n {
			res++
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
