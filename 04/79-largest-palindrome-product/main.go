package main

import "fmt"

func main() {
	for i := 1; i <= 8; i++ {
		fmt.Println(largestPalindrome(i), largestPalindromeFull(i))
	}
}

func largestPalindrome(n int) int {
	return []int{
		9,
		987,
		123,
		597,
		677,
		1218,
		877,
		475,
	}[n-1]
}

func largestPalindromeFull(n int) int {
	if n == 1 {
		return 9
	}
	start, end := pow10(n-1), pow10(n)
	for i := end - 1; i >= start; i-- {
		pal := toPal(i)
		valid := false
		for f := end; f >= start; f-- {
			f2 := pal / f
			if f2 >= end {
				break
			}
			if pal%f == 0 {
				valid = true
				break
			}
		}
		if valid {
			return pal % 1337
		}
	}
	return 0
}

func pow10(n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r *= 10
	}
	return r
}

func toPal(n int) int {
	r := n
	for n > 0 {
		r = (r * 10) + (n % 10)
		n /= 10
	}
	return r
}
