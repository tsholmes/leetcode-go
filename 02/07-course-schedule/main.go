package main

import "fmt"

func main() {
	fmt.Println(canFinish(2, [][]int{
		{1, 0},
	}))
	fmt.Println(canFinish(2, [][]int{
		{0, 1},
		{1, 0},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	pres := map[int][]int{}
	for _, pair := range prerequisites {
		pres[pair[0]] = append(pres[pair[0]], pair[1])
	}

	seen := map[int]bool{}
	var walk func(int) bool
	walk = func(i int) bool {
		if seen[i] {
			return true
		}
		seen[i] = true
		res := false
		for _, n := range pres[i] {
			if walk(n) {
				res = true
				break
			}
		}
		seen[i] = false
		return res
	}

	for i := 0; i < numCourses; i++ {
		if walk(i) {
			return false
		}
	}
	return true
}
