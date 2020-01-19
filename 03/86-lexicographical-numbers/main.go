package main

import "fmt"

func main() {
	fmt.Println(lexicalOrder(120))
}

func lexicalOrder(n int) []int {
	res := make([]int, 0, n)
	var search func(int, int)
	search = func(v int, pow int) {
		if v > n {
			return
		}
		res = append(res, v)
		search(v*10, pow*10)
		if pow > 1 && v%10 < 9 {
			search(v+1, pow)
		}
	}
	for i := 1; i <= 9; i++ {
		search(i, 1)
	}
	return res
}
