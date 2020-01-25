package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(wordsTyping([]string{"hello", "world"}, 2, 8))
	fmt.Println(wordsTyping([]string{"a", "bcd", "e"}, 3, 6))
	fmt.Println(wordsTyping([]string{"I", "had", "apple", "pie"}, 4, 5))
	fmt.Println(wordsTyping([]string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}, 20000, 20000))
}

func wordsTyping(sentence []string, rows int, cols int) int {
	N := len(sentence)
	lens := make([][]int, N)
	for i := range lens {
		lens[i] = make([]int, N)
		tlen := 0
		for j := i; j < N; j++ {
			tlen += len(sentence[j]) + 1
			lens[i][j] = tlen
		}
	}

	var res int
	wpos := 0
	for row := 0; row < rows; row++ {
		used := -1
		for lens[wpos][N-1]+used <= cols {
			used += lens[wpos][N-1]
			wpos = 0
			res++
		}
		wpos = sort.Search(N, func(i int) bool {
			return used+lens[wpos][i] > cols
		})
	}

	return res
}
