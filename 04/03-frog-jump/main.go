package main

import "fmt"

func main() {
	fmt.Println(canCross([]int{0, 1, 3, 4, 5, 7, 9}))
	fmt.Println(canCross([]int{0, 1, 3, 6, 10, 15, 16, 21}))
}

func canCross(stones []int) bool {
	if stones[1] != 1 {
		return false
	}
	seen := map[[2]int]bool{}
	var search func(p int, k int, i int) bool
	search = func(p int, k int, i int) bool {
		if i == len(stones)-1 {
			return true
		}
		key := [2]int{p, k}
		if seen[key] {
			return false
		}
		seen[key] = true

		i++
		for i < len(stones) && stones[i] <= p+k+1 {
			if stones[i] >= p+k-1 && search(stones[i], stones[i]-p, i) {
				return true
			}
			i++
		}

		return false
	}
	return search(1, 1, 1)
}
