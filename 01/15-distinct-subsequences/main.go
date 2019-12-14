package main

import "fmt"

func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}

func numDistinct(s string, t string) int {
	memo := map[[2]int]int{}
	var search func(int, int) int
	search = func(i, j int) int {
		if j == len(t) {
			return 1
		}
		if i == len(s) {
			return 0
		}
		key := [2]int{i, j}
		v, ok := memo[key]
		if ok {
			return v
		}
		for k := i; k < len(s); k++ {
			if s[k] == t[j] {
				v += search(k+1, j+1)
			}
		}
		memo[key] = v
		return v
	}
	return search(0, 0)
}
