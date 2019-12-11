package main

import "fmt"

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("AEAEEEBCA", "AABC"))
	fmt.Println(minWindow("bba", "ab"))
}

func minWindow(s string, t string) string {
	counts := map[byte]int{}
	for i := 0; i < len(t); i++ {
		counts[t[i]]++
	}

	start := 0
	for start < len(s) && counts[s[start]] == 0 {
		start++
	}
	if start == len(s) {
		return ""
	}
	seen := map[byte]int{}
	seenCount := 0
	end := start
	for end < len(s) && seenCount < len(t) {
		c := s[end]
		if seen[c] < counts[c] {
			seenCount++
		}
		seen[c]++
		end++
		for start < end && seen[s[start]] > counts[s[start]] {
			seen[s[start]]--
			start++
		}
	}
	if seenCount < len(t) {
		return ""
	}

	min := s[start:end]

	for end < len(s) {
		c := s[end]
		seen[c]++
		end++
		for start < end && seen[s[start]] > counts[s[start]] {
			seen[s[start]]--
			start++
		}
		if end-start < len(min) {
			min = s[start:end]
		}
	}

	return min
}
