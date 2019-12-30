package main

import "fmt"

func main() {
	fmt.Println(canPermutePalindrome("aab"))
}

func canPermutePalindrome(s string) bool {
	var low, high uint

	for _, c := range []byte(s) {
		if c < 64 {
			low ^= 1 << uint(c)
		} else {
			high ^= 1 << uint(c-64)
		}
	}

	if len(s)&1 == 0 {
		return low == 0 && high == 0
	} else {
		low ^= high
		return low&-low == low && low != 0
	}
}
