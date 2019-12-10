package main

import "fmt"

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9, 9, 9, 9, 9}))
}

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0; i-- {
		v := digits[i] + carry
		digits[i] = v % 10
		carry = v / 10
	}
	if carry == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
