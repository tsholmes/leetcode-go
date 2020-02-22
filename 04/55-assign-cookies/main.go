package main

import "sort"

func main() {

}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	i, j := len(g)-1, len(s)-1
	count := 0
	for i >= 0 && j >= 0 {
		if s[j] >= g[i] {
			count++
			j--
		}
		i--
	}

	return count
}
