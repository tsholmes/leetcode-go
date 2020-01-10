package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findMinHeightTrees(4, [][]int{{1, 0}, {1, 2}, {1, 3}}))
	fmt.Println(findMinHeightTrees(6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}))
	fmt.Println(findMinHeightTrees(2, [][]int{{0, 1}}))
	fmt.Println(findMinHeightTrees(1, [][]int{}))
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return nil
	}
	es := map[int][]int{}
	for _, e := range edges {
		es[e[0]] = append(es[e[0]], e[1])
		es[e[1]] = append(es[e[1]], e[0])
	}

	ns := make([]int, n)
	for i := range ns {
		ns[i] = i
	}

	exclude := map[int]bool{}
	for {
		count := 0
		for i := 0; i < len(ns); i++ {
			ec := 0
			for _, e := range es[ns[i]] {
				if !exclude[e] {
					ec++
				}
			}
			if ec <= 1 {
				ns[i], ns[count] = ns[count], ns[i]
				count++
			}
		}
		if count == len(ns) {
			sort.Ints(ns)
			return ns
		}
		for _, n := range ns[:count] {
			exclude[n] = true
		}
		ns = ns[count:]
	}
}
