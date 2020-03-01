package main

import "fmt"

func main() {
	fmt.Println(getNoZeroIntegers(1010))
}

func hasZeros(n int) bool {
	for n > 0 {
		d := n % 10
		if d == 0 {
			return true
		}
		n = n / 10
	}
	return false
}

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		j := n - i
		if !hasZeros(i) && !hasZeros(j) {
			return []int{i, j}
		}
	}
	return nil
}
