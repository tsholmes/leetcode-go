package main

import "fmt"

func main() {
	fmt.Println(minFlips(2, 6, 5))
	fmt.Println(minFlips(4, 2, 7))
	fmt.Println(minFlips(1, 2, 3))
	fmt.Println(minFlips(8, 3, 5))
}
func minFlips(a int, b int, c int) int {
	count := 0
	for bb := 0; bb < 32; bb++ {
		bit := 1 << uint(bb)
		ai, bi, ci := (a & bit), (b & bit), (c & bit)
		if (ai | bi) != ci {
			if ci == 0 {
				if ai != 0 {
					count++
				}
				if bi != 0 {
					count++
				}
			} else {
				count += 1
			}
		}
	}
	return count
}
