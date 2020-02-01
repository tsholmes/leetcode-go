package main

import "fmt"

func main() {
	for i := 1; i <= 13; i++ {
		fmt.Println(findKthNumber(13, i))
	}
}

func findKthNumber(n int, k int) int {
	countPrefix := func(p int) int {
		res := 0
		p1 := p + 1
		for p <= n {
			upper := p1
			if upper > n+1 {
				upper = n + 1
			}
			res += upper - p
			p *= 10
			p1 *= 10
		}
		return res
	}
	var search func(p int, pk int) int
	search = func(p int, pk int) int {
		if pk == 0 {
			return p
		}
		count := countPrefix(p)
		if count <= pk {
			return search(p+1, pk-count)
		}
		return search(p*10, pk-1)
	}
	return search(1, k-1)
}
