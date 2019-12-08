package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	b := []byte(s)
	lasts := [256]int{}
	for i := range lasts {
		lasts[i] = -1
	}
	starts := make([]int, len(s))
	for i, c := range b {
		starts[i] = lasts[c] + 1
		lasts[c] = i
	}

	max := 0
	curStart := 0
	for i := range b {
		if starts[i] > curStart {
			curStart = starts[i]
		}
		l := i - curStart + 1
		if l > max {
			max = l
		}
	}

	return max
}
