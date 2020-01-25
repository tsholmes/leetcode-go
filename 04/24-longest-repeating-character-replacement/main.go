package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(characterReplacement("ABAB", 2))
	fmt.Println(characterReplacement("AABABBA", 1))
	fmt.Println(characterReplacement("BAAAB", 2))
}

func characterReplacement(s string, k int) int {
	counts := make([][26]int, len(s))
	max := 0
	for i := 0; i < len(s); i++ {
		ix := int(s[i] - 'A')
		for j := 0; j < 26; j++ {
			if i > 0 {
				counts[i][j] += counts[i-1][j]
			}
			if j != ix {
				counts[i][j]++
			}
		}
		for ix := 0; ix < 26; ix++ {
			j := sort.Search(i+1, func(j int) bool {
				count := counts[i][ix]
				if j > 0 {
					count -= counts[j-1][ix]
				}
				return count <= k
			})
			if i-j+1 > max {
				max = i - j + 1
			}
		}
	}
	return max
}
