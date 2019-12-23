package main

import "fmt"

func main() {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("AB"))
	fmt.Println(titleToNumber("ZY"))
}

func titleToNumber(s string) int {
	base := 0
	mul := 1
	for i := 1; i < len(s); i++ {
		mul *= 26
		base += mul
	}
	for i := 0; i < len(s); i++ {
		base += int(s[i]-'A') * mul
		mul /= 26
	}
	return base + 1
}
