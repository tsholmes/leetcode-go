package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 6, 1, 3}))
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
}

func verifyPreorder(preorder []int) bool {
	var search func(int, int, int) bool
	search = func(i, j, min int) bool {
		if j == i {
			return true
		}
		if j == i+1 {
			return preorder[i] > min
		}
		root := preorder[i]
		if root < min {
			return false
		}
		end := i + 1
		for end < j && preorder[end] < root {
			end++
		}
		return search(i+1, end, min) && search(end, j, root)
	}
	return search(0, len(preorder), math.MinInt64)
}
