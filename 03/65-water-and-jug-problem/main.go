package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))
	fmt.Println(canMeasureWater(2, 6, 5))
}

func canMeasureWater(x int, y int, z int) bool {
	seen := map[[2]int]bool{}
	var search func(int, int) bool
	search = func(a int, b int) bool {
		if a+b == z {
			return true
		}
		k := [2]int{a, b}
		if seen[k] {
			return false
		}
		seen[k] = true
		if search(a, 0) {
			return true
		}
		if search(0, b) {
			return true
		}
		if search(x, b) {
			return true
		}
		if search(a, x) {
			return true
		}
		if b < y && a > 0 {
			t := y - b
			if t > a {
				t = a
			}
			if search(a-t, b+t) {
				return true
			}
		}
		if a < x && b > 0 {
			t := x - a
			if t > b {
				t = b
			}
			if search(a+t, b-t) {
				return true
			}
		}
		return false
	}
	return search(0, 0)
}
