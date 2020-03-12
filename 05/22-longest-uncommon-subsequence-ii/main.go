package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findLUSlength([]string{"aba", "cdc", "eae"}))
	fmt.Println(findLUSlength([]string{"aabbcc", "aabbcc", "c"}))
}

func findLUSlength(strs []string) int {
	byLen := map[int][]string{}
	counts := map[string]int{}
	for _, s := range strs {
		byLen[len(s)] = append(byLen[len(s)], s)
		counts[s]++
	}

	var lens []int
	for k := range byLen {
		lens = append(lens, k)
	}
	sort.Ints(lens)

	for j := len(lens) - 1; j >= 0; j-- {
		for _, w := range byLen[lens[j]] {
			if counts[w] > 1 {
				continue
			}
			valid := true
			for k := j + 1; k < len(lens) && valid; k++ {
				for _, w2 := range byLen[lens[k]] {
					if isSubSeq(w, w2) {
						valid = false
						break
					}
				}
			}
			if valid {
				return len(w)
			}
		}
	}

	return -1
}

func isSubSeq(a string, b string) bool {
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] == b[j] {
			i++
			j++
		} else {
			j++
		}
	}
	return i == len(a)
}
