package main

import "fmt"

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-2, 3))
	fmt.Println(getSum(3, -2))
}

func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
