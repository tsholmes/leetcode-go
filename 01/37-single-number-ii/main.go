package main

import "fmt"

func main() {
	fmt.Println(singleNumber([]int{2, 2, 3, 2}))
	fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, -3, 1, -3, -3, -4, -2}))
}

func singleNumber(nums []int) int {
	m1 := 0
	m2 := 0
	for _, n := range nums {
		m1, m2 = m1^(n&^m2), (m2&^n)|(m1&n)
	}
	return m1
}
