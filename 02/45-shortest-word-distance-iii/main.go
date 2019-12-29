package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(shortestWordDistance([]string{
		"practice", "makes", "perfect", "coding", "makes",
	}, "makes", "coding"))
	fmt.Println(shortestWordDistance([]string{
		"practice", "makes", "perfect", "coding", "makes",
	}, "makes", "makes"))
}

func shortestWordDistance(words []string, word1 string, word2 string) int {
	min := math.MaxInt64
	last1 := math.MinInt32
	last2 := math.MinInt32
	if word1 == word2 {
		for i, w := range words {
			if w == word1 {
				last1, last2 = last2, i
				if i-last1 < min {
					min = i - last1
				}
			}
		}
	} else {
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
	}
	return min
}
