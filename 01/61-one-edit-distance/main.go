package main

import "fmt"

func main() {
	fmt.Println(isOneEditDistance("ab", "acb"))
	fmt.Println(isOneEditDistance("cab", "ad"))
	fmt.Println(isOneEditDistance("1203", "1213"))
	fmt.Println(isOneEditDistance("ab", "abc"))
	fmt.Println(isOneEditDistance("abc", "ab"))
	fmt.Println(isOneEditDistance("abc", "abd"))
}

func isOneEditDistance(s string, t string) bool {
	if s == t {
		return false
	}
	for i := 0; i <= len(s); i++ {
		// insert
		if i < len(t) {
			if s[:i] == t[:i] && s[i:] == t[i+1:] {
				return true
			}
		}
		// delete
		if i < len(s) && i <= len(t) {
			if s[:i] == t[:i] && s[i+1:] == t[i:] {
				return true
			}
		}
		// replace
		if i < len(s) && i < len(t) {
			if s[:i] == t[:i] && s[i+1:] == t[i+1:] {
				return true
			}
		}
	}
	return false
}
