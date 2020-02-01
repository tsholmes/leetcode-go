package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(arrangeCoins(3))
}

func arrangeCoins(n int) int {
	i := sort.Search(n, func(i int) bool {
		return i*(i+1)/2 >= n
	})
	if i*(i+1)/2 > n {
		i--
	}
	return i
}
