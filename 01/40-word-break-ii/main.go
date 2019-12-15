package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
	fmt.Println(wordBreak("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"}))
	fmt.Println(wordBreak(
		"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		[]string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"},
	))
}

func wordBreak(s string, wordDict []string) []string {
	maxLen := 0
	exists := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		exists[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	var res []string
	var work []string
	var search func(int) bool
	bad := map[int]bool{}
	search = func(i int) bool {
		if i == len(s) {
			res = append(res, strings.Join(work, " "))
			return true
		}
		if bad[i] {
			return false
		}
		wl := len(work)
		good := false
		for l := 1; l <= maxLen && i+l <= len(s); l++ {
			if !exists[s[i:i+l]] {
				continue
			}
			work = append(work, s[i:i+l])
			good = search(i+l) || good
			work = work[:wl]
		}
		if !good {
			bad[i] = true
		}
		return good
	}

	search(0)

	return res
}
