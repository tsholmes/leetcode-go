package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, canWinNim(i))
	}
}

func canWinNim(n int) bool {
	return n%4 != 0
}
