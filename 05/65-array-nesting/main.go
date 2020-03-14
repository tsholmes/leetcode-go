package main

import "fmt"

func main() {
	fmt.Println(arrayNesting([]int{5, 4, 0, 3, 1, 6, 2}))
}

func arrayNesting(nums []int) int {
	max := 0

	seen := map[int]bool{}
	var path []int
	var searchRing func(i int)
	searchRing = func(i int) {
		if !seen[i] {
			seen[i] = true
			path = append(path, i)
			searchRing(nums[i])
			return
		}
		rlen := 1
		for path[len(path)-rlen] != i {
			rlen++
		}
		if rlen > max {
			max = rlen
		}
	}

	for i := range nums {
		if !seen[i] {
			path = path[:0]
			searchRing(i)
		}
	}

	return max
}
