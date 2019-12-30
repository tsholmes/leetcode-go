package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(i, addDigits(i))
	}
}

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	return (num-1)%9 + 1
}
