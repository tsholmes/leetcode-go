package main

import "fmt"

func main() {
	fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("226"))
}

func numDecodings(s string) int {
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	counts := make([]int, len(s)+1)
	counts[0] = 1
	counts[1] = 1
	for i := 2; i <= len(s); i++ {
		c := s[i-1]
		b := s[i-2]
		if c > '0' {
			counts[i] += counts[i-1]
		}
		if b == '1' || (b == '2' && c <= '6') {
			counts[i] += counts[i-2]
		}
	}
	return counts[len(s)]
}
