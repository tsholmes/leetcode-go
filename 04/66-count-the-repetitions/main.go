package main

import (
	"fmt"
)

func main() {
	fmt.Println(getMaxRepetitions("acb", 4, "ab", 2))
	fmt.Println(getMaxRepetitions("abab", 4, "ababab", 2))
	fmt.Println(getMaxRepetitions("bacaba", 3, "abacab", 1))
	fmt.Println(getMaxRepetitions("baba", 11, "baab", 1))
	fmt.Println(getMaxRepetitions("nlhqgllunmelayl", 2, "lnl", 1))
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	if n1 == 0 {
		return 0
	}

	seen := map[[3]int]int{}
	var counts []int
	var nexts []int

	{
		i, j := 0, 0
		cur := 0
		iter := 0
		for {
			// simple if slow detection of no match
			if iter > 1e7 {
				return 0
			}
			iter++
			key := [3]int{i, j, cur}
			if i2, ok := seen[key]; ok {
				if len(nexts) == 0 {
					return 0
				}
				nexts[len(nexts)-1] = i2
				break
			}
			seen[key] = len(counts)

			if s1[i] != s2[j] {
				i++
			} else {
				j++
				i++
			}
			if j == len(s2) {
				j = 0
				cur++
			}
			if i == len(s1) {
				i = 0
				counts = append(counts, cur)
				nexts = append(nexts, len(counts))
				cur = 0
			}
		}
	}

	count := 0
	pos := 0
	for i := 0; i < n1; i++ {
		count += counts[pos]
		pos = nexts[pos]
	}

	return count / n2
}
