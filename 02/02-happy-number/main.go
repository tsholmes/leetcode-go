package main

import "fmt"

func main() {
	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))
}

func isHappy(n int) bool {
	seen := map[int]bool{}
	for n > 1 {
		if seen[n] {
			return false
		}
		seen[n] = true
		next := 0
		for n > 0 {
			d := n % 10
			n /= 10
			next += d * d
		}
		n = next
	}
	return true
}
