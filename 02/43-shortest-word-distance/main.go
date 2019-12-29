package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestDistance([]string{
		"practice", "makes", "perfect", "coding", "makes",
	}, "coding", "practice"))
	fmt.Println(shortestDistance([]string{
		"practice", "makes", "perfect", "coding", "makes",
	}, "makes", "coding"))
}

func shortestDistance(words []string, word1 string, word2 string) int {
	min := math.MaxInt64
	last1 := math.MinInt32
	last2 := math.MinInt32
	for i, w := range words {
		if w == word1 {
			if d := i - last2; d < min {
				min = d
			}
			last1 = i
		} else if w == word2 {
			if d := i - last1; d < min {
				min = d
			}
			last2 = i
		}
	}
	return min
}
