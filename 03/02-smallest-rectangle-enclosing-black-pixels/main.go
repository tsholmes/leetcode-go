package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minArea([][]byte{
		[]byte(`0010`),
		[]byte(`0110`),
		[]byte(`0100`),
	}, 0, 2))
	fmt.Println(minArea([][]byte{
		[]byte(`01`),
	}, 0, 1))
}

func minArea(image [][]byte, x int, y int) int {
	N, M := len(image), len(image[0])

	left := sort.Search(x, func(i int) bool {
		any := false
		for j := 0; j < M; j++ {
			any = any || image[i][j] == '1'
		}
		return any
	})
	right := x + sort.Search(N-x, func(i int) bool {
		any := false
		for j := 0; j < M; j++ {
			any = any || image[i+x][j] == '1'
		}
		return !any
	}) - 1
	top := sort.Search(y, func(j int) bool {
		any := false
		for i := 0; i < N; i++ {
			any = any || image[i][j] == '1'
		}
		return any
	})
	bottom := y + sort.Search(M-y, func(j int) bool {
		any := false
		for i := 0; i < N; i++ {
			any = any || image[i][j+y] == '1'
		}
		return !any
	}) - 1
	return (right - left + 1) * (bottom - top + 1)
}
