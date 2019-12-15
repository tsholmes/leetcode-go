package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	maxLen := 0
	exists := make(map[string]bool, len(wordDict))
	for _, w := range wordDict {
		exists[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	good := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		for l := 1; l <= maxLen && (i-l+1) >= 0; l++ {
			j, k := i-l+1, i+1
			if j > 0 && !good[j-1] {
				continue
			}
			w := s[j:k]
			if exists[w] {
				good[i] = true
				break
			}
		}
	}

	return good[len(s)-1]
}
