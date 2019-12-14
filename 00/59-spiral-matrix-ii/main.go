package main

import "fmt"

func main() {
	do := func(n int) {
		res := generateMatrix(n)
		for _, row := range res {
			fmt.Println(row)
		}
		fmt.Println()
	}
	do(0)
	do(1)
	do(3)
	do(6)
}

func generateMatrix(n int) [][]int {
	if n == 0 {
		return nil
	}
	pos := [2]int{0, 0}
	dir := [2]int{0, 1}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for i := 1; i <= n*n; i++ {
		res[pos[0]][pos[1]] = i

		pos2 := [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		if pos2[0] < 0 || pos2[0] == n || pos2[1] < 0 || pos2[1] == n || res[pos2[0]][pos2[1]] != 0 {
			dir = [2]int{dir[1], -dir[0]}
			pos2 = [2]int{pos[0] + dir[0], pos[1] + dir[1]}
		}
		pos = pos2
	}

	return res
}
