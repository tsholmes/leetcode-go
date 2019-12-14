package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a"))
	fmt.Println(isMatch("aa", "*"))
	fmt.Println(isMatch("cb", "?a"))
	fmt.Println(isMatch("adceb", "*a*b"))
	fmt.Println(isMatch("acdcb", "a*c?b"))
}

func isMatch(s string, p string) bool {
	fwds := make([][]int, len(p))
	for i := 0; i < len(p); i++ {
		if p[i] == '*' {
			fwds[i] = append(fwds[i], i)
		}
		fwds[i] = append(fwds[i], i+1)
		for j := i + 1; j < len(p) && p[j] == '*'; j++ {
			fwds[i] = append(fwds[i], j+1)
		}
	}

	state := map[int]struct{}{}
	state[0] = struct{}{}
	for i := 0; i < len(p) && p[i] == '*'; i++ {
		state[i+1] = struct{}{}
	}

	for i := 0; i < len(s); i++ {
		next := map[int]struct{}{}
		for n := range state {
			if n < len(p) && (p[n] == '*' || p[n] == '?' || p[n] == s[i]) {
				for _, f := range fwds[n] {
					next[f] = struct{}{}
				}
			}
		}
		state = next
	}

	_, ok := state[len(p)]
	return ok
}
