package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("cbbd"))
	fmt.Println(longestPalindrome("bbbb"))
	fmt.Println(longestPalindrome("bbbbb"))
}

func longestPalindrome(s string) string {
	max := ""

	check := func(l, r int) {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		if r-l > len(max) {
			max = s[l+1 : r]
		}
	}

	for i := range s {
		check(i, i)
		check(i, i+1)
	}

	return max
}
