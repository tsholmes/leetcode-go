package main

import "fmt"

func main() {
	fmt.Println(longestPrefix("leetcodeleet"))
	fmt.Println(longestPrefix("abcda"))
}

func longestPrefix(s string) string {
	lps := make([]int, len(s))

	cur := 0
	i := 1
	for i < len(s) {
		if s[i] == s[cur] {
			cur++
			lps[i] = cur
			i++
		} else {
			if cur != 0 {
				cur = lps[cur-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}

	return s[:lps[len(s)-1]]
}
