package main

import "fmt"

func main() {
	fmt.Println(shortestPalindrome("aacecaaa"))
	fmt.Println(shortestPalindrome("abcd"))
}

func shortestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	for l := len(s); ; l++ {
		valid := true
		for i := l / 2; i >= 0; i-- {
			j := l - i - 1
			si := i - (l - len(s))
			sj := j - (l - len(s))
			if si < 0 {
				break
			}
			if s[si] != s[sj] {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}

		res := make([]byte, l)
		copy(res[l-len(s):], s)
		for i := 0; i < l/2; i++ {
			j := l - i - 1
			res[i] = res[j]
		}
		return string(res)
	}
}
