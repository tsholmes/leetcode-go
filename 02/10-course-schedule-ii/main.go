package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(findOrder(2, [][]int{{1, 0}}))
	fmt.Println(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	pres := make([][]int, numCourses)
	for _, pair := range prerequisites {
		pres[pair[0]] = append(pres[pair[0]], pair[1])
	}

	depths := make([]int, numCourses)
	var walk func(int, int, map[int]bool) bool
	walk = func(i int, depth int, seen map[int]bool) bool {
		if seen[i] {
			return false
		}
		seen[i] = true
		defer func() {
			seen[i] = false
		}()

		if depth > 0 && depth <= depths[i] {
			return true
		}
		depths[i] = depth

		for _, j := range pres[i] {
			if !walk(j, depth+1, seen) {
				return false
			}
		}
		return true
	}

	for i := 0; i < numCourses; i++ {
		if depths[i] == 0 && !walk(i, 0, map[int]bool{}) {
			return nil
		}
	}

	res := make([]int, numCourses)
	for i := range res {
		res[i] = i
	}
	sort.Slice(res, func(i, j int) bool {
		di, dj := depths[res[i]], depths[res[j]]
		return di > dj
	})

	return res
}
