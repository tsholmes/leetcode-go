package main

import "fmt"

func main() {
	fmt.Println(isUgly(6))
	fmt.Println(isUgly(8))
	fmt.Println(isUgly(14))
}

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	for num&1 == 0 {
		num = num >> 1
	}
	for num%3 == 0 {
		num = num / 3
	}
	for num%5 == 0 {
		num = num / 5
	}
	return num == 1
}
