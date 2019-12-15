package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
}

func longestConsecutive(nums []int) int {
	parent := map[int]int{}
	get := func(a int) (int, bool) {
		b, ok := parent[a]
		if !ok {
			return 0, false
		}
		for a != b {
			parent[a] = parent[b]
			a, b = b, parent[b]
		}
		return b, true
	}
	join := func(a, b int) {
		pa, oka := get(a)
		pb, okb := get(b)
		if oka && okb {
			parent[pa] = parent[pb]
		}
	}

	for _, n := range nums {
		if _, ok := parent[n]; !ok {
			parent[n] = n
		}
		join(n, n-1)
		join(n, n+1)
	}

	counts := map[int]int{}
	for k := range parent {
		p, _ := get(k)
		counts[p]++
	}

	max := 0
	for _, v := range counts {
		if v > max {
			max = v
		}
	}

	return max
}
