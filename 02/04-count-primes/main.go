package main

import "fmt"

func main() {
	fmt.Println(countPrimes(10))
	fmt.Println(countPrimes(100))
}

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	invalid := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if invalid[i] {
			continue
		}
		count++
		for j := i * i; j < n; j += i {
			invalid[j] = true
		}
	}
	return count
}
