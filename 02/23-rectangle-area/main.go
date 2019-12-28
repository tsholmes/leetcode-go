package main

import "fmt"

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	fmt.Println(computeArea(-2, -2, 2, 2, -1, 4, 1, 6))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	return -max(-a, -b)
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	wid := min(C, G) - max(A, E)
	hei := min(D, H) - max(B, F)

	intA := 0
	if wid > 0 && hei > 0 {
		intA = wid * hei
	}

	return (C-A)*(D-B) + (G-E)*(H-F) - intA
}
